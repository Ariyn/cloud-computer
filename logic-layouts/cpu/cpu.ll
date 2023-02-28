inputs 11

define BUS 4bit-bus 4 4
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

define ACC 4bit-register 6 4
connect ACC.i1 BUS.o1
connect ACC.i2 BUS.o2
connect ACC.i3 BUS.o3
connect ACC.i4 BUS.o4
connect ACC.i6 CU.o1

define ACC-OUT 4bit-and 8 4
connect ACC-OUT.i1 ACC.o1
connect ACC-OUT.i2 ACC.o2
connect ACC-OUT.i3 ACC.o3
connect ACC-OUT.i4 ACC.o4
connect ACC-OUT.i5 CU.o2
connect ACC-OUT.i6 CU.o2
connect ACC-OUT.i7 CU.o2
connect ACC-OUT.i8 CU.o2

connect BUS.i1 ACC-OUT.o1
connect BUS.i2 ACC-OUT.o2
connect BUS.i3 ACC-OUT.o3
connect BUS.i4 ACC-OUT.o4

define TEMP 4bit-register 6 4
connect TEMP.i1 BUS.o1
connect TEMP.i2 BUS.o2
connect TEMP.i3 BUS.o3
connect TEMP.i4 BUS.o4
connect TEMP.i6 CU.o3

define TEMP-OUT 4bit-and 8 4
connect TEMP-OUT.i1 TEMP.o1
connect TEMP-OUT.i2 TEMP.o2
connect TEMP-OUT.i3 TEMP.o3
connect TEMP-OUT.i4 TEMP.o4
connect TEMP-OUT.i5 CU.o4
connect TEMP-OUT.i6 CU.o4
connect TEMP-OUT.i7 CU.o4
connect TEMP-OUT.i8 CU.o4

connect BUS.i1 TEMP-OUT.o1
connect BUS.i2 TEMP-OUT.o2
connect BUS.i3 TEMP-OUT.o3
connect BUS.i4 TEMP-OUT.o4

define ALU 4bit-alu 11 5
connect ALU.i1 ACC.o1
connect ALU.i2 ACC.o2
connect ALU.i3 ACC.o3
connect ALU.i4 ACC.o4
connect ALU.i5 TEMP.o1
connect ALU.i6 TEMP.o2
connect ALU.i7 TEMP.o3
connect ALU.i8 TEMP.o4

define ALU-OUT 4bit-and 8 4
connect ALU-OUT.i1 ALU.o1
connect ALU-OUT.i2 ALU.o2
connect ALU-OUT.i3 ALU.o3
connect ALU-OUT.i4 ALU.o4
connect ALU-OUT.i5 CU.o5
connect ALU-OUT.i6 CU.o5
connect ALU-OUT.i7 CU.o5
connect ALU-OUT.i8 CU.o5

connect BUS.i1 ALU-OUT.o1
connect BUS.i2 ALU-OUT.o2
connect BUS.i3 ALU-OUT.o3
connect BUS.i4 ALU-OUT.o4

define IR 4bit-register 6 4
connect IR.i1 BUS.o1
connect IR.i2 BUS.o2
connect IR.i3 BUS.o3
connect IR.i4 BUS.o4
connect IR.i6 CU.o6

connect ALU.i9 IR.o1
connect ALU.i10 IR.o2
connect ALU.i11 IR.o3
# connect ALU.i12 IR.o4

define MAR 4bit-register 6 4
connect MAR.i1 BUS.o1
connect MAR.i2 BUS.o2
connect MAR.i3 BUS.o3
connect MAR.i4 BUS.o4
connect MAR.i6 CU.o7

define RAM 16word-ram 10 8
connect RAM.i5 MAR.o1
connect RAM.i6 MAR.o2
connect RAM.i7 MAR.o3
connect RAM.i8 MAR.o4

connect BUS.i1 RAM.o1
connect BUS.i2 RAM.o2
connect BUS.i3 RAM.o3
connect BUS.i4 RAM.o4