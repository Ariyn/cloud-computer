inputs 66

#  inputs.5 ~ inputs.32 is not implemented yet.
#  inputs.37 ~ inputs.64 is not implemented yet.

# inputs.1 ~ inputs.32 = inputs bits
#  inputs.1 is LSB, inputs.32 is MSB
# inputs.33 ~ inputs.64 = selector
#  inputs.33 is LSB, inputs.64 is MSB
# inputs.65 = reset
# inputs.66 clk
#  clock will used for write.
#  output will be selected immediately after selector input.

define out1 or 16 1
alias o1 out1.o1

define out2 or 16 1
alias o2 out2.o1

define out3 or 16 1
alias o3 out3.o1

define out4 or 16 1
alias o4 out4.o1

define out5 or 16 1
alias o5 out5.o1

define out6 or 16 1
alias o6 out6.o1

define out7 or 16 1
alias o7 out7.o1

define out8 or 16 1
alias o8 out8.o1

define selector 4bit-decoder 4 16
connect selector.i1 $inputs.33
connect selector.i2 $inputs.34
connect selector.i3 $inputs.35
connect selector.i4 $inputs.36

# ram address 1
define reg1 4bit-register 6 4
connect reg1.i1 $inputs.1
connect reg1.i2 $inputs.2
connect reg1.i3 $inputs.3
connect reg1.i4 $inputs.4
connect reg1.i5 $inputs.65

define and1 and 2 1
connect and1.i1 $inputs.66
connect and1.i2 selector.o1
connect reg1.i6 and1.o1

define o1-1 and 2 1
connect o1-1.i1 selector.o1
connect o1-1.i2 reg1.o1

connect out1.i1 o1-1.o1

define o1-2 and 2 1
connect o1-2.i1 selector.o1
connect o1-2.i2 reg1.o2

connect out2.i1 o1-2.o1

define o1-3 and 2 1
connect o1-3.i1 selector.o1
connect o1-3.i2 reg1.o3

connect out3.i1 o1-3.o1

define o1-4 and 2 1
connect o1-4.i1 selector.o1
connect o1-4.i2 reg1.o4

connect out4.i1 o1-4.o1

# ram address 2
define reg2 4bit-register 6 4
connect reg2.i1 $inputs.1
connect reg2.i2 $inputs.2
connect reg2.i3 $inputs.3
connect reg2.i4 $inputs.4
connect reg2.i5 $inputs.65

define and2 and 2 1
connect and2.i1 $inputs.66
connect and2.i2 selector.o2
connect reg2.i6 and2.o1

define o2-1 and 2 1
connect o2-1.i1 selector.o2
connect o2-1.i2 reg2.o1

connect out1.i2 o2-1.o1

define o2-2 and 2 1
connect o2-2.i1 selector.o2
connect o2-2.i2 reg2.o2

connect out2.i2 o2-2.o1

define o2-3 and 2 1
connect o2-3.i1 selector.o2
connect o2-3.i2 reg2.o3

connect out3.i2 o2-3.o1

define o2-4 and 2 1
connect o2-4.i1 selector.o2
connect o2-4.i2 reg2.o4

connect out4.i2 o2-4.o1

# ram address 3
define reg3 4bit-register 6 4
connect reg3.i1 $inputs.1
connect reg3.i2 $inputs.2
connect reg3.i3 $inputs.3
connect reg3.i4 $inputs.4
connect reg3.i5 $inputs.65

define and3 and 2 1
connect and3.i1 $inputs.66
connect and3.i2 selector.o3
connect reg3.i6 and3.o1

define o3-1 and 2 1
connect o3-1.i1 selector.o3
connect o3-1.i2 reg3.o1

connect out1.i3 o3-1.o1

define o3-2 and 2 1
connect o3-2.i1 selector.o3
connect o3-2.i2 reg3.o2

connect out2.i3 o3-2.o1

define o3-3 and 2 1
connect o3-3.i1 selector.o3
connect o3-3.i2 reg3.o3

