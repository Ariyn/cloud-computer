inputs 4

# inputs.1 = a
# inputs.2 = b
# inputs.3 = li (loan in)
# inputs.4 = output selector
#   when output selector is 1, output is a
#   when output selector is 0, output is a-b

define sub subtractor 3 1
connect sub.i1 $inputs.1
connect sub.i2 $inputs.2
connect sub.i3 $inputs.3

alias lo sub.lo

define and1 and 2 1
connect and1.i1 $inputs.1
connect and1.i2 $inputs.4

define not1 not 1 1
connect not1.i1 $inputs.4

define and2 and 2 1
connect and2.i1 sub.d
connect and2.i2 not1.o1

define or or 2 1
connect or.i1 and1.o1
connect or.i2 and2.o1

alias d or.o1