inputs 7

# this register save data when clock rising edge triggered and enabled.
# rising edge duration does not decided. (depends on 3 not gates duration)

# inputs.1 ~ inputs.4 = bits
# inputs.5 = reset
# inputs.6 = clk
# inputs.7 = enabled
# if enabled is false, register does not update data

define trigger complex/falling-edge-trigger 1 1
connect trigger.i1 $inputs.6

define enabled_clock and 2 1
connect enabled_clock.i1 trigger.o1
connect enabled_clock.i2 $inputs.7

define ff1 flipflop 3 2
connect ff1.i1 $inputs.1
connect ff1.i2 $inputs.5
connect ff1.i3 enabled_clock.o1

define ff2 flipflop 3 2
connect ff2.i1 $inputs.2
connect ff2.i2 $inputs.5
connect ff2.i3 enabled_clock.o1

define ff3 flipflop 3 2
connect ff3.i1 $inputs.3
connect ff3.i2 $inputs.5
connect ff3.i3 enabled_clock.o1

define ff4 flipflop 3 2
connect ff4.i1 $inputs.4
connect ff4.i2 $inputs.5
connect ff4.i3 enabled_clock.o1

alias o1 ff1.o1
alias o2 ff2.o1
alias o3 ff3.o1
alias o4 ff4.o1