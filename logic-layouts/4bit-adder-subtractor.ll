inputs 9

# inputs.9 adder or subtractor selector
#   0 = add, 1 = subtractor

define fa1 full-adder 3 3
define fa2 full-adder 3 3
define fa3 full-adder 3 3
define fa4 full-adder 3 3

define x1 xor 2 1
connect x1.i2 $inputs.9

define x2 xor 2 1
connect x2.i2 $inputs.9

define x3 xor 2 1
connect x3.i2 $inputs.9

define x4 xor 2 1
connect x4.i2 $inputs.9

connect fa1.i1 $inputs.1
connect x1.i1 $inputs.5
connect fa1.i2 x1.o1

connect fa2.i1 $inputs.2
connect x2.i1 $inputs.6
connect fa2.i2 x2.o1

connect fa3.i1 $inputs.3
connect x3.i1 $inputs.7
connect fa3.i2 x3.o1

connect fa4.i1 $inputs.4
connect x4.i1 $inputs.8
connect fa4.i2 x4.o1

connect fa1.i3 $inputs.9
connect fa2.i3 fa1.carry
connect fa3.i3 fa2.carry
connect fa4.i3 fa3.carry

alias o1 fa1.sum
alias o2 fa2.sum
alias o3 fa3.sum
alias o4 fa4.sum
alias o5 fa4.carry