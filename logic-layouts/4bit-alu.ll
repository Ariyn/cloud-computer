inputs 10

# inputs.9 ~ inputs.10 = opcode
# inputs.9 = LSB, inputs.10 = MSB
# opcode 0 = A and B
# opcode 1 = A or B
# opcode 2 = not A
# opcode 3 = A + B + (if exists) carry

define alu1 1bit-alu 5 2
connect alu1.i1 $inputs.1
connect alu1.i2 $inputs.5
connect alu1.i3 0
connect alu1.i4 $inputs.9
connect alu1.i5 $inputs.10

define alu2 1bit-alu 5 2
connect alu2.i1 $inputs.2
connect alu2.i2 $inputs.6
connect alu2.i3 alu1.carry
connect alu2.i4 $inputs.9
connect alu2.i5 $inputs.10

define alu3 1bit-alu 5 2
connect alu3.i1 $inputs.3
connect alu3.i2 $inputs.7
connect alu3.i3 alu2.carry
connect alu3.i4 $inputs.9
connect alu3.i5 $inputs.10

define alu4 1bit-alu 5 2
connect alu4.i1 $inputs.4
connect alu4.i2 $inputs.8
connect alu4.i3 alu3.carry
connect alu4.i4 $inputs.9
connect alu4.i5 $inputs.10

alias o1 alu1.output
alias o2 alu2.output
alias o3 alu3.output
alias o4 alu4.output
alias o5 alu4.carry
