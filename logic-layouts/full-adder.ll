define ha1 half-adder 2 2
define ha2 half-adder 2 2
define or1 or 2 1

connect ha2.inputs.1 ${name}.ha1.sum
connect ha2.inputs.2 ${name}.inputs.3

connect ha1.inputs.1 ${name}.inputs.1
connect ha1.inputs.2 ${name}.inputs.2

connect or1.inputs.1 ${name}.ha1.carry
connect or1.inputs.2 ${name}.ha2.carry

alias carry ${name}.or1.o1
alias sum ${name}.ha2.carry