connect out3.i3 o3-3.o1

define o3-4 and 2 1
connect o3-4.i1 selector.o3
connect o3-4.i2 reg3.o4

connect out4.i3 o3-4.o1

# ram address 4
define reg4 4bit-register 6 4
connect reg4.i1 $inputs.1
connect reg4.i2 $inputs.2
connect reg4.i3 $inputs.3
connect reg4.i4 $inputs.4
connect reg4.i5 $inputs.65

define and4 and 2 1
connect and4.i1 $inputs.66
connect and4.i2 selector.o4
connect reg4.i6 and4.o1

define o4-1 and 2 1
connect o4-1.i1 selector.o4
connect o4-1.i2 reg4.o1

connect out1.i4 o4-1.o1

define o4-2 and 2 1
connect o4-2.i1 selector.o4
connect o4-2.i2 reg4.o2

connect out2.i4 o4-2.o1

define o4-3 and 2 1
connect o4-3.i1 selector.o4
connect o4-3.i2 reg4.o3

connect out3.i4 o4-3.o1

define o4-4 and 2 1
connect o4-4.i1 selector.o4
connect o4-4.i2 reg4.o4

connect out4.i4 o4-4.o1

# ram address 5
define reg5 4bit-register 6 4
connect reg5.i1 $inputs.1
connect reg5.i2 $inputs.2
connect reg5.i3 $inputs.3
connect reg5.i4 $inputs.4
connect reg5.i5 $inputs.65

define and5 and 2 1
connect and5.i1 $inputs.66
connect and5.i2 selector.o5
connect reg5.i6 and5.o1

define o5-1 and 2 1
connect o5-1.i1 selector.o5
connect o5-1.i2 reg5.o1

connect out1.i5 o5-1.o1

define o5-2 and 2 1
connect o5-2.i1 selector.o5
connect o5-2.i2 reg5.o2

connect out2.i5 o5-2.o1

define o5-3 and 2 1
connect o5-3.i1 selector.o5
connect o5-3.i2 reg5.o3

connect out3.i5 o5-3.o1

define o5-4 and 2 1
connect o5-4.i1 selector.o5
connect o5-4.i2 reg5.o4

connect out4.i5 o5-4.o1

# ram address 6
define reg6 4bit-register 6 4
connect reg6.i1 $inputs.1
connect reg6.i2 $inputs.2
connect reg6.i3 $inputs.3
connect reg6.i4 $inputs.4
connect reg6.i5 $inputs.65

define and6 and 2 1
connect and6.i1 $inputs.66
connect and6.i2 selector.o6
connect reg6.i6 and6.o1

define o6-1 and 2 1
connect o6-1.i1 selector.o6
connect o6-1.i2 reg6.o1

connect out1.i6 o6-1.o1

define o6-2 and 2 1
connect o6-2.i1 selector.o6
connect o6-2.i2 reg6.o2

connect out2.i6 o6-2.o1

define o6-3 and 2 1
connect o6-3.i1 selector.o6
connect o6-3.i2 reg6.o3

connect out3.i6 o6-3.o1

define o6-4 and 2 1
connect o6-4.i1 selector.o6
connect o6-4.i2 reg6.o4

connect out4.i6 o6-4.o1

# ram address 7
define reg7 4bit-register 6 4
connect reg7.i1 $inputs.1
connect reg7.i2 $inputs.2
connect reg7.i3 $inputs.3
connect reg7.i4 $inputs.4
connect reg7.i5 $inputs.65

define and7 and 2 1
connect and7.i1 $inputs.66
connect and7.i2 selector.o7
connect reg7.i6 and7.o1

define o7-1 and 2 1
connect o7-1.i1 selector.o7
connect o7-1.i2 reg7.o1

connect out1.i7 o7-1.o1

define o7-2 and 2 1
connect o7-2.i1 selector.o7
connect o7-2.i2 reg7.o2

connect out2.i7 o7-2.o1

