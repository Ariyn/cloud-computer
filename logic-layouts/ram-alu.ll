inputs 13

# inputs.1 ~ inputs.4 = selector
#  inputs.1 = LSB, inputs.4 = MSB
# inputs.5 = clock
# inputs.6 ~ inputs.9 = alu parameter
#  inputs.6 = LSB, inputs.9 = MSB
# inputs.10 ~ inputs.11 = alu operand
#   10 is LSB, 11 is MSB.
#   00: A and B, 01: A or B, 10: not A, 11: A + B
# inputs.12 = load selector
#  load selector is 0: load from ram
#  load selector is 1: load from alu result
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

define alu 4bit-alu 10 5
connect alu.i1 $inputs.6
connect alu.i2 $inputs.7
connect alu.i3 $inputs.8
connect alu.i4 $inputs.9
# alu.i5 ~ alu.i8 connected to l
connect alu.i9 $inputs.10
connect alu.i10 $inputs.11

define demux 8x4-demux 9 4
connect demux.i1 ram.o1
connect demux.i2 ram.o2
connect demux.i3 ram.o3
connect demux.i4 ram.o4

connect demux.i5 alu.o1
connect demux.i6 alu.o2
connect demux.i7 alu.o3
connect demux.i8 alu.o4

connect demux.i9 $inputs.12

define l 4bit-register 6 4
connect l.i1 demux.o1
connect l.i2 demux.o2
connect l.i3 demux.o3
connect l.i4 demux.o4

connect l.i6 $inputs.5

define mux 4x8-mux 5 8
connect mux.i1 l.o1
connect mux.i2 l.o2
connect mux.i3 l.o3
connect mux.i4 l.o4
connect mux.i5 $inputs.13

connect ram.i1 mux.o1
connect ram.i2 mux.o2
connect ram.i3 mux.o3
connect ram.i4 mux.o4

connect alu.i5 mux.o5
connect alu.i6 mux.o6
connect alu.i7 mux.o7
connect alu.i8 mux.o8