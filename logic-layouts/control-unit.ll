inputs 6

# inputs.1 ~ inputs.4 = rom selector
#  inputs.1 = LSB, inputs.4 = MSB
# inputs.5 = clock

define ram-alu ram-alu

define rom sample.rom

connect rom.i1 $inputs.1
connect rom.i2 $inputs.2
connect rom.i3 $inputs.3
connect rom.i4 $inputs.4
connect rom.i5 $inputs.5

connect ram-alu.i10 rom.o1
connect ram-alu.i11 rom.o2
connect ram-alu.i12 rom.o3

connect ram-alu.i13 $inputs.6

connect ram-alu.i5 $inputs.5
connect ram-alu.i6 rom.o5
connect ram-alu.i7 rom.o6
connect ram-alu.i8 rom.o7
connect ram-alu.i9 rom.o8
