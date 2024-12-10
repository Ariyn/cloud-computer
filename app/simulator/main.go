package main

import (
	"context"
	"flag"
	cloud_computer "github.com/ariyn/cloud-computer"
	"github.com/ariyn/cloud-computer/gate"
	"io"
	"log"
	"os"
)

var file string

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	flag.StringVar(&file, "file", "", "file name")

	cloud_computer.IsVerbose = true
}

func getTruthTableCode(file string) string {
	var truthTableCode string
	if f, err := os.Open(file); err != nil {
		panic(err)
	} else {
		defer f.Close()
		if bytes, err := io.ReadAll(f); err != nil {
			panic(err)
		} else {
			truthTableCode = string(bytes)
		}
	}

	return truthTableCode
}

// TODO: simulator라는 이름보다 좀 더 명확한 이름이 있을 것
func main() {
	flag.Parse()
	if file == "" {
		panic("no file")
	}

	truthTableCode := getTruthTableCode(file)
	tt := cloud_computer.NewTruthTableWithCode(truthTableCode)

	inputElements := cloud_computer.ParseInputs(tt.Inputs...)
	outputElements := cloud_computer.CreateOutputs(tt.OutputSize)

	pg := gate.NewPGGate("pg", inputElements, outputElements, tt)
	ds := &cloud_computer.MockupDataSource{}

	pg.Init(context.TODO(), ds)
	results, changed := pg.Handler(0, true)
	log.Println(results, changed)

	results, changed = pg.Handler(1, true)
	log.Println(results, changed)

	results, changed = pg.Handler(1, false)
	log.Println(results, changed)

	results, changed = pg.Handler(0, false)
	log.Println(results, changed)
}
