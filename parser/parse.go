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

const (
	bashInputParameterFormat = "-i%d) i%d=$2; shift; ;;"
	bashInputCheckFormat     = `if [ -z ${i%d+x} ]; then
	i%d="inputs.%d"
fi`
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

	f, err := os.Open(strings.Join([]string{path, file}, "/"))
	if err != nil {
		panic(err)
	}

	defer f.Close()

	bytes, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	llScript := string(bytes)
	if strings.HasSuffix(file, "rom") {
		llScript, err = parseRom(string(bytes))
		if err != nil {
			panic(err)
		}
	}

	bashScript, err := parse(llScript)
	if err != nil {
		panic(err)
	}

	fmt.Println(bashScript)
}

func parse(script string) (bash string, err error) {
	commandsByName := make(map[string]command)
	commandsByName["inputs"] = &definition{
		typ:        "input",
		name:       "inputs",
		outputSize: 2,
	}

	inputSize := 2

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
			var inputSize, outputSize int
			if len(words) <= 3 && !isPreCompiledGate(words[1]) {
				//inputSize
			} else {
				inputSize, err = strconv.Atoi(words[3])
				if err != nil {
					panic(err)
					return bash, err
				}

				if len(words) <= 4 {
					log.Println("Wrong define", strings.Join(words, " "))
					panic("wrong")
				}

				outputSize, err = strconv.Atoi(words[4])
				if err != nil {
					log.Println("Wrong define", strings.Join(words, " "))
					panic(err)
					return bash, err
				}
			}

			noOptimization := false
			if 4 < len(words) && words[len(words)-1] == "no-optimization" {
				noOptimization = true
			}

			def := parseDefine(words[1], words[2], inputSize, outputSize, noOptimization)
			commandsByName[def.name] = &def
		case "connect":
			e1, err := parseElement(words[1])
			if err != nil {
				return bash, err
			}
			if _, ok := commandsByName[e1.GateName]; !ok {
				log.Println(e1.GateName, "not declared gate name")
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
				if !(e2.IsParameter || e2.IsStaticValue) {
					err = cc.InvalidElement
					log.Println(e2.GateName)
					panic(err)
					return bash, err
				}
			}
			//outputGate := definitionsByName[e2.GateName]
			//log.Println("e1.part", e1.Part)
			inputGate.inputConnection = append(inputGate.inputConnection, fmt.Sprintf(`-%s "%s"`, e1.Part, e2.Bash()))

			commandsByName[e1.GateName] = inputGate

		case "alias":
			alias, err := parseAlias(words[1], words[2])
			if err != nil {
				panic(err)
				return "", nil
			}

			commandsByName[alias.Name] = alias

		case "inputs":
			inputSize, err = strconv.Atoi(words[1])
			if err != nil {
				panic(err)
			}
		}

		//log.Println(len(line), line)
	}

	delete(commandsByName, "inputs")

	bashLines := make([]string, 0)
	for i := 0; i < inputSize; i++ {
		bashLines = append(bashLines, fmt.Sprintf("bin/input -name \"${name_variable}i%d\" -debug \"${debug}\"  -i1 \"${i%d}\"", i+1, i+1))
	}

	for _, cmd := range commandsByName {
		//log.Println(cmd.Bash())
		bashLines = append(bashLines, cmd.Bash())
		//bashLines = append(bashLines, fmt.Sprintf(`(%s && wait)`, cmd.Bash()))
	}

	bash = `#!/bin/bash

trap "kill 0" SIGINT

while (( $# )); do
  case $1 in
    -name) name=$2;shift; ;;
    -child) child=1; ;;
    -no-optimization) ;;
    -debug) debug=$2; shift;;

{input_parameters}

    *) echo "unknown $1"; break
  esac
  shift
done

{input_checks}

if [ -n "$name" ]; then
  name_variable="${name}."
fi

if [ -z ${child+x} ]; then
{input_names}
fi

` + strings.Join(bashLines, "& \\\n")

	inputParameters := make([]string, 0)
	inputCheckers := make([]string, 0)
	inputNames := make([]string, 0)
	for i := 1; i <= inputSize; i++ {
		pf := "    " + fmt.Sprintf(bashInputParameterFormat, i, i)
		inputParameters = append(inputParameters, pf)

		cf := fmt.Sprintf(bashInputCheckFormat, i, i, i)
		inputCheckers = append(inputCheckers, cf)

		inputNames = append(inputNames, fmt.Sprintf("  i%d=${name_variable}${i%d}", i, i))
	}

	bash = strings.ReplaceAll(bash, "{input_parameters}", strings.Join(inputParameters, "\n\n"))

	bash = strings.ReplaceAll(bash, "{input_checks}", strings.Join(inputCheckers, "\n\n"))

	bash = strings.ReplaceAll(bash, "{input_names}", strings.Join(inputNames, "\n\n"))

	return bash, nil
}

