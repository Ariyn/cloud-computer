inputs 8

define and1-1 and 2 1
define and1-2 and 2 1
define and1-3 and 2 1
define and1-4 and 2 1

connect and1-1.i1 $inputs.1
connect and1-2.i1 $inputs.2
connect and1-3.i1 $inputs.3
connect and1-4.i1 $inputs.4

connect and1-1.i2 $inputs.5
connect and1-2.i2 $inputs.5
connect and1-3.i2 $inputs.5
connect and1-4.i2 $inputs.5

# bit 2
define and2-1 and 2 1
define and2-2 and 2 1
define and2-3 and 2 1
define and2-4 and 2 1

connect and2-1.i1 $inputs.1
connect and2-2.i1 $inputs.2
connect and2-3.i1 $inputs.3
connect and2-4.i1 $inputs.4

connect and2-1.i2 $inputs.6
connect and2-2.i2 $inputs.6
connect and2-3.i2 $inputs.6
connect and2-4.i2 $inputs.6

define adder1 4bit-full-adder 9 5
connect adder1.i1 and1-2.o1
connect adder1.i2 and1-3.o1
connect adder1.i3 and1-4.o1

connect adder1.i5 and2-1.o1
connect adder1.i6 and2-2.o1
connect adder1.i7 and2-3.o1
connect adder1.i8 and2-4.o1

# bit 3
define and3-1 and 2 1
define and3-2 and 2 1
define and3-3 and 2 1
define and3-4 and 2 1

connect and3-1.i1 $inputs.1
connect and3-2.i1 $inputs.2
connect and3-3.i1 $inputs.3
connect and3-4.i1 $inputs.4

connect and3-1.i2 $inputs.7
connect and3-2.i2 $inputs.7
connect and3-3.i2 $inputs.7
connect and3-4.i2 $inputs.7

define adder2 4bit-full-adder 9 5
connect adder2.i1 adder1.o2
connect adder2.i2 adder1.o3
connect adder2.i3 adder1.o4
connect adder2.i4 adder1.carry

connect adder2.i5 and3-1.o1
connect adder2.i6 and3-2.o1
connect adder2.i7 and3-3.o1
connect adder2.i8 and3-4.o1

# bit 4
define and4-1 and 2 1
define and4-2 and 2 1
define and4-3 and 2 1
define and4-4 and 2 1

connect and4-1.i1 $inputs.1
connect and4-2.i1 $inputs.2
connect and4-3.i1 $inputs.3
connect and4-4.i1 $inputs.4

connect and4-1.i2 $inputs.8
connect and4-2.i2 $inputs.8
connect and4-3.i2 $inputs.8
connect and4-4.i2 $inputs.8

define adder3 4bit-full-adder 9 5
connect adder3.i1 adder2.o2
connect adder3.i2 adder2.o3
connect adder3.i3 adder2.o4
connect adder3.i4 adder2.carry

connect adder3.i5 and4-1.o1
connect adder3.i6 and4-2.o1
connect adder3.i7 and4-3.o1
connect adder3.i8 and4-4.o1

alias o1 and1-1.o1
alias o2 adder1.o1
alias o3 adder2.o1
alias o4 adder3.o1
alias carry adder3.o2
