inputs 13

# inputs.1 ~ inputs.4 = selector
#  inputs.1 = LSB, inputs.4 = MSB
# inputs.5 = clock
# inputs.6 ~ inputs.9 = alu parameter
#  inputs.6 = LSB, inputs.9 = MSB
# inputs.10 ~ inputs.12 = alu operand
#   10 is LSB, 12 is MSB.
#   0: A and B, 1: A or B, 2: not A, 3: A + B, 4: LOAD A
# inputs.13 = destination selector
#  destination selector is 0: destination to ram
#  destination selector is 1: destination to alu inputs b

define ram 16word-ram 10 4
connect ram.i5 $inputs.1
connect ram.i6 $inputs.2
connect ram.i7 $inputs.3
connect ram.i8 $inputs.4

define ramClock and 2 1
connect ramClock.i1 $inputs.5

define ramClockNot not 1 1
connect ramClockNot.i1 $inputs.13
connect ramClock.i2 ramClockNot.o1

connect ram.i10 ramClock.o1

define alu 4bit-alu 11 5
connect alu.i1 $inputs.6
connect alu.i2 $inputs.7
connect alu.i3 $inputs.8
connect alu.i4 $inputs.9
# alu.i5 ~ alu.i8 connected to l
connect alu.i9 $inputs.10
connect alu.i10 $inputs.11
connect alu.i11 $inputs.12

define accm 4bit-register 6 4
connect accm.i1 alu.o1
connect accm.i2 alu.o2
connect accm.i3 alu.o3
connect accm.i4 alu.o4
connect accm.i6 $inputs.5

define accm2 4bit-register 6 4
connect accm2.i1 accm.o1
connect accm2.i2 accm.o2
connect accm2.i3 accm.o3
connect accm2.i4 accm.o4

define clockNot not 1 1
connect clockNot.i1 $inputs.5
connect accm2.i6 clockNot.o1

connect ram.i1 accm2.o1
connect ram.i2 accm2.o2
connect ram.i3 accm2.o3
connect ram.i4 accm2.o4

connect alu.i5 accm2.o1
connect alu.i6 accm2.o2
connect alu.i7 accm2.o3
connect alu.i8 accm2.o4