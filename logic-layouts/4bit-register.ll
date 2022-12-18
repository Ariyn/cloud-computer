inputs 6

# inputs.1 ~ inputs.4 = bits
# inputs.5 = reset
# inputs.6 = clk

define ff1 flipflop 3 2
connect ff1.i1 $inputs.1
connect ff1.i2 $inputs.5
connect ff1.i3 $inputs.6

define ff2 flipflop 3 2
connect ff2.i1 $inputs.2
connect ff2.i2 $inputs.5
connect ff2.i3 $inputs.6

define ff3 flipflop 3 2
connect ff3.i1 $inputs.3
connect ff3.i2 $inputs.5
connect ff3.i3 $inputs.6

define ff4 flipflop 3 2
connect ff4.i1 $inputs.4
connect ff4.i2 $inputs.5
connect ff4.i3 $inputs.6

alias o1 ff1.o1
alias o2 ff2.o1
alias o3 ff3.o1
alias o4 ff4.o1