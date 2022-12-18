inputs 9
# inputs.1 ~ inputs.4 = a
# inputs.5 ~ inputs.8 = b
# inputs.9 = carry in

define fa1 full-adder 3 3
define fa2 full-adder 3 3
define fa3 full-adder 3 3
define fa4 full-adder 3 3

connect fa1.i1 $inputs.1
connect fa2.i1 $inputs.2
connect fa3.i1 $inputs.3
connect fa4.i1 $inputs.4

connect fa1.i2 $inputs.5
connect fa2.i2 $inputs.6
connect fa3.i2 $inputs.7
connect fa4.i2 $inputs.8

connect fa1.i3 $inputs.9
connect fa2.i3 fa1.carry
connect fa3.i3 fa2.carry
connect fa4.i3 fa3.carry

alias o1 fa1.sum
alias o2 fa2.sum
alias o3 fa3.sum
alias o4 fa4.sum
alias o5 fa4.carry
alias carry fa4.carry