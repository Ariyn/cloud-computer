package cloud_computer

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func parseArguments() {
	for index := 1; index < len(os.Args); index++ {
		arg := os.Args[index]
		if !strings.HasPrefix(arg, "-") {
			continue
		}

		if strings.HasPrefix(arg, "-i") {
			// TODO: -i 뒤의 숫자를 파싱해서 해당 번호에 넣기
			//log.Println(arg, os.Args[index+1])
			Inputs = append(Inputs, os.Args[index+1])
			index += 1
		} else if arg == "-name" {
			Name = os.Args[index+1]
		} else if arg == "-no-optimization" {
			UseOptimization = false
		} else if strings.HasPrefix(arg, "-o") {
			// TODO: -o 뒤의 숫자를 파싱해서 해당 번호에 넣기
			Outputs = append(Outputs, os.Args[index+1])
			index += 1
		} else if arg == "-debug" {
			if os.Args[index+1] == "1" {
				IsDebugging = true
				log.Println("RUNNING DEBUGGING MODE")
			}
			index += 1
		}
	}
}

func equalOutputs(v1, v2 []bool) bool {
	if len(v1) != len(v2) {
		return false
	}

	for i := range v1 {
		if v1[i] != v2[i] {
			return false
		}
	}

	return true
}

func ParseInputs(inputs ...string) (elements []Element) {
	for _, i := range inputs {
		words := strings.Split(i, ".")
		e := parseElement(words...)
		elements = append(elements, e)
	}

	return
}

func CreateOutputs(size int) (elements []Element) {
	for i := 0; i < size; i++ {
		e := Element{
			Part:    fmt.Sprintf("o%d", i+1),
			IsAlias: false,
		}

		elements = append(elements, e)
	}
	return
}
