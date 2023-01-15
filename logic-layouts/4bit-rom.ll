inputs 5
# inputs.1 ~ inputs.4 = selector
#  inputs.1 = LSB, 4 = MSB
# inputs.5 = clock

define selector 4bit-decoder
connect selector.i1 $inputs.1
connect selector.i2 $inputs.2
connect selector.i3 $inputs.3
connect selector.i4 $inputs.4

define or1 or 2 1
define or2 or 2 1
define or3 or 2 1
define or4 or 2 1

define and1-1 and 2 1
connect and1-1.i1 selector.o1
connect and1-1.i2 1
connect or1.i1 and1-1.o1

define and1-2 and 2 1
connect and1-2.i1 selector.o1
connect and1-2.i2 0
connect or2.i1 and1-2.o1

define and1-3 and 2 1
connect and1-3.i1 selector.o1
connect and1-3.i2 0
connect or3.i1 and1-3.o1

define and1-4 and 2 1
connect and1-4.i1 selector.o1
connect and1-4.i2 1
connect or4.i1 and1-4.o1

define r 4bit-register
connect r.i1 or1.o1
connect r.i2 or2.o1
connect r.i3 or3.o1
connect r.i4 or4.o1
connect r.i6 $inputs.5

alias o1 r.o1
alias o2 r.o2
alias o3 r.o3
alias o4 r.o4