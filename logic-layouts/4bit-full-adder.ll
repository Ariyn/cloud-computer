inputs 8

define fa1 full-adder 3 3
define fa2 full-adder 3 3
define fa3 full-adder 3 3
define fa4 full-adder 3 3

connect fa1.i1 $inputs.1
connect fa1.i2 $inputs.2

connect fa2.i1 $inputs.3
connect fa2.i2 $inputs.4

connect fa3.i1 $inputs.5
connect fa3.i2 $inputs.6

connect fa4.i1 $inputs.7
connect fa4.i2 $inputs.8

connect fa2.i3 fa1.carry
connect fa3.i3 fa2.carry
connect fa4.i3 fa3.carry

alias b1 fa1.sum
alias b2 fa2.sum
alias b3 fa3.sum
alias b4 fa4.sum
alias b5 fa4.carry