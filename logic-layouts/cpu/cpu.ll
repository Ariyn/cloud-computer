inputs 20

# instructions
# LMA	MOV M, A	MEM = ACC // use DataWord(bus17~20) as memory address
#    2, 7, 8 = 1, 17 ~ 20 = MEM ADDRESS
# LAM   MOV A, M    ACC = MEM // use DataWord(bus17~20) as memory address
#    3, 6, 9, 11 = 1, 17~ 20 = MEM ADDRESS (ADDRESS - 1)
# ADI   ADD A, I    ACC = ACC + Immediate Value
#    3, 6, 13, 14 = 1, 17~20 = Immediate Value
# NDI	ANI	    	A = A ∧ Immediate Value
#    3, 6, 7 = 1, 17~20 = Immediate Value

# input 1 ~ 12 = manual-cu
# ~~ i1 = acc in ~~
# i2 = acc out
# i3 = temp in
# i4 = temp out
# i5 = alu out
# i6 = ir in
# i7 = mar in
# i8 = MEM in
# i9 = MEM out

# input 12 = clock
# input 13 ~ 16 = bus opcode
#   input 13 = LSB, 16 = MSB
# input 17 ~ 20 = bus opcode
#   input 17 = LSB, 20 = MSB

# RC = Reverse Clock
define RC not 1 1
connect RC.i1 $inputs.12

# bus 1-4 = opcode (OPERATION WORD)
# bus 5-8 = data (DATA WORD)
define BUS bus/8bit 8 8

define BUS-i1 or 6 1
connect BUS-i1.i1 $inputs.13
connect BUS.i1 BUS-i1.o1

define BUS-i2 or 6 1
connect BUS-i2.i1 $inputs.14
connect BUS.i2 BUS-i2.o1

define BUS-i3 or 6 1
connect BUS-i3.i1 $inputs.15
connect BUS.i3 BUS-i3.o1

define BUS-i4 or 6 1
connect BUS-i4.i1 $inputs.16
connect BUS.i4 BUS-i4.o1

define BUS-i5 or 6 1
connect BUS-i5.i1 $inputs.17
connect BUS.i5 BUS-i5.o1

define BUS-i6 or 6 1
connect BUS-i6.i1 $inputs.18
connect BUS.i6 BUS-i6.o1

define BUS-i7 or 6 1
connect BUS-i7.i1 $inputs.19
connect BUS.i7 BUS-i7.o1

define BUS-i8 or 6 1
connect BUS-i8.i1 $inputs.20
connect BUS.i8 BUS-i8.o1

define CU cpu/manual-cu 11 11
connect CU.i1 $inputs.1
connect CU.i2 $inputs.2
connect CU.i3 $inputs.3
connect CU.i4 $inputs.4
connect CU.i5 $inputs.5
connect CU.i6 $inputs.6
connect CU.i7 $inputs.7
connect CU.i8 $inputs.8
connect CU.i9 $inputs.9
connect CU.i10 $inputs.10
connect CU.i11 $inputs.11

define ACC-FIRST 4bit-register 6 4
#connect ACC-FIRST.i1 BUS.o5
#connect ACC-FIRST.i2 BUS.o6
#connect ACC-FIRST.i3 BUS.o7
#connect ACC-FIRST.i4 BUS.o8
connect ACC-FIRST.i6 $inputs.12

define ACC-SECOND 4bit-register 6 4
connect ACC-SECOND.i1 ACC-FIRST.o1
connect ACC-SECOND.i2 ACC-FIRST.o2
connect ACC-SECOND.i3 ACC-FIRST.o3
connect ACC-SECOND.i4 ACC-FIRST.o4
connect ACC-SECOND.i6 RC.o1

define ACC-OUT 4bit-and 8 4
connect ACC-OUT.i1 ACC-SECOND.o1
connect ACC-OUT.i2 ACC-SECOND.o2
connect ACC-OUT.i3 ACC-SECOND.o3
connect ACC-OUT.i4 ACC-SECOND.o4
connect ACC-OUT.i5 CU.o2
connect ACC-OUT.i6 CU.o2
connect ACC-OUT.i7 CU.o2
connect ACC-OUT.i8 CU.o2

#connect BUS-i5.i6 ACC-OUT.o1
#connect BUS-i6.i6 ACC-OUT.o2
#connect BUS-i7.i6 ACC-OUT.o3
#connect BUS-i8.i6 ACC-OUT.o4

define TEMP complex/4bit-rising-edge-enable-register 7 4
connect TEMP.i1 BUS.o5
connect TEMP.i2 BUS.o6
connect TEMP.i3 BUS.o7
connect TEMP.i4 BUS.o8
connect TEMP.i6 $inputs.12
connect TEMP.i7 CU.o3

