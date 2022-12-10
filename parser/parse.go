package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var InvalidElement error
var file string

var (
	commentReg = regexp.MustCompile(`(#.+)$`)
)

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	flag.StringVar(&file, "file", "", "file name")

	InvalidElement = errors.New("invalid element")
}

func main() {
	flag.Parse()

	if file == "" {
		panic("no file")
	}

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	file, err := os.Open(strings.Join([]string{path, file}, "/"))
	if err != nil {
		panic(err)
	}

	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	script, err := parse(string(bytes))
	if err != nil {
		panic(err)
	}

	fmt.Println(script)
}

func parse(script string) (bash string, err error) {
	definitionsByName := make(map[string]definition)
	definitionsByName["inputs"] = definition{
		typ:        "input",
		name:       "inputs",
		outputSize: 2,
	}

	lines := strings.Split(script, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		line = commentReg.ReplaceAllString(line, "")

		if line == "" {
			continue
		}

		// TODO: rename words to more accurate name
		words := strings.Split(line, " ")

		// TODO: validate words count
		switch words[0] {
		case "define":
			inputSize, err := strconv.Atoi(words[3])
			if err != nil {
				return bash, err
			}

			outputSize, err := strconv.Atoi(words[4])
			if err != nil {
				return bash, err
			}

			def := parseDefine(words[1], words[2], inputSize, outputSize)
			definitionsByName[def.name] = def

		case "connect":
			e1, err := parseElement(words[1])
			if err != nil {
				return bash, err
			}
			if _, ok := definitionsByName[e1.GateName]; !ok {
				err = InvalidElement
				return bash, err
			}
			inputGate := definitionsByName[e1.GateName]

			e2, err := parseElement(words[2])
			if err != nil {
				return bash, err
			}
			if _, ok := definitionsByName[e2.GateName]; !ok {
				err = InvalidElement
				return bash, err
			}
			//outputGate := definitionsByName[e2.GateName]

			inputGate.inputConnection = append(inputGate.inputConnection, fmt.Sprintf("-%s %s", e1.Part, e2.String()))

			definitionsByName[e1.GateName] = inputGate
		}

		//log.Println(len(line), line)
	}

	delete(definitionsByName, "inputs")

	bashLines := []string{
		"#!/bin/bash",
	}
	for _, def := range definitionsByName {
		log.Println(def.Bash())
		bashLines = append(bashLines, fmt.Sprintf(`(%s && wait) & \`, def.Bash()))
	}

	return strings.Join(bashLines, "\n"), nil
}

type definition struct {
	typ             string
	bin             string
	name            string
	inputSize       int
	outputSize      int
	inputConnection []string
}

func (def *definition) Bash() string {
	inputs := make([]string, def.inputSize)
	for i, c := range def.inputConnection {
		inputs[i] = fmt.Sprintf(`%s`, c)
	}

	inputArgument := strings.Join(inputs, " ")
	return fmt.Sprintf(`%s -name '%s' %s`, def.bin, def.name, inputArgument)
}

func parseDefine(name, typ string, inputSize, outputSize int) (def definition) {
	def.typ = typ
	def.bin = "bin/" + typ
	def.name = name
	def.inputSize = inputSize
	def.outputSize = outputSize
	def.inputConnection = make([]string, 0)

	return
}

type Element struct {
	GateName string
	Part     string
}

func (e Element) String() string {
	return fmt.Sprintf(`%s.%s`, e.GateName, e.Part)
}

func (e Element) IsValidPart() (err error) {
	// TODO: not enough validation
	if e.Part == "i1" || e.Part == "i2" || e.Part == "1" || e.Part == "2" || e.Part == "o1" {
		return nil
	}

	return InvalidElement
}

func parseElement(word string) (element Element, err error) {
	elements := strings.Split(word, ".")
	if len(elements) < 1 {
		err = InvalidElement
		return
	}

	element.GateName = elements[0]
	element.Part = elements[1]

	err = element.IsValidPart()
	if err != nil {
		return
	}

	return
}
