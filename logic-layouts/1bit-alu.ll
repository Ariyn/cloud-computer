# https://www.101computing.net/8-bit-alu-using-logic-gates/
# 1bit alu has 4 methods. not a, a or b, a and b, a + b
# there should be 2 selector bits, 1 bit for each a and b, 1 bit for carry
# there would be 2 outputs for output and carry out.
inputs 6

# inputs.4 ~ 6 = opcode
#  inputs.4 = LSB, 6 = MSB
# opcode 0 = A and B
# opcode 1 = A or B
# opcode 2 = not A
# opcode 3 = A + B + (if exists) carry
# opcode 4 = load A

define result or 8 1

define selector 3bit-decoder 3 8
connect selector.i1 $inputs.4
connect selector.i2 $inputs.5
connect selector.i3 $inputs.6

# and operation opcode 0
# will return A and B
define and and 2 1
connect and.i1 $inputs.1
connect and.i2 $inputs.2

define and-output and 2 1
connect and-output.i1 and.o1
connect and-output.i2 selector.o1

connect result.i1 and-output.o1

# or operation opcode 1
# will return A or B

define or or 2 1
connect or.i1 $inputs.1
connect or.i2 $inputs.2

define or-output and 2 1
connect or-output.i1 or.o1
connect or-output.i2 selector.o2

connect result.i2 or-output.o1

# not operation opcode 2
# will return not A
define not not 1 1
connect not.i1 $inputs.1

define not-output and 2 1
connect not-output.i1 not.o1
connect not-output.i2 selector.o3

alias output not-output.o1
connect result.i3 not-output.o1

# add operation opcode 3
# will return a + b + (if exists) carry
define adder full-adder 3 2
connect adder.i1 $inputs.1
connect adder.i2 $inputs.2
connect adder.i3 $inputs.3

define adder-sum-output and 2 1
connect adder-sum-output.i1 adder.sum
connect adder-sum-output.i2 selector.o4

connect result.i4 adder-sum-output.o1

define adder-carry-output and 2 1
connect adder-carry-output.i1 adder.carry
connect adder-carry-output.i2 selector.o4

alias carry adder-carry-output.o1

define loader and 2 1
connect loader.i1 $inputs.1
connect loader.i2 selector.o5

connect result.i5 loader.o1

alias output result.o1