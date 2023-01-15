inputs 9

# inputs.1 ~ inputs.4 = opcode # inputs.4 is not implemented for now
#  inputs.1 = LSB, inputs.4 = MSB
# inputs.4 ~ inputs.8 = operand
#  inputs.4 = LSB, inputs.8 = MSB
# inputs.9 = clock

# opcs = opcode-selector
define opcs 3bit-decoder

# opcode
connect opcs.i1 $inputs.1
connect opcs.i2 $inputs.2
connect opcs.i3 $inputs.3

define ram-alu ram-alu

define op1 or 8 1
define op2 or 8 1
define op3 or 8 1

connect ram-alu.i10 $inputs.1
connect ram-alu.i11 $inputs.2
connect ram-alu.i12 $inputs.3


# opcode 011 = add operand from accumulator and store to accumulator

define op4not not 1 1
connect op4not.i1 opcs.o4

connect op1.i4 opcs.o4
connect op2.i4 opcs.o4
connect op3.i4 opcs.o4

connect ram-alu.i1 $inputs.5
connect ram-alu.i2 $inputs.6
connect ram-alu.i3 $inputs.7
connect ram-alu.i4 $inputs.8
connect ram-alu.i5 $inputs.9


# opcode 100 = store L register to ram {operand} address

define op5not not 1 1
connect op5not.i1 opcs.o5

connect op1.i5 op5not.o1
connect op2.i5 op5not.o1
connect op3.i5 opcs.

connect ram-alu.i1 $inputs.5
connect ram-alu.i2 $inputs.6
connect ram-alu.i3 $inputs.7
connect ram-alu.i4 $inputs.8
connect ram-alu.i5 $inputs.9
connect ram-alu.i13
