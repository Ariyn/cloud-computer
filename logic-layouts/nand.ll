inputs 2

define and and 2 1 no-optimization
define not not 2 1 no-optimization

connect and.i1 $inputs.1
connect and.i2 $inputs.2
connect not.i1 and.o1

alias output not.o1
alias o1 not.o1