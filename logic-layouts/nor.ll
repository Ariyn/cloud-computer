inputs 2

define or or 2 1
define not not 1 1

connect or.i1 $inputs.1
connect or.i2 $inputs.2
connect not.i1 or.o1

alias output not.o1