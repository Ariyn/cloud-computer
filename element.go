package cloud_computer

import (
	"fmt"
	"strconv"
	"strings"
)

// TODO: merge this and parse.Element
type Element struct {
	GateName      string
	Part          string
	IsAlias       bool
	IsParameter   bool
	StaticValue   bool
	IsStaticValue bool
}

func (e Element) String() string {
	es := []string{e.GateName}
	if !e.IsAlias && e.Part != "" {
		es = append(es, e.Part)
	}

	return strings.Join(es, ".")
}

func (e Element) Bash() string {
	name := e.String()

	if name == "" && e.IsStaticValue {
		name = ""
		if e.StaticValue == true {
			name = "1"
		} else {
			name = "0"
		}
		return name
	}

	if name[0] == '$' && 7 <= len(name) && name[1:7] == "inputs" {
		name = fmt.Sprintf("${i%s}", strings.ReplaceAll(name, "$inputs.", ""))
		return name
	}

	return fmt.Sprintf("${name_variable}%s", name)
}

func parseElement(words ...string) Element {
	if n, err := strconv.Atoi(words[0]); err == nil {
		e := Element{
			IsStaticValue: true,
		}

		if n == 0 {
			e.StaticValue = false
		} else {
			e.StaticValue = true
		}

		return e
	}

	e := Element{
		GateName: strings.Join(words, "."),
	}
	return e
}