define TEMP-OUT 4bit-and 8 4
connect TEMP-OUT.i1 TEMP.o1
connect TEMP-OUT.i2 TEMP.o2
connect TEMP-OUT.i3 TEMP.o3
connect TEMP-OUT.i4 TEMP.o4
connect TEMP-OUT.i5 CU.o4
connect TEMP-OUT.i6 CU.o4
connect TEMP-OUT.i7 CU.o4
connect TEMP-OUT.i8 CU.o4

connect BUS-i5.i2 TEMP-OUT.o1
connect BUS-i6.i2 TEMP-OUT.o2
connect BUS-i7.i2 TEMP-OUT.o3
connect BUS-i8.i2 TEMP-OUT.o4

define ALU 4bit-alu 11 5
connect ALU.i1 TEMP.o1
connect ALU.i2 TEMP.o2
connect ALU.i3 TEMP.o3
connect ALU.i4 TEMP.o4
connect ALU.i5 ACC-SECOND.o1
connect ALU.i6 ACC-SECOND.o2
connect ALU.i7 ACC-SECOND.o3
connect ALU.i8 ACC-SECOND.o4

connect ACC-FIRST.i1 ALU.o1
connect ACC-FIRST.i2 ALU.o2
connect ACC-FIRST.i3 ALU.o3
connect ACC-FIRST.i4 ALU.o4

#define ALU-OUT 4bit-and 8 4
#connect ALU-OUT.i1 ALU.o1
#connect ALU-OUT.i2 ALU.o2
#connect ALU-OUT.i3 ALU.o3
#connect ALU-OUT.i4 ALU.o4
#connect ALU-OUT.i5 CU.o5
#connect ALU-OUT.i6 CU.o5
#connect ALU-OUT.i7 CU.o5
#connect ALU-OUT.i8 CU.o5

#connect BUS-i5.i3 ALU-OUT.o1
#connect BUS-i2.i3 ALU-OUT.o2
#connect BUS-i3.i3 ALU-OUT.o3
#connect BUS-i4.i3 ALU-OUT.o4

define IR complex/4bit-rising-edge-enable-register 7 4
connect IR.i1 BUS.o1
connect IR.i2 BUS.o2
connect IR.i3 BUS.o3
connect IR.i4 BUS.o4
connect IR.i6 $inputs.12
connect IR.i7 CU.o6

connect ALU.i9 IR.o1
connect ALU.i10 IR.o2
connect ALU.i11 IR.o3
# connect ALU.i12 IR.o4

define MAR complex/4bit-rising-edge-enable-register 7 4
connect MAR.i1 BUS.o5
connect MAR.i2 BUS.o6
connect MAR.i3 BUS.o7
connect MAR.i4 BUS.o8
connect MAR.i6 $inputs.12
connect MAR.i7 CU.o7

define RAM 16word-ram 10 8
connect RAM.i1 ACC-OUT.o1
connect RAM.i2 ACC-OUT.o2
connect RAM.i3 ACC-OUT.o3
connect RAM.i4 ACC-OUT.o4
connect RAM.i5 MAR.o1
connect RAM.i6 MAR.o2
connect RAM.i7 MAR.o3
connect RAM.i8 MAR.o4

define ram-trigger complex/rising-edge-trigger 1 1
connect ram-trigger.i1 RC.o1

define ram-enabled-clock and 2 1
connect ram-enabled-clock.i1 ram-trigger.o1
connect ram-enabled-clock.i2 CU.o8
connect RAM.i10 ram-enabled-clock.o1


define RAM-OUT 8bit-and 16 8
connect RAM-OUT.i1 RAM.o1
connect RAM-OUT.i2 RAM.o2
connect RAM-OUT.i3 RAM.o3
connect RAM-OUT.i4 RAM.o4
connect RAM-OUT.i5 RAM.o5
connect RAM-OUT.i6 RAM.o6
connect RAM-OUT.i7 RAM.o7
connect RAM-OUT.i8 RAM.o8
connect RAM-OUT.i9 CU.o9
connect RAM-OUT.i10 CU.o9
connect RAM-OUT.i11 CU.o9
connect RAM-OUT.i12 CU.o9
connect RAM-OUT.i13 CU.o9
connect RAM-OUT.i14 CU.o9
connect RAM-OUT.i15 CU.o9
connect RAM-OUT.i16 CU.o9

# bus를 MEM의 Address 값과 MEM에서 꺼내오는 값, 두 개를 동시에 사용하다보니 currupt가 일어남
# single cycle cpu를 구현하려면, BUS가 분리되어야 한다.
# Address register를 새롭게 만들까?
connect BUS-i1.i5 RAM-OUT.o1
connect BUS-i2.i5 RAM-OUT.o2
connect BUS-i3.i5 RAM-OUT.o3
connect BUS-i4.i5 RAM-OUT.o4
connect BUS-i5.i5 RAM-OUT.o5
connect BUS-i6.i5 RAM-OUT.o6
connect BUS-i7.i5 RAM-OUT.o7
connect BUS-i8.i5 RAM-OUT.o8