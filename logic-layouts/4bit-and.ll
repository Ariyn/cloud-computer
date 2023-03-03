inputs 8

define and1 and 2 1
connect and1.i1 $inputs.1
connect and1.i2 $inputs.5

define and2 and 2 1
connect and2.i1 $inputs.2
connect and2.i2 $inputs.6

define and3 and 2 1
connect and3.i1 $inputs.3
connect and3.i2 $inputs.7

define and4 and 2 1
connect and4.i1 $inputs.4
connect and4.i2 $inputs.8

alias o1 and1.o1
alias o2 and2.o1
alias o3 and3.o1
alias o4 and4.o1
