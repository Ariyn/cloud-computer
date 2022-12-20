package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	cc "github.com/ariyn/cloud-computer"
	"log"
	"os"
	"strings"
)

var InvalidInput = errors.New("invalid input: input must contain nothing but 1 or 0")

func main() {
	log.Println("running")
	client := cc.ConnectRedis()
	log.Println("connected")

	log.Println(cc.Outputs)
	outputs := make([]chan<- bool, 0)
	for _, output := range cc.Outputs {
		outputs = append(outputs, cc.WriteAsyncRedis(context.Background(), client, output))
	}

	for {
		raw, currentValues, err := inputsFromUser()
		if err != nil {
			log.Printf("'%s' is invalid error", raw)
		}

		printBits(currentValues)
		for i, ch := range outputs {
			ch <- currentValues[i]
		}
	}
}

func inputsFromUser() (raw string, values []bool, err error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	raw, err = reader.ReadString('\n')
	if err != nil {
		return
	}

	raw = strings.TrimSpace(raw)

	values, err = parseBinaryString(raw)
	log.Println(values)
	return
}

func parseBinaryString(s string) (values []bool, err error) {
	if len(s) != (strings.Count(s, "1") + strings.Count(s, "0")) {
		err = InvalidInput
		return
	}

	length := len(s)
	values = make([]bool, length)
	for i, c := range s {
		if c == '0' {
			values[length-i-1] = false
		} else {
			values[length-i-1] = true
		}
	}

	return
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