func isPreCompiledGate(s string) (r bool) {
	switch s {
	case "xor":
		r = true
	case "and":
		r = true
	case "or":
		r = true
	case "input":
		r = true
	case "alias":
		r = true
	case "not":
		r = true
	case "flipflop":
		r = true
	}
	return
}

type definition struct {
	typ             string
	bin             string
	name            string
	inputSize       int
	outputSize      int
	inputConnection []string
	noOptimization  bool
}

func (def *definition) Bash() string {
	child := ""
	if !(def.typ == "and" || def.typ == "or" || def.typ == "not" || def.typ == "xor" || def.typ == "flipflop") {
		child = "-child"
	}

	inputs := make([]string, len(def.inputConnection))
	for i, c := range def.inputConnection {
		inputs[i] = fmt.Sprintf(`%s`, c)
	}

	inputArgument := strings.Join(inputs, " ")

	noOptimization := ""
	if def.noOptimization {
		noOptimization = "-no-optimization"
	}
	return fmt.Sprintf(`%s %s -name "${name_variable}%s" -debug "${debug}" %s %s`, def.bin, noOptimization, def.name, child, inputArgument)
}

func parseDefine(name, typ string, inputSize, outputSize int, noOptimization bool) (def definition) {
	def.typ = typ
	def.bin = "bin/" + typ
	def.name = name
	def.inputSize = inputSize
	def.outputSize = outputSize
	def.noOptimization = noOptimization
	def.inputConnection = make([]string, 0)

	return
}

func parseElement(word string) (element cc.Element, err error) {
	if word[0] == '$' {
		element.GateName = word
		element.IsParameter = true
		return
	}

	if n, err := strconv.Atoi(word); err == nil {
		element.IsStaticValue = true
		if n == 0 {
			element.StaticValue = false
		} else {
			element.StaticValue = true
		}
		//log.Println(word, "skipping")
		return element, nil
	}

	elements := strings.Split(word, ".")
	if len(elements) < 1 {
		err = cc.InvalidElement
		panic(err)
		return
	}

	element.GateName = elements[0]

	if _, err := strconv.Atoi(element.GateName); err == nil {
		return element, nil
	}

	element.Part = strings.Join(elements[1:], ".")

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
	target := a.Target.String()

	if target == "" && a.Target.IsStaticValue {
		if a.Target.StaticValue == true {
			target = "1"
		} else {
			target = "0"
		}
		return fmt.Sprintf(`bin/alias -name "${name_variable}%s" -i1 "%s"`, a.Name, target)
	}

	if target[0] == '$' && 7 <= len(target) && target[1:7] == "inputs" {
		target = fmt.Sprintf("${i%s}", strings.ReplaceAll(target, "$inputs.", ""))
	}

	if !(target == "0" || target == "1") && target[0] != '$' {
		target = "${name_variable}" + target
	}

	return fmt.Sprintf(`bin/alias -name "${name_variable}%s" -i1 "%s"`, a.Name, target)
}

func parseAlias(name, target string) (alias Alias, err error) {
	alias.Target, err = parseElement(target)
	if err != nil {
		return
	}

	alias.Name = name

	return
}
