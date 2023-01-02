package main

import (
	"context"
	"flag"
	cc "github.com/ariyn/cloud-computer"
	"github.com/go-redis/redis"
	"log"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

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

var nameRegex *regexp.Regexp

func init() {
	//flag.Var(&names, "names", "names for watch")
	flag.StringVar(&name, "name", "", "name for watch")
	//flag.BoolVar(&input, "I", false, "isInput")
	//flag.BoolVar(&output, "O", false, "isOutput")

	nameRegex = regexp.MustCompile(`.+?\.[io](\d+)`)
}

type Watch struct {
	Input  Watches
	Output Watches
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

func getWatches(client *redis.Client, name string) Watches {
	w, err := client.SMembers(name).Result()
	if err != nil {
		panic(err)
	}

	sort.Slice(w, func(i, j int) bool {
		first := findNumber(w[i])
		second := findNumber(w[j])
		return first < second
	})

	channels := make([]<-chan bool, 0)
	for _, name := range w {
		channels = append(channels, cc.ReadAsyncRedis(context.TODO(), client, name))
	}

	return Watches{
		channels: channels,
		values:   make([]bool, len(channels)),
	}
}

// TODO: 이부분 RunRedis와 거의 동일함. 추상화 할 방법 찾아보기
func main() {
	flag.Parse()
	if input && output {
		panic("input and output can not be both set")
	}

	log.Println("running")
	client := cc.ConnectRedis()
	log.Println("connected")

	watches := make(map[string]Watch)
	w := Watch{}
	w.Input = getWatches(client, name+".inputs")
	w.Output = getWatches(client, name+".outputs")

	watches[name] = w

	cases := make([]reflect.SelectCase, 0)
	i := watches[name].Input
	cases = append(cases, i.getCases()...)

	o := watches[name].Output
	cases = append(cases, o.getCases()...)

	for {
		index, value, ok := reflect.Select(cases)
		if !ok {
			break
		}

		if index < i.Length() {
			i.setValue(index, value.Bool())
		} else {
			o.setValue(index-i.Length(), value.Bool())
		}

		printBits("inputs", i.values)
		printBits("outputs", o.values)
		//log.Printf("%d: %v", index, previousValues[index])
	}
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

func printBits(prefix string, bits []bool) {
	l := ""
	for i := range bits {
		b := bits[len(bits)-i-1]
		if b {
			l += "1"
		} else {
			l += "0"
		}
	}

	log.Println(prefix, l)
}
