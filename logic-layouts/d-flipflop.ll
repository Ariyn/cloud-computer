inputs 2
# $inputs.1 = data
# $inputs.2 = clock

define sr sr-latch 2 1
define and1 and 2 1
define and2 and 2 1
define not1 not 2 1

connect not1.i1 $inputs.1
connect and1.i1 not1.o1
connect sr.i1 and1.o1

connect and1.i2 $inputs.2

connect and2.i1 $inputs.1
connect and2.i2 $inputs.2

alias q sr.q
alias q-bar sr.q-bar