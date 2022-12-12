inputs 3

define ha1 half-adder 2 2
define ha2 half-adder 2 2
define or1 or 2 1

connect ha2.i1 ha1.sum
connect ha2.i2 $inputs.3

connect ha1.i1 $inputs.1
connect ha1.i2 $inputs.2

connect or1.i1 ha1.carry
connect or1.i2 ha2.carry

alias carry or1.o1
alias sum ha2.sum
