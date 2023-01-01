package main

import (
	"context"
	"flag"
	cc "github.com/ariyn/cloud-computer"
	"log"
	"reflect"
	"regexp"
	"sort"
	"strconv"
)

// type arrayFlags []string
//
//	func (af *arrayFlags) String() string {
//		return strings.Join(*af, "\n")
//	}
//
//	func (af *arrayFlags) Set(value string) error {
//		*af = append(*af, value)
//		return nil
//	}
//
// var watches arrayFlags
var name string
var input bool
var output bool

var nameRegex *regexp.Regexp

func init() {
	//flag.Var(&watches, "names", "names for watch")
	flag.StringVar(&name, "name", "", "name for watch")
	flag.BoolVar(&input, "I", false, "isInput")
	flag.BoolVar(&output, "O", false, "isOutput")

	nameRegex = regexp.MustCompile(`.+?\.[io](\d+)`)
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

	memberName := name
	if input {
		memberName += ".inputs"
	}
	if output {
		memberName += ".outputs"
	}
	watches, err := client.SMembers(memberName).Result()
	if err != nil {
		panic(err)
	}

	sort.Slice(watches, func(i, j int) bool {
		first := findNumber(watches[i])
		second := findNumber(watches[j])
		return first < second
	})

	inputs := make([]<-chan bool, 0)
	for _, name := range watches {
		log.Println(name)
		inputs = append(inputs, cc.ReadAsyncRedis(context.TODO(), client, name))
	}

	cases := make([]reflect.SelectCase, 0)
	for _, ch := range inputs {
		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(ch),
		})
	}

	previousValues := make([]bool, len(inputs))
	for {
		index, value, ok := reflect.Select(cases)
		if !ok {
			break
		}

		previousValues[index] = value.Bool()

		printBits(previousValues)
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

func printBits(bits []bool) {
	l := ""
	for i := range bits {
		b := bits[len(bits)-i-1]
		if b {
			l += "1"
		} else {
			l += "0"
		}
	}

	log.Println(l)
}
