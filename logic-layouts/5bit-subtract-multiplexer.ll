inputs 11
# inputs 1~5 = a
# inputs 6~10 = b
# inputs 11 = li (loan in)
# output selector is connected to last bit's lo

define sub1 subtract-multiplexer 4 2
define sub2 subtract-multiplexer 4 2
define sub3 subtract-multiplexer 4 2
define sub4 subtract-multiplexer 4 2
define sub5 subtract-multiplexer 4 2

connect sub1.i4 sub5.lo
connect sub2.i4 sub5.lo
connect sub3.i4 sub5.lo
connect sub4.i4 sub5.lo
connect sub5.i4 sub5.lo

connect sub1.i1 $inputs.1
connect sub2.i1 $inputs.2
connect sub3.i1 $inputs.3
connect sub4.i1 $inputs.4
connect sub5.i1 $inputs.5

connect sub1.i2 $inputs.6
connect sub2.i2 $inputs.7
connect sub3.i2 $inputs.8
connect sub4.i2 $inputs.9
connect sub5.i2 $inputs.10

connect sub1.i3 $inputs.11
connect sub2.i3 sub1.lo
connect sub3.i3 sub2.lo
connect sub4.i3 sub3.lo
connect sub5.i3 sub4.lo

# o1 = LSB, o5 = MSB
alias o1 sub1.d
alias o2 sub2.d
alias o3 sub3.d
alias o4 sub4.d
alias o5 sub5.d

alias lo sub5.lo