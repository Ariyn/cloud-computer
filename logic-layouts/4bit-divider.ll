inputs 8

# inputs.1 ~ inputs.4 = a
#    inputs.1 = LSB
#    inputs.4 = MSB
# inputs.5 ~ inputs.8 = b
# a / b = o1 ~ o2
# remain = r1 ~ r2

# o1 = LSB, o4 = MSB
# r1 = LSB, r4 = MSB

# o4 MSB
define 5b-sm4 5bit-subtract-multiplexer 11 5
connect 5b-sm4.i1 $inputs.4
connect 5b-sm4.i2 0
connect 5b-sm4.i3 0
connect 5b-sm4.i4 0
connect 5b-sm4.i5 0
connect 5b-sm4.i6 $inputs.5
connect 5b-sm4.i7 $inputs.6
connect 5b-sm4.i8 $inputs.7
connect 5b-sm4.i9 $inputs.8
connect 5b-sm4.i10 0
connect 5b-sm4.i11 0

define not4 not 1 1
connect not4.i1 5b-sm4.lo
alias o4 not4.o1

# o3
define 5b-sm3 5bit-subtract-multiplexer 11 5
connect 5b-sm3.i1 $inputs.3
connect 5b-sm3.i2 5b-sm4.o1
connect 5b-sm3.i3 5b-sm4.o2
connect 5b-sm3.i4 5b-sm4.o3
connect 5b-sm3.i5 5b-sm4.o4
connect 5b-sm3.i6 $inputs.5
connect 5b-sm3.i7 $inputs.6
connect 5b-sm3.i8 $inputs.7
connect 5b-sm3.i9 $inputs.8
connect 5b-sm3.i10 0
connect 5b-sm3.i11 0

define not3 not 1 1
connect not3.i1 5b-sm3.lo
alias o3 not3.o1

# o2
define 5b-sm2 5bit-subtract-multiplexer 11 5
connect 5b-sm2.i1 $inputs.2
connect 5b-sm2.i2 5b-sm3.o1
connect 5b-sm2.i3 5b-sm3.o2
connect 5b-sm2.i4 5b-sm3.o3
connect 5b-sm2.i5 5b-sm3.o4
connect 5b-sm2.i6 $inputs.5
connect 5b-sm2.i7 $inputs.6
connect 5b-sm2.i8 $inputs.7
connect 5b-sm2.i9 $inputs.8
connect 5b-sm2.i10 0
connect 5b-sm2.i11 0

define not2 not 1 1
connect not2.i1 5b-sm2.lo
alias o2 not2.o1

# o1 LSB
define 5b-sm1 5bit-subtract-multiplexer 11 5
connect 5b-sm1.i1 $inputs.1
connect 5b-sm1.i2 5b-sm2.o1
connect 5b-sm1.i3 5b-sm2.o2
connect 5b-sm1.i4 5b-sm2.o3
connect 5b-sm1.i5 5b-sm2.o4
connect 5b-sm1.i6 $inputs.5
connect 5b-sm1.i7 $inputs.6
connect 5b-sm1.i8 $inputs.7
connect 5b-sm1.i9 $inputs.8
connect 5b-sm1.i10 0
connect 5b-sm1.i11 0

define not1 not 1 1
connect not1.i1 5b-sm1.lo
alias o1 not1.o1

# remains
alias r1 5b-sm1.o1
alias r2 5b-sm1.o2
alias r3 5b-sm1.o3
alias r4 5b-sm1.o4