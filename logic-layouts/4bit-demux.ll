inputs 4

# 1 inputs, 2 selectors, 4 outputs

define a1 and 3 1 no-optimization
define a2 and 3 1 no-optimization
define a3 and 3 1 no-optimization
define a4 and 3 1 no-optimization

connect a1.i3 $inputs.3
connect a2.i3 $inputs.3
connect a3.i3 $inputs.3
connect a4.i3 $inputs.3

define selector-not1-1 not 1 1 no-optimization
connect selector-not1-1.i1 $inputs.1
define selector-not2-1 not 1 1 no-optimization
connect selector-not2-1.i1 $inputs.2

define selector-not1-2 not 1 1 no-optimization
connect selector-not1-2.i1 selector-not1-1.o1

define selector-not2-2 not 1 1 no-optimization
connect selector-not2-2.i1 selector-not2-1.o1

connect a1.i1 selector-not1-1.o1
connect a1.i2 selector-not2-1.o1

connect a2.i1 selector-not1-2.o1
connect a2.i2 selector-not2-1.o1

connect a3.i1 selector-not1-1.o1
connect a3.i2 selector-not2-2.o1

connect a4.i1 selector-not1-2.o1
connect a4.i2 selector-not2-2.o1

alias o1 a1.o1
alias o2 a2.o1
alias o3 a3.o1
alias o4 a4.o1
