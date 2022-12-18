inputs 2

# 2 selectors, 4 outputs

define a1 and 2 1 no-optimization
define a2 and 2 1 no-optimization
define a3 and 2 1 no-optimization
define a4 and 2 1 no-optimization

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
