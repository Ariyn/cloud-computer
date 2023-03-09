inputs 64

# inputs 1~32 = A
#    inputs.1 = LSB, inputs.32 = MSB

# inputs 33 ~ 64 = B
#    inputs.33 = LSB, inputs.64 = MSB

define fa1 full-adder 3 3
define fa2 full-adder 3 3
define fa3 full-adder 3 3
define fa4 full-adder 3 3
define fa5 full-adder 3 3
define fa6 full-adder 3 3
define fa7 full-adder 3 3
define fa8 full-adder 3 3
define fa9 full-adder 3 3
define fa10 full-adder 3 3
define fa11 full-adder 3 3
define fa12 full-adder 3 3
define fa13 full-adder 3 3
define fa14 full-adder 3 3
define fa15 full-adder 3 3
define fa16 full-adder 3 3
define fa17 full-adder 3 3
define fa18 full-adder 3 3
define fa19 full-adder 3 3
define fa20 full-adder 3 3
define fa21 full-adder 3 3
define fa22 full-adder 3 3
define fa23 full-adder 3 3
define fa24 full-adder 3 3
define fa25 full-adder 3 3
define fa26 full-adder 3 3
define fa27 full-adder 3 3
define fa28 full-adder 3 3
define fa29 full-adder 3 3
define fa30 full-adder 3 3
define fa31 full-adder 3 3
define fa32 full-adder 3 3

connect fa1.i1 $inputs.1
connect fa1.i2 $inputs.33
connect fa1.i3 0

connect fa2.i1 $inputs.2
connect fa2.i2 $inputs.34
connect fa2.i3 fa1.carry

connect fa3.i1 $inputs.3
connect fa3.i2 $inputs.35
connect fa3.i3 fa2.carry

connect fa4.i1 $inputs.4
connect fa4.i2 $inputs.36
connect fa4.i3 fa3.carry

connect fa5.i1 $inputs.5
connect fa5.i2 $inputs.37
connect fa5.i3 fa4.carry

connect fa6.i1 $inputs.6
connect fa6.i2 $inputs.38
connect fa6.i3 fa5.carry

connect fa7.i1 $inputs.7
connect fa7.i2 $inputs.39
connect fa7.i3 fa6.carry

connect fa8.i1 $inputs.8
connect fa8.i2 $inputs.40
connect fa8.i3 fa7.carry

connect fa9.i1 $inputs.9
connect fa9.i2 $inputs.41
connect fa9.i3 fa8.carry

connect fa10.i1 $inputs.10
connect fa10.i2 $inputs.42
connect fa10.i3 fa9.carry

connect fa11.i1 $inputs.11
connect fa11.i2 $inputs.43
connect fa11.i3 fa10.carry

connect fa12.i1 $inputs.12
connect fa12.i2 $inputs.44
connect fa12.i3 fa11.carry

connect fa13.i1 $inputs.13
connect fa13.i2 $inputs.45
connect fa13.i3 fa12.carry

connect fa14.i1 $inputs.14
connect fa14.i2 $inputs.46
connect fa14.i3 fa13.carry

connect fa15.i1 $inputs.15
connect fa15.i2 $inputs.47
connect fa15.i3 fa14.carry

connect fa16.i1 $inputs.16
connect fa16.i2 $inputs.48
connect fa16.i3 fa15.carry

connect fa17.i1 $inputs.17
connect fa17.i2 $inputs.49
connect fa17.i3 fa16.carry

connect fa18.i1 $inputs.18
connect fa18.i2 $inputs.50
connect fa18.i3 fa17.carry

connect fa19.i1 $inputs.19
connect fa19.i2 $inputs.51
connect fa19.i3 fa18.carry

connect fa20.i1 $inputs.20
connect fa20.i2 $inputs.52
connect fa20.i3 fa19.carry

connect fa21.i1 $inputs.21
connect fa21.i2 $inputs.53
connect fa21.i3 fa20.carry

connect fa22.i1 $inputs.22
connect fa22.i2 $inputs.54
connect fa22.i3 fa21.carry

connect fa23.i1 $inputs.23
connect fa23.i2 $inputs.55
connect fa23.i3 fa22.carry

connect fa24.i1 $inputs.24
connect fa24.i2 $inputs.56
connect fa24.i3 fa23.carry

connect fa25.i1 $inputs.25
connect fa25.i2 $inputs.57
connect fa25.i3 fa24.carry

connect fa26.i1 $inputs.26
connect fa26.i2 $inputs.58
connect fa26.i3 fa25.carry

connect fa27.i1 $inputs.27
connect fa27.i2 $inputs.59
connect fa27.i3 fa26.carry

connect fa28.i1 $inputs.28
connect fa28.i2 $inputs.60
connect fa28.i3 fa27.carry

connect fa29.i1 $inputs.29
connect fa29.i2 $inputs.61
connect fa29.i3 fa28.carry

connect fa30.i1 $inputs.30
connect fa30.i2 $inputs.62
connect fa30.i3 fa29.carry

connect fa31.i1 $inputs.31
connect fa31.i2 $inputs.63
connect fa31.i3 fa30.carry

connect fa32.i1 $inputs.32
connect fa32.i2 $inputs.64
connect fa32.i3 fa31.carry

alias o1 ffa1.sum
alias o2 ffa2.sum
alias o3 ffa3.sum
alias o4 ffa4.sum
alias o5 ffa5.sum
alias o6 ffa6.sum
alias o7 ffa7.sum
alias o8 ffa8.sum
alias o9 ffa9.sum
alias o10 ffa10.sum
alias o11 ffa11.sum
alias o12 ffa12.sum
alias o13 ffa13.sum
alias o14 ffa14.sum
alias o15 ffa15.sum
alias o16 ffa16.sum
alias o17 ffa17.sum
alias o18 ffa18.sum
alias o19 ffa19.sum
alias o20 ffa20.sum
alias o21 ffa21.sum
alias o22 ffa22.sum
alias o23 ffa23.sum
alias o24 ffa24.sum
alias o25 ffa25.sum
alias o26 ffa26.sum
alias o27 ffa27.sum
alias o28 ffa28.sum
alias o29 ffa29.sum
alias o30 ffa30.sum
alias o31 ffa31.sum
alias o32 ffa32.sum