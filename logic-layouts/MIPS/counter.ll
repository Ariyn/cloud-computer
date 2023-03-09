inputs 35

# inputs 1 ~ 32 is not implemented yet.
# inputs 1~32 = inputs
#   inputs 1 = LSB, inputs 32 = MSB
# inputs 33 = reset
# inputs 34 = clock
# inputs 35 = enable

define r MIPS/register 35 32
connect r.i33 $inputs.33
connect r.i34 $inputs.34
connect r.i35 $inputs.35

define sa MIPS/32bit-adder 64 32
connect sa.i1 r.o1
connect sa.i2 r.o2
connect sa.i3 r.o3
connect sa.i4 r.o4
connect sa.i5 r.o5
connect sa.i6 r.o6
connect sa.i7 r.o7
connect sa.i8 r.o8
connect sa.i9 r.o9
connect sa.i10 r.o10
connect sa.i11 r.o11
connect sa.i12 r.o12
connect sa.i13 r.o13
connect sa.i14 r.o14
connect sa.i15 r.o15
connect sa.i16 r.o16
connect sa.i17 r.o17
connect sa.i18 r.o18
connect sa.i19 r.o19
connect sa.i20 r.o20
connect sa.i21 r.o21
connect sa.i22 r.o22
connect sa.i23 r.o23
connect sa.i24 r.o24
connect sa.i25 r.o25
connect sa.i26 r.o26
connect sa.i27 r.o27
connect sa.i28 r.o28
connect sa.i29 r.o29
connect sa.i30 r.o30
connect sa.i31 r.o31
connect sa.i32 r.o32

connect sa.i33 0
connect sa.i34 0
# i35 = 4
connect sa.i35 1
connect sa.i36 0
connect sa.i37 0
connect sa.i38 0
connect sa.i39 0
connect sa.i40 0
connect sa.i41 0
connect sa.i42 0
connect sa.i43 0
connect sa.i44 0
connect sa.i45 0
connect sa.i46 0
connect sa.i47 0
connect sa.i48 0
connect sa.i49 0
connect sa.i50 0
connect sa.i51 0
connect sa.i52 0
connect sa.i53 0
connect sa.i54 0
connect sa.i55 0
connect sa.i56 0
connect sa.i57 0
connect sa.i58 0
connect sa.i59 0
connect sa.i60 0
connect sa.i61 0
connect sa.i62 0
connect sa.i63 0
connect sa.i64 0

connect r.i1 sa.o1
connect r.i2 sa.o2
connect r.i3 sa.o3
connect r.i4 sa.o4
connect r.i5 sa.o5
connect r.i6 sa.o6
connect r.i7 sa.o7
connect r.i8 sa.o8
connect r.i9 sa.o9
connect r.i10 sa.o10
connect r.i11 sa.o11
connect r.i12 sa.o12
connect r.i13 sa.o13
connect r.i14 sa.o14
connect r.i15 sa.o15
connect r.i16 sa.o16
connect r.i17 sa.o17
connect r.i18 sa.o18
connect r.i19 sa.o19
connect r.i20 sa.o20
connect r.i21 sa.o21
connect r.i22 sa.o22
connect r.i23 sa.o23
connect r.i24 sa.o24
connect r.i25 sa.o25
connect r.i26 sa.o26
connect r.i27 sa.o27
connect r.i28 sa.o28
connect r.i29 sa.o29
connect r.i30 sa.o30
connect r.i31 sa.o31
connect r.i32 sa.o32

alias o1 r.o1
alias o2 r.o2
alias o3 r.o3
alias o4 r.o4
alias o5 r.o5
alias o6 r.o6
alias o7 r.o7
alias o8 r.o8
alias o9 r.o9
alias o10 r.o10
alias o11 r.o11
alias o12 r.o12
alias o13 r.o13
alias o14 r.o14
alias o15 r.o15
alias o16 r.o16
alias o17 r.o17
alias o18 r.o18
alias o19 r.o19
alias o20 r.o20
alias o21 r.o21
alias o22 r.o22
alias o23 r.o23
alias o24 r.o24
alias o25 r.o25
alias o26 r.o26
alias o27 r.o27
alias o28 r.o28
alias o29 r.o29
alias o30 r.o30
alias o31 r.o31
alias o32 r.o32