define o7-3 and 2 1
connect o7-3.i1 selector.o7
connect o7-3.i2 reg7.o3

connect out3.i7 o7-3.o1

define o7-4 and 2 1
connect o7-4.i1 selector.o7
connect o7-4.i2 reg7.o4

connect out4.i7 o7-4.o1

# ram address 8
define reg8 4bit-register 6 4
connect reg8.i1 $inputs.1
connect reg8.i2 $inputs.2
connect reg8.i3 $inputs.3
connect reg8.i4 $inputs.4
connect reg8.i5 $inputs.65

define and8 and 2 1
connect and8.i1 $inputs.66
connect and8.i2 selector.o8
connect reg8.i6 and8.o1

define o8-1 and 2 1
connect o8-1.i1 selector.o8
connect o8-1.i2 reg8.o1

connect out1.i8 o8-1.o1

define o8-2 and 2 1
connect o8-2.i1 selector.o8
connect o8-2.i2 reg8.o2

connect out2.i8 o8-2.o1

define o8-3 and 2 1
connect o8-3.i1 selector.o8
connect o8-3.i2 reg8.o3

connect out3.i8 o8-3.o1

define o8-4 and 2 1
connect o8-4.i1 selector.o8
connect o8-4.i2 reg8.o4

connect out4.i8 o8-4.o1

# ram address 9
define reg9 4bit-register 6 4
connect reg9.i1 $inputs.1
connect reg9.i2 $inputs.2
connect reg9.i3 $inputs.3
connect reg9.i4 $inputs.4
connect reg9.i5 $inputs.65

define and9 and 2 1
connect and9.i1 $inputs.66
connect and9.i2 selector.o9
connect reg9.i6 and9.o1

define o9-1 and 2 1
connect o9-1.i1 selector.o9
connect o9-1.i2 reg9.o1

connect out1.i9 o9-1.o1

define o9-2 and 2 1
connect o9-2.i1 selector.o9
connect o9-2.i2 reg9.o2

connect out2.i9 o9-2.o1

define o9-3 and 2 1
connect o9-3.i1 selector.o9
connect o9-3.i2 reg9.o3

connect out3.i9 o9-3.o1

define o9-4 and 2 1
connect o9-4.i1 selector.o9
connect o9-4.i2 reg9.o4

connect out4.i9 o9-4.o1

# ram address 10
define reg10 4bit-register 6 4
connect reg10.i1 $inputs.1
connect reg10.i2 $inputs.2
connect reg10.i3 $inputs.3
connect reg10.i4 $inputs.4
connect reg10.i5 $inputs.65

define and10 and 2 1
connect and10.i1 $inputs.66
connect and10.i2 selector.o10
connect reg10.i6 and10.o1

define o10-1 and 2 1
connect o10-1.i1 selector.o10
connect o10-1.i2 reg10.o1

connect out1.i10 o10-1.o1

define o10-2 and 2 1
connect o10-2.i1 selector.o10
connect o10-2.i2 reg10.o2

connect out2.i10 o10-2.o1

define o10-3 and 2 1
connect o10-3.i1 selector.o10
connect o10-3.i2 reg10.o3

connect out3.i10 o10-3.o1

define o10-4 and 2 1
connect o10-4.i1 selector.o10
connect o10-4.i2 reg10.o4

connect out4.i10 o10-4.o1

# ram address 11
define reg11 4bit-register 6 4
connect reg11.i1 $inputs.1
connect reg11.i2 $inputs.2
connect reg11.i3 $inputs.3
connect reg11.i4 $inputs.4
connect reg11.i5 $inputs.65

define and11 and 2 1
connect and11.i1 $inputs.66
connect and11.i2 selector.o11
connect reg11.i6 and11.o1

define o11-1 and 2 1
connect o11-1.i1 selector.o11
connect o11-1.i2 reg11.o1

connect out1.i11 o11-1.o1

define o11-2 and 2 1
connect o11-2.i1 selector.o11
connect o11-2.i2 reg11.o2

connect out2.i11 o11-2.o1

