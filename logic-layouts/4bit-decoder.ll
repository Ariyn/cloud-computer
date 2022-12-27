inputs 4
# 4 inputs, 16 outputs
# selected output pin will have "true" value

define s1 not 1 1 no-optimization
connect s1.i1 $inputs.1

define s2 not 1 1 no-optimization
connect s2.i1 $inputs.2

define s3 not 1 1 no-optimization
connect s3.i1 $inputs.3

define s4 not 1 1 no-optimization
connect s4.i1 $inputs.4

# a1 = 0000
define a1 and 4 1 no-optimization
connect a1.i1 s1.o1
connect a1.i2 s2.o1
connect a1.i3 s3.o1
connect a1.i4 s4.o1

# a2 = 0001
define a2 and 4 1 no-optimization
connect a2.i1 $inputs.1
connect a2.i2 s2.o1
connect a2.i3 s3.o1
connect a2.i4 s4.o1

# a3 = 0010
define a3 and 4 1 no-optimization
connect a3.i1 s1.o1
connect a3.i2 $inputs.2
connect a3.i3 s3.o1
connect a3.i4 s4.o1

# a4 = 0011
define a4 and 4 1 no-optimization
connect a4.i1 $inputs.1
connect a4.i2 $inputs.2
connect a4.i3 s3.o1
connect a4.i4 s4.o1

# a5 = 0100
define a5 and 4 1 no-optimization
connect a5.i1 s1.o1
connect a5.i2 s2.o1
connect a5.i3 $inputs.3
connect a5.i4 s4.o1

# a6 = 0101
define a6 and 4 1 no-optimization
connect a6.i1 $inputs.1
connect a6.i2 s2.o1
connect a6.i3 $inputs.3
connect a6.i4 s4.o1

# a7 = 0110
define a7 and 4 1 no-optimization
connect a7.i1 s1.o1
connect a7.i2 $inputs.2
connect a7.i3 $inputs.3
connect a7.i4 s4.o1

# a8 = 0111
define a8 and 4 1 no-optimization
connect a8.i1 $inputs.2
connect a8.i2 $inputs.2
connect a8.i3 $inputs.3
connect a8.i4 s4.o1

# a9 = 1000
define a9 and 4 1 no-optimization
connect a9.i1 s1.o1
connect a9.i2 s2.o1
connect a9.i3 s3.o1
connect a9.i4 $inputs.4

# a10 = 1001
define a10 and 4 1 no-optimization
connect a10.i1 $inputs.1
connect a10.i2 s2.o1
connect a10.i3 s3.o1
connect a10.i4 $inputs.4

# a11 = 1010
define a11 and 4 1 no-optimization
connect a11.i1 s1.o1
connect a11.i2 $inputs.2
connect a11.i3 s3.o1
connect a11.i4 $inputs.4

# a12 = 1011
define a12 and 4 1 no-optimization
connect a12.i1 $inputs.1
connect a12.i2 $inputs.2
connect a12.i3 s3.o1
connect a12.i4 $inputs.4

# a13 = 1100
define a13 and 4 1 no-optimization
connect a13.i1 s1.o1
connect a13.i2 s2.o1
connect a13.i3 $inputs.3
connect a13.i4 $inputs.4

# a14 = 1101
define a14 and 4 1 no-optimization
connect a14.i1 $inputs.1
connect a14.i2 s2.o1
connect a14.i3 $inputs.3
connect a14.i4 $inputs.4

# a15 = 1110
define a15 and 4 1 no-optimization
connect a15.i1 s1.o1
connect a15.i2 $inputs.2
connect a15.i3 $inputs.3
connect a15.i4 $inputs.4

# a16 = 1111
define a16 and 4 1 no-optimization
connect a16.i1 $inputs.1
connect a16.i2 $inputs.2
connect a16.i3 $inputs.3
connect a16.i4 $inputs.4

alias o1 a1.o1
alias o2 a2.o1
alias o3 a3.o1
alias o4 a4.o1
alias o5 a5.o1
alias o6 a6.o1
alias o7 a7.o1
alias o8 a8.o1
alias o9 a9.o1
alias o11 a11.o1
alias o12 a12.o1
alias o13 a13.o1
alias o13 a13.o1
alias o14 a14.o1
alias o15 a15.o1
alias o16 a16.o1
