inputs 3

# a - b
# inputs.1 = a
# inputs.2 = b
# inputs.3 = loan in

# output d = difference
# output lo = loan out

define xor1 xor 2 1
define xor2 xor 2 1

connect xor1.i1 $inputs.1
connect xor1.i2 $inputs.2

connect xor2.i1 xor1.o1
connect xor2.i2 $inputs.3

alias d xor2.o1

define not1 not 1 1
define and1 and 2 1

# loan_out = not a & b | loan_in & !( a xor b )
connect not1.i1 $inputs.1
connect and1.i1 not1.o1
connect and1.i2 $inputs.2

define not2 not 1 1
connect not2.i1 xor1.o1

define and2 and 2 1
connect and2.i1 not2.o1
connect and2.i2 $inputs.3

define or1 or 2 1
connect or1.i1 and1.o1
connect or1.i2 and2.o1

alias lo or1.o1