define o11-3 and 2 1
connect o11-3.i1 selector.o11
connect o11-3.i2 reg11.o3

connect out3.i11 o11-3.o1

define o11-4 and 2 1
connect o11-4.i1 selector.o11
connect o11-4.i2 reg11.o4

connect out4.i11 o11-4.o1

# ram address 12
define reg12 4bit-register 6 4
connect reg12.i1 $inputs.1
connect reg12.i2 $inputs.2
connect reg12.i3 $inputs.3
connect reg12.i4 $inputs.4
connect reg12.i5 $inputs.65

define and12 and 2 1
connect and12.i1 $inputs.66
connect and12.i2 selector.o12
connect reg12.i6 and12.o1

define o12-1 and 2 1
connect o12-1.i1 selector.o12
connect o12-1.i2 reg12.o1

connect out1.i12 o12-1.o1

define o12-2 and 2 1
connect o12-2.i1 selector.o12
connect o12-2.i2 reg12.o2

connect out2.i12 o12-2.o1

define o12-3 and 2 1
connect o12-3.i1 selector.o12
connect o12-3.i2 reg12.o3

connect out3.i12 o12-3.o1

define o12-4 and 2 1
connect o12-4.i1 selector.o12
connect o12-4.i2 reg12.o4

connect out4.i12 o12-4.o1

# ram address 13
define reg13 4bit-register 6 4
connect reg13.i1 $inputs.1
connect reg13.i2 $inputs.2
connect reg13.i3 $inputs.3
connect reg13.i4 $inputs.4
connect reg13.i5 $inputs.65

define and13 and 2 1
connect and13.i1 $inputs.66
connect and13.i2 selector.o13
connect reg13.i6 and13.o1

define o13-1 and 2 1
connect o13-1.i1 selector.o13
connect o13-1.i2 reg13.o1

connect out1.i13 o13-1.o1

define o13-2 and 2 1
connect o13-2.i1 selector.o13
connect o13-2.i2 reg13.o2

connect out2.i13 o13-2.o1

define o13-3 and 2 1
connect o13-3.i1 selector.o13
connect o13-3.i2 reg13.o3

connect out3.i13 o13-3.o1

define o13-4 and 2 1
connect o13-4.i1 selector.o13
connect o13-4.i2 reg13.o4

connect out4.i13 o13-4.o1

# ram address 14
define reg14 4bit-register 6 4
connect reg14.i1 $inputs.1
connect reg14.i2 $inputs.2
connect reg14.i3 $inputs.3
connect reg14.i4 $inputs.4
connect reg14.i5 $inputs.65

define and14 and 2 1
connect and14.i1 $inputs.66
connect and14.i2 selector.o14
connect reg14.i6 and14.o1

define o14-1 and 2 1
connect o14-1.i1 selector.o14
connect o14-1.i2 reg14.o1

connect out1.i14 o14-1.o1

define o14-2 and 2 1
connect o14-2.i1 selector.o14
connect o14-2.i2 reg14.o2

connect out2.i14 o14-2.o1

define o14-3 and 2 1
connect o14-3.i1 selector.o14
connect o14-3.i2 reg14.o3

connect out3.i14 o14-3.o1

define o14-4 and 2 1
connect o14-4.i1 selector.o14
connect o14-4.i2 reg14.o4

connect out4.i14 o14-4.o1

# ram address 15
define reg15 4bit-register 6 4
connect reg15.i1 $inputs.1
connect reg15.i2 $inputs.2
connect reg15.i3 $inputs.3
connect reg15.i4 $inputs.4
connect reg15.i5 $inputs.65

define and15 and 2 1
connect and15.i1 $inputs.66
connect and15.i2 selector.o15
connect reg15.i6 and15.o1

define o15-1 and 2 1
connect o15-1.i1 selector.o15
connect o15-1.i2 reg15.o1

connect out1.i15 o15-1.o1

define o15-2 and 2 1
connect o15-2.i1 selector.o15
connect o15-2.i2 reg15.o2

