package main

import (
	"flag"
	"fmt"
	cc "github.com/ariyn/cloud-computer"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var file string

var (
	commentReg = regexp.MustCompile(`(#.+)$`)
)

type command interface {
	Bash() string
}

var _ command = (*definition)(nil)

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	flag.StringVar(&file, "file", "", "file name")
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
	commandsByName := make(map[string]command)
	commandsByName["inputs"] = &definition{
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
				panic(err)
				return bash, err
			}

			outputSize, err := strconv.Atoi(words[4])
			if err != nil {
				panic(err)
				return bash, err
			}

			def := parseDefine(words[1], words[2], inputSize, outputSize)
			commandsByName[def.name] = &def
		case "connect":
			e1, err := parseElement(words[1])
			if err != nil {
				return bash, err
			}
			if _, ok := commandsByName[e1.GateName]; !ok {
				err = cc.InvalidElement
				panic(err)
				return bash, err
			}

			inputGate := commandsByName[e1.GateName].(*definition)

			e2, err := parseElement(words[2])
			if err != nil {
				panic(err)
				return bash, err
			}
			if _, ok := commandsByName[e2.GateName]; !ok {
				if !e2.IsParameter {
					err = cc.InvalidElement
					log.Println(e2.GateName)
					panic(err)
					return bash, err
				}
			}
			//outputGate := definitionsByName[e2.GateName]

			inputGate.inputConnection = append(inputGate.inputConnection, fmt.Sprintf("-inputs %s", e2.String()))

			commandsByName[e1.GateName] = inputGate

		case "alias":
			alias, err := parseAlias(words[1], words[2])
			if err != nil {
				panic(err)
				return "", nil
			}

			commandsByName[alias.Name] = alias
		}

		//log.Println(len(line), line)
	}

	delete(commandsByName, "inputs")

	bashLines := make([]string, 0)
	for _, cmd := range commandsByName {
		log.Println(cmd.Bash())
		bashLines = append(bashLines, fmt.Sprintf(`(%s && wait)`, cmd.Bash()))
	}

	bash = `#!/bin/bash

while (( $# )); do
  echo "first argument $1"
  case $1 in
    --name) name=$2;shift;
    ;;
	--i1) i1=$2; shift;
	;;
	--i2) i2=$2; shift;
	;;
    *) echo "unknown $1"; break
  esac
  shift
done

if [ -z ${i1+x} ]; then
	i1="inputs.1"
fi

if [ -z ${i2+x} ]; then
	i2="inputs.2"
fi

` + strings.Join(bashLines, "& \\\n")

	return bash, nil
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
	return fmt.Sprintf(`%s -name ${name}.%s %s`, def.bin, def.name, inputArgument)
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

func parseElement(word string) (element cc.Element, err error) {
	if word[0] == '$' {
		element.GateName = word
		element.IsParameter = true
		return
	}

	elements := strings.Split(word, ".")
	if len(elements) < 1 {
		err = cc.InvalidElement
		panic(err)
		return
	}

	element.GateName = elements[0]
	element.Part = elements[1]

	//err = element.IsValidPart()
	//if err != nil {
	//	panic(err)
	//	return
	//}

	return
}

var _ command = (*Alias)(nil)

type Alias struct {
	Name   string
	Target cc.Element
}

func (a Alias) Bash() string {
	return fmt.Sprintf(`bin/alias -name "${name}"."%s" -inputs "${name}".%s`, a.Name, a.Target.String())
}

func parseAlias(name, target string) (alias Alias, err error) {
	alias.Target, err = parseElement(target)
	if err != nil {
		return
	}

	alias.Name = name

	return
}
