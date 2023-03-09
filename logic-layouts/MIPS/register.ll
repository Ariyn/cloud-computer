inputs 35
# in mips 1 word is 32 bits
# inputs 1~32 = inputs
#   inputs 1 = LSB, inputs 32 = MSB
# inputs 33 = reset
# inputs 34 = clock
# inputs 35 = enable

define trigger complex/rising-edge-trigger 1 1
connect trigger.i1 $inputs.34

# ec = enabled clock
define ec and 2 1
connect ec.i1 $inputs.35
connect ec.i2 trigger.o1

define ff1 flipflop 3 2
connect ff1.i1 $inputs.1
connect ff1.i2 $inputs.33
connect ff1.i3 ec.o1

define ff2 flipflop 3 2
connect ff2.i1 $inputs.2
connect ff2.i2 $inputs.33
connect ff2.i3 ec.o1

define ff3 flipflop 3 2
connect ff3.i1 $inputs.3
connect ff3.i2 $inputs.33
connect ff3.i3 ec.o1

define ff4 flipflop 3 2
connect ff4.i1 $inputs.4
connect ff4.i2 $inputs.33
connect ff4.i3 ec.o1

define ff5 flipflop 3 2
connect ff5.i1 $inputs.5
connect ff5.i2 $inputs.33
connect ff5.i3 ec.o1

define ff6 flipflop 3 2
connect ff6.i1 $inputs.6
connect ff6.i2 $inputs.33
connect ff6.i3 ec.o1

define ff7 flipflop 3 2
connect ff7.i1 $inputs.7
connect ff7.i2 $inputs.33
connect ff7.i3 ec.o1

define ff8 flipflop 3 2
connect ff8.i1 $inputs.8
connect ff8.i2 $inputs.33
connect ff8.i3 ec.o1

define ff9 flipflop 3 2
connect ff9.i1 $inputs.9
connect ff9.i2 $inputs.33
connect ff9.i3 ec.o1

define ff10 flipflop 3 2
connect ff10.i1 $inputs.10
connect ff10.i2 $inputs.33
connect ff10.i3 ec.o1

define ff11 flipflop 3 2
connect ff11.i1 $inputs.11
connect ff11.i2 $inputs.33
connect ff11.i3 ec.o1

define ff12 flipflop 3 2
connect ff12.i1 $inputs.12
connect ff12.i2 $inputs.33
connect ff12.i3 ec.o1

define ff13 flipflop 3 2
connect ff13.i1 $inputs.13
connect ff13.i2 $inputs.33
connect ff13.i3 ec.o1

define ff14 flipflop 3 2
connect ff14.i1 $inputs.14
connect ff14.i2 $inputs.33
connect ff14.i3 ec.o1

define ff15 flipflop 3 2
connect ff15.i1 $inputs.15
connect ff15.i2 $inputs.33
connect ff15.i3 ec.o1

define ff16 flipflop 3 2
connect ff16.i1 $inputs.16
connect ff16.i2 $inputs.33
connect ff16.i3 ec.o1

define ff17 flipflop 3 2
connect ff17.i1 $inputs.17
connect ff17.i2 $inputs.33
connect ff17.i3 ec.o1

define ff18 flipflop 3 2
connect ff18.i1 $inputs.18
connect ff18.i2 $inputs.33
connect ff18.i3 ec.o1

define ff19 flipflop 3 2
connect ff19.i1 $inputs.19
connect ff19.i2 $inputs.33
connect ff19.i3 ec.o1

define ff20 flipflop 3 2
connect ff20.i1 $inputs.20
connect ff20.i2 $inputs.33
connect ff20.i3 ec.o1

define ff21 flipflop 3 2
connect ff21.i1 $inputs.21
connect ff21.i2 $inputs.33
connect ff21.i3 ec.o1

define ff22 flipflop 3 2
connect ff22.i1 $inputs.22
connect ff22.i2 $inputs.33
connect ff22.i3 ec.o1

define ff23 flipflop 3 2
connect ff23.i1 $inputs.23
connect ff23.i2 $inputs.33
connect ff23.i3 ec.o1

define ff24 flipflop 3 2
connect ff24.i1 $inputs.24
connect ff24.i2 $inputs.33
connect ff24.i3 ec.o1

define ff25 flipflop 3 2
connect ff25.i1 $inputs.25
connect ff25.i2 $inputs.33
connect ff25.i3 ec.o1

define ff26 flipflop 3 2
connect ff26.i1 $inputs.26
connect ff26.i2 $inputs.33
connect ff26.i3 ec.o1

define ff27 flipflop 3 2
connect ff27.i1 $inputs.27
connect ff27.i2 $inputs.33
connect ff27.i3 ec.o1

define ff28 flipflop 3 2
connect ff28.i1 $inputs.28
connect ff28.i2 $inputs.33
connect ff28.i3 ec.o1

define ff29 flipflop 3 2
connect ff29.i1 $inputs.29
connect ff29.i2 $inputs.33
connect ff29.i3 ec.o1

define ff30 flipflop 3 2
connect ff30.i1 $inputs.30
connect ff30.i2 $inputs.33
connect ff30.i3 ec.o1

define ff31 flipflop 3 2
connect ff31.i1 $inputs.31
connect ff31.i2 $inputs.33
connect ff31.i3 ec.o1

define ff32 flipflop 3 2
connect ff32.i1 $inputs.32
connect ff32.i2 $inputs.33
connect ff32.i3 ec.o1

alias o1 ff1.o1
alias o2 ff2.o1
alias o3 ff3.o1
alias o4 ff4.o1
alias o5 ff5.o1
alias o6 ff6.o1
alias o7 ff7.o1
alias o8 ff8.o1
alias o9 ff9.o1
alias o10 ff10.o1
alias o11 ff11.o1
alias o12 ff12.o1
alias o13 ff13.o1
alias o14 ff14.o1
alias o15 ff15.o1
alias o16 ff16.o1
alias o17 ff17.o1
alias o18 ff18.o1
alias o19 ff19.o1
alias o20 ff20.o1
alias o21 ff21.o1
alias o22 ff22.o1
alias o23 ff23.o1
alias o24 ff24.o1
alias o25 ff25.o1
alias o26 ff26.o1
alias o27 ff27.o1
alias o28 ff28.o1
alias o29 ff29.o1
alias o30 ff30.o1
alias o31 ff31.o1
alias o32 ff32.o1