connect out2.i15 o15-2.o1

define o15-3 and 2 1
connect o15-3.i1 selector.o15
connect o15-3.i2 reg15.o3

connect out3.i15 o15-3.o1

define o15-4 and 2 1
connect o15-4.i1 selector.o15
connect o15-4.i2 reg15.o4

connect out4.i15 o15-4.o1

# ram address 16
define reg16 4bit-register 6 4
connect reg16.i1 $inputs.1
connect reg16.i2 $inputs.2
connect reg16.i3 $inputs.3
connect reg16.i4 $inputs.4
connect reg16.i5 $inputs.65

define and16 and 2 1
connect and16.i1 $inputs.66
connect and16.i2 selector.o16
connect reg16.i6 and16.o1

define o16-1 and 2 1
connect o16-1.i1 selector.o16
connect o16-1.i2 reg16.o1

connect out1.i16 o16-1.o1

define o16-2 and 2 1
connect o16-2.i1 selector.o16
connect o16-2.i2 reg16.o2

connect out2.i16 o16-2.o1

define o16-3 and 2 1
connect o16-3.i1 selector.o16
connect o16-3.i2 reg16.o3

connect out3.i16 o16-3.o1

define o16-4 and 2 1
connect o16-4.i1 selector.o16
connect o16-4.i2 reg16.o4

connect out4.i16 o16-4.o1

define sa 4bit-adder-subtractor 9 1
connect sa.i1 $inputs.5
connect sa.i2 $inputs.6
connect sa.i3 $inputs.7
connect sa.i4 $inputs.8
connect sa.i5 1
connect sa.i6 0
connect sa.i7 0
connect sa.i8 0
connect sa.i9 0

define selector2 4bit-decoder 4 16
connect selector2.i1 sa.o1
connect selector2.i2 sa.o2
connect selector2.i3 sa.o3
connect selector2.i4 sa.o4
define o1-5 and 2 1
connect o1-5.i1 selector2.o1
connect o1-5.i2 reg1.o1

connect out5.i1 o1-5.o1

define o1-6 and 2 1
connect o1-6.i1 selector2.o1
connect o1-6.i2 reg1.o2

connect out6.i1 o1-6.o1

define o1-7 and 2 1
connect o1-7.i1 selector2.o1
connect o1-7.i2 reg1.o3

connect out7.i1 o1-7.o1

define o1-8 and 2 1
connect o1-8.i1 selector2.o1
connect o1-8.i2 reg1.o4

connect out8.i1 o1-8.o1


define o2-5 and 2 1
connect o2-5.i1 selector2.o2
connect o2-5.i2 reg2.o1

connect out5.i2 o2-5.o1

define o2-6 and 2 1
connect o2-6.i1 selector2.o2
connect o2-6.i2 reg2.o2

connect out6.i2 o2-6.o1

define o2-7 and 2 1
connect o2-7.i1 selector2.o2
connect o2-7.i2 reg2.o3

connect out7.i2 o2-7.o1

define o2-8 and 2 1
connect o2-8.i1 selector2.o2
connect o2-8.i2 reg2.o4

connect out8.i2 o2-8.o1


define o3-5 and 2 1
connect o3-5.i1 selector2.o3
connect o3-5.i2 reg3.o1

connect out5.i3 o3-5.o1

define o3-6 and 2 1
connect o3-6.i1 selector2.o3
connect o3-6.i2 reg3.o2

connect out6.i3 o3-6.o1

define o3-7 and 2 1
connect o3-7.i1 selector2.o3
connect o3-7.i2 reg3.o3

connect out7.i3 o3-7.o1

define o3-8 and 2 1
connect o3-8.i1 selector2.o3
connect o3-8.i2 reg3.o4

connect out8.i3 o3-8.o1


define o4-5 and 2 1
connect o4-5.i1 selector2.o4
connect o4-5.i2 reg4.o1

connect out5.i4 o4-5.o1

define o4-6 and 2 1
connect o4-6.i1 selector2.o4
connect o4-6.i2 reg4.o2

connect out6.i4 o4-6.o1

define o4-7 and 2 1
connect o4-7.i1 selector2.o4
connect o4-7.i2 reg4.o3

connect out7.i4 o4-7.o1

define o4-8 and 2 1
connect o4-8.i1 selector2.o4
connect o4-8.i2 reg4.o4

connect out8.i4 o4-8.o1


define o5-5 and 2 1
connect o5-5.i1 selector2.o5
connect o5-5.i2 reg5.o1

connect out5.i5 o5-5.o1

define o5-6 and 2 1
connect o5-6.i1 selector2.o5
connect o5-6.i2 reg5.o2

connect out6.i5 o5-6.o1

define o5-7 and 2 1
connect o5-7.i1 selector2.o5
connect o5-7.i2 reg5.o3

connect out7.i5 o5-7.o1

define o5-8 and 2 1
connect o5-8.i1 selector2.o5
connect o5-8.i2 reg5.o4

connect out8.i5 o5-8.o1


define o6-5 and 2 1
connect o6-5.i1 selector2.o6
connect o6-5.i2 reg6.o1

connect out5.i6 o6-5.o1

define o6-6 and 2 1
connect o6-6.i1 selector2.o6
connect o6-6.i2 reg6.o2

connect out6.i6 o6-6.o1

define o6-7 and 2 1
connect o6-7.i1 selector2.o6
connect o6-7.i2 reg6.o3

connect out7.i6 o6-7.o1

define o6-8 and 2 1
connect o6-8.i1 selector2.o6
connect o6-8.i2 reg6.o4

connect out8.i6 o6-8.o1


define o7-5 and 2 1
connect o7-5.i1 selector2.o7
connect o7-5.i2 reg7.o1

connect out5.i7 o7-5.o1

define o7-6 and 2 1
connect o7-6.i1 selector2.o7
connect o7-6.i2 reg7.o2

connect out6.i7 o7-6.o1

define o7-7 and 2 1
connect o7-7.i1 selector2.o7
connect o7-7.i2 reg7.o3

connect out7.i7 o7-7.o1

define o7-8 and 2 1
connect o7-8.i1 selector2.o7
connect o7-8.i2 reg7.o4

connect out8.i7 o7-8.o1


define o8-5 and 2 1
connect o8-5.i1 selector2.o8
connect o8-5.i2 reg8.o1

connect out5.i8 o8-5.o1

define o8-6 and 2 1
connect o8-6.i1 selector2.o8
connect o8-6.i2 reg8.o2

connect out6.i8 o8-6.o1

define o8-7 and 2 1
connect o8-7.i1 selector2.o8
connect o8-7.i2 reg8.o3

connect out7.i8 o8-7.o1

define o8-8 and 2 1
connect o8-8.i1 selector2.o8
connect o8-8.i2 reg8.o4

connect out8.i8 o8-8.o1


define o9-5 and 2 1
connect o9-5.i1 selector2.o9
connect o9-5.i2 reg9.o1

connect out5.i9 o9-5.o1

define o9-6 and 2 1
connect o9-6.i1 selector2.o9
connect o9-6.i2 reg9.o2

connect out6.i9 o9-6.o1

define o9-7 and 2 1
connect o9-7.i1 selector2.o9
connect o9-7.i2 reg9.o3

connect out7.i9 o9-7.o1

define o9-8 and 2 1
connect o9-8.i1 selector2.o9
connect o9-8.i2 reg9.o4

connect out8.i9 o9-8.o1


define o10-5 and 2 1
connect o10-5.i1 selector2.o10
connect o10-5.i2 reg10.o1

connect out5.i10 o10-5.o1

define o10-6 and 2 1
connect o10-6.i1 selector2.o10
connect o10-6.i2 reg10.o2

connect out6.i10 o10-6.o1

define o10-7 and 2 1
connect o10-7.i1 selector2.o10
connect o10-7.i2 reg10.o3

connect out7.i10 o10-7.o1

define o10-8 and 2 1
connect o10-8.i1 selector2.o10
connect o10-8.i2 reg10.o4

connect out8.i10 o10-8.o1


define o11-5 and 2 1
connect o11-5.i1 selector2.o11
connect o11-5.i2 reg11.o1

connect out5.i11 o11-5.o1

define o11-6 and 2 1
connect o11-6.i1 selector2.o11
connect o11-6.i2 reg11.o2

connect out6.i11 o11-6.o1

define o11-7 and 2 1
connect o11-7.i1 selector2.o11
connect o11-7.i2 reg11.o3

connect out7.i11 o11-7.o1

define o11-8 and 2 1
connect o11-8.i1 selector2.o11
connect o11-8.i2 reg11.o4

connect out8.i11 o11-8.o1


define o12-5 and 2 1
connect o12-5.i1 selector2.o12
connect o12-5.i2 reg12.o1

connect out5.i12 o12-5.o1

define o12-6 and 2 1
connect o12-6.i1 selector2.o12
connect o12-6.i2 reg12.o2

connect out6.i12 o12-6.o1

define o12-7 and 2 1
connect o12-7.i1 selector2.o12
connect o12-7.i2 reg12.o3

connect out7.i12 o12-7.o1

define o12-8 and 2 1
connect o12-8.i1 selector2.o12
connect o12-8.i2 reg12.o4

connect out8.i12 o12-8.o1


define o13-5 and 2 1
connect o13-5.i1 selector2.o13
connect o13-5.i2 reg13.o1

connect out5.i13 o13-5.o1

define o13-6 and 2 1
connect o13-6.i1 selector2.o13
connect o13-6.i2 reg13.o2

connect out6.i13 o13-6.o1

define o13-7 and 2 1
connect o13-7.i1 selector2.o13
connect o13-7.i2 reg13.o3

connect out7.i13 o13-7.o1

define o13-8 and 2 1
connect o13-8.i1 selector2.o13
connect o13-8.i2 reg13.o4

connect out8.i13 o13-8.o1


define o14-5 and 2 1
connect o14-5.i1 selector2.o14
connect o14-5.i2 reg14.o1

connect out5.i14 o14-5.o1

define o14-6 and 2 1
connect o14-6.i1 selector2.o14
connect o14-6.i2 reg14.o2

connect out6.i14 o14-6.o1

define o14-7 and 2 1
connect o14-7.i1 selector2.o14
connect o14-7.i2 reg14.o3

connect out7.i14 o14-7.o1

define o14-8 and 2 1
connect o14-8.i1 selector2.o14
connect o14-8.i2 reg14.o4

connect out8.i14 o14-8.o1


define o15-5 and 2 1
connect o15-5.i1 selector2.o15
connect o15-5.i2 reg15.o1

connect out5.i15 o15-5.o1

define o15-6 and 2 1
connect o15-6.i1 selector2.o15
connect o15-6.i2 reg15.o2

connect out6.i15 o15-6.o1

define o15-7 and 2 1
connect o15-7.i1 selector2.o15
connect o15-7.i2 reg15.o3

connect out7.i15 o15-7.o1

define o15-8 and 2 1
connect o15-8.i1 selector2.o15
connect o15-8.i2 reg15.o4

connect out8.i15 o15-8.o1


define o16-5 and 2 1
connect o16-5.i1 selector2.o16
connect o16-5.i2 reg16.o1

connect out5.i16 o16-5.o1

define o16-6 and 2 1
connect o16-6.i1 selector2.o16
connect o16-6.i2 reg16.o2

connect out6.i16 o16-6.o1

define o16-7 and 2 1
connect o16-7.i1 selector2.o16
connect o16-7.i2 reg16.o3

connect out7.i16 o16-7.o1

define o16-8 and 2 1
connect o16-8.i1 selector2.o16
connect o16-8.i2 reg16.o4

connect out8.i16 o16-8.o1