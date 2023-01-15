package main

import (
	"strconv"
	"strings"
)

func parseRom(script string) (llScript string, err error) {
	llScript = `inputs 5
# inputs.1 ~ inputs.4 = selector
#  inputs.1 = LSB, 4 = MSB
# inputs.5 = clock

define selector 4bit-decoder
connect selector.i1 $inputs.1
connect selector.i2 $inputs.2
connect selector.i3 $inputs.3
connect selector.i4 $inputs.4

define sa 4bit-adder-subtractor 9 1
connect sa.i1 $inputs.1
connect sa.i2 $inputs.2
connect sa.i3 $inputs.3
connect sa.i4 $inputs.4
connect sa.i5 1
connect sa.i6 0
connect sa.i7 0
connect sa.i8 0
connect sa.i9 0

define selector2 4bit-decoder 4 16
connect selector2.i1 sa.o1
connect selector2.i2 sa.o2
connect selector2.i3 sa.o3
connect selector2.i4 sa.o4

define or1 or 16 1
define or2 or 16 1
define or3 or 16 1
define or4 or 16 1
define or5 or 16 1
define or6 or 16 1
define or7 or 16 1
define or8 or 16 1

{rom_data}

define r 4bit-register
connect r.i1 or1.o1
connect r.i2 or2.o1
connect r.i3 or3.o1
connect r.i4 or4.o1
connect r.i6 $inputs.5

define r2 4bit-register
connect r2.i1 or5.o1
connect r2.i2 or6.o1
connect r2.i3 or7.o1
connect r2.i4 or8.o1
connect r2.i6 $inputs.5

alias o1 r.o1
alias o2 r.o2
alias o3 r.o3
alias o4 r.o4

alias o5 r2.o1
alias o6 r2.o2
alias o7 r2.o3
alias o8 r2.o4`

	iterationFormat := `define and{i}-{j} and 2 1
connect and{i}-{j}.i1 selector.o{i}
connect and{i}-{j}.i2 {bit}
connect or{j}.i{i} and{i}-{j}.o1

define and2-{i}-{j} and 2 1
connect and2-{i}-{j}.i1 selector2.o{i}
connect and2-{i}-{j}.i2 {bit}
connect or{j+4}.i{i} and2-{i}-{j}.o1`

	romData := make([]string, 0)
	for i, line := range strings.Split(script, "\n") {
		if line[0] == '#' {
			continue
		}

		for j, c := range line {
			compliedLine := strings.ReplaceAll(iterationFormat, "{i}", strconv.Itoa(i+1))
			compliedLine = strings.ReplaceAll(compliedLine, "{j}", strconv.Itoa(j+1))
			compliedLine = strings.ReplaceAll(compliedLine, "{j+4}", strconv.Itoa(j+5))
			compliedLine = strings.ReplaceAll(compliedLine, "{bit}", string(c))
			romData = append(romData, compliedLine)
		}
	}

	llScript = strings.ReplaceAll(llScript, "{rom_data}", strings.Join(romData, "\n\n"))

	return
}
