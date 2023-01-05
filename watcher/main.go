package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	cc "github.com/ariyn/cloud-computer"
	"github.com/go-redis/redis/v9"
	"github.com/gosuri/uilive"
	"log"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const ESC = 27

var clear = fmt.Sprintf("%c[%dA%c[2K", ESC, 1, ESC)
var isPrinted = false

type arrayFlags []string

func (af *arrayFlags) String() string {
	return strings.Join(*af, "\n")
}

func (af *arrayFlags) Set(value string) error {
	*af = append(*af, value)
	return nil
}

// var names arrayFlags
var name string
var input bool
var output bool
var excepts arrayFlags

var nameRegex *regexp.Regexp

func init() {
	flag.Var(&excepts, "e", "will not watch itself, and it's children")
	flag.StringVar(&name, "name", "", "name for watch")

	nameRegex = regexp.MustCompile(`.+?\.[io](\d+)`)
}

type Watch struct {
	Input       *Watches
	InputIndex  int
	InputSize   int
	Output      *Watches
	OutputIndex int
	OutputSize  int
	Name        string
}

type Watches struct {
	channels []<-chan bool
	values   []bool
}

func (w *Watches) getCases() (cases []reflect.SelectCase) {
	for _, ch := range w.channels {
		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(ch),
		})
	}

	return
}

func (w *Watches) setValue(index int, value bool) {
	w.values[index] = value
}

func (w *Watches) Length() int {
	return len(w.values)
}

func getWatches(ctx context.Context, client *redis.Client, name string) *Watches {
	w, err := client.SMembers(ctx, name).Result()
	if err != nil {
		panic(err)
	}

	sort.Slice(w, func(i, j int) bool {
		first := findNumber(w[i])
		second := findNumber(w[j])
		return first < second
	})

	values := make([]bool, len(w))
	channels := make([]<-chan bool, len(w))
	for i, name := range w {
		channels[i] = cc.ReadAsyncRedis(context.TODO(), client, name)
		v, err := cc.ReadRedis(ctx, client, name+".status")
		if err != nil {
			panic(err)
		}
		values[i] = v
	}

	return &Watches{
		channels: channels,
		values:   values,
	}
}

var children []string

var writer *uilive.Writer

// TODO: 이부분 RunRedis와 거의 동일함. 추상화 할 방법 찾아보기
func main() {
	flag.Parse()
	if input && output {
		panic("input and output can not be both set")
	}

	log.Println("running")
	client := cc.ConnectRedis()
	log.Println("connected")

	var err error
	children, err = getChildren(client, name)
	if err != nil {
		log.Println("can't get children")
		panic(err)
	}

	watches := make([]Watch, 0)

	index := 0
	for _, c := range children {
		w := Watch{}
		w.Name = c
		w.Input = getWatches(context.Background(), client, c+".inputs")
		w.InputSize = w.Input.Length()
		w.InputIndex = index

		w.Output = getWatches(context.Background(), client, c+".outputs")
		w.OutputSize = w.Output.Length()
		w.OutputIndex = index + w.InputSize

		watches = append(watches, w)
		index += w.InputSize + w.OutputSize
	}

	printWatches(watches)

	cases := make([]reflect.SelectCase, 0)
	for _, w := range watches {
		cases = append(cases, w.Input.getCases()...)  // even = inputs
		cases = append(cases, w.Output.getCases()...) // odd = outputs
	}

	for {
		index, value, ok := reflect.Select(cases)
		if !ok {
			break
		}

		w, watchesIndex, err := getWatch(watches, index)
		if err != nil {
			log.Println("no such index", index)
			panic(err)
		}

		w.setValue(index-watchesIndex, value.Bool())
		printWatches(watches)
	}
}

func getWatch(watches []Watch, index int) (*Watches, int, error) {
	for _, w := range watches {
		if w.InputIndex <= index && index < (w.InputIndex+w.InputSize) {
			return w.Input, w.InputIndex, nil
		}
		if w.OutputIndex <= index && index < (w.OutputIndex+w.OutputSize) {
			return w.Output, w.OutputIndex, nil
		}
	}

	return nil, 0, errors.New("No such watch index")
}

func getChildren(client *redis.Client, name string) (children []string, err error) {
	rawChildren, err := client.SMembers(context.Background(), name+".children").Result()
	if err != nil {
		return
	}

	for _, c := range rawChildren {
		if isExcept(c) {
			continue
		}

		children = append(children, c)
	}

	// TODO: need numeric sort
	sort.Slice(children, func(i, j int) bool {
		return children[i] < children[j]
	})

	return
}

func isExcept(name string) bool {
	for _, e := range excepts {
		if strings.HasPrefix(name, e) {
			return true
		}
	}

	return false
}

func findNumber(name string) int {
	l := nameRegex.FindStringSubmatch(name)
	if len(l) == 0 {
		return -1
	}

	n, err := strconv.Atoi(l[1])
	if err != nil {
		return -1
	}

	return n
}

func printWatches(watches []Watch) {
	if isPrinted {
		for range watches {
			fmt.Print(clear)
		}
	}

	maximumSizes := make([]int, 4)
	for _, w := range watches {
		inputNameSize := len(w.Name) + 7
		inputSize := w.InputSize
		outputNameSize := len(w.Name) + 8
		outputSize := w.OutputSize

		if maximumSizes[0] < inputNameSize {
			maximumSizes[0] = inputNameSize
		}

		if maximumSizes[1] < inputSize {
			maximumSizes[1] = inputSize
		}

		if maximumSizes[2] < outputNameSize {
			maximumSizes[2] = outputNameSize
		}

		if maximumSizes[3] < outputSize {
			maximumSizes[3] = outputSize
		}
	}

	for _, w := range watches {
		printBits(maximumSizes, w.Name, w.Input.values, w.Output.values)
	}
	isPrinted = true
}

func printBits(maximumSizes []int, prefix string, inputBits, outputBits []bool) {
	inputName := fmt.Sprintf("%s.inputs", prefix)
	inputName += strings.Repeat(" ", maximumSizes[0]-len(inputName)+3)

	input := getBitString(inputBits)
	input += strings.Repeat(" ", maximumSizes[1]-len(input)+3)

	outputName := fmt.Sprintf("%s.outputs", prefix)
	outputName += strings.Repeat(" ", maximumSizes[2]-len(outputName)+3)

	output := getBitString(outputBits)
	output += strings.Repeat(" ", maximumSizes[3]-len(output)+3)

	fmt.Println(inputName, input, outputName, output)
}

func getBitString(bits []bool) string {
	l := ""
	for i := range bits {
		b := bits[len(bits)-i-1]
		if b {
			l += "1"
		} else {
			l += "0"
		}
	}

	return l
}
