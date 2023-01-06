inputs 3

# inputs.1 = LSB, inputs.3 = MSB
# parse inputs as binary number, select number of output pin

# 3 selectors, 8 outputs

define a1 and 3 1 no-optimization
define a2 and 3 1 no-optimization
define a3 and 3 1 no-optimization
define a4 and 3 1 no-optimization
define a5 and 3 1 no-optimization
define a6 and 3 1 no-optimization
define a7 and 3 1 no-optimization
define a8 and 3 1 no-optimization

define sn1 not 1 1 no-optimization
connect sn1.i1 $inputs.1

define sn2 not 1 1 no-optimization
connect sn2.i1 $inputs.2

define sn3 not 1 1 no-optimization
connect sn3.i1 $inputs.3

connect a1.i1 sn1.o1
connect a1.i2 sn2.o1
connect a1.i3 sn3.o1

connect a2.i1 $inputs.1
connect a2.i2 sn2.o1
connect a2.i3 sn3.o1

connect a3.i1 sn1.o1
connect a3.i2 $inputs.2
connect a3.i3 sn3.o1

connect a4.i1 $inputs.1
connect a4.i2 $inputs.2
connect a4.i3 sn3.o1

connect a5.i1 sn1.o1
connect a5.i2 sn2.o1
connect a5.i3 $inputs.3

connect a6.i1 $inputs.1
connect a6.i2 sn2.o1
connect a6.i3 $inputs.3

connect a7.i1 sn1.o1
connect a7.i2 $inputs.2
connect a7.i3 $inputs.3

connect a8.i1 $inputs.1
connect a8.i2 $inputs.2
connect a8.i3 $inputs.3

alias o1 a1.o1
alias o2 a2.o1
alias o3 a3.o1
alias o4 a4.o1
alias o5 a5.o1
alias o6 a6.o1
alias o7 a7.o1
alias o8 a8.o1
