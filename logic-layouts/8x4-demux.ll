inputs 9

# TODO: Rename to demux-bus?
# select inputs and pass to output.
# inputs.1 ~ inputs.4 = inputs a. 1 = LSB, 4 = MSB
# inputs.5 ~ inputs.8 = inputs b. 4 = LSB, 8 = MSB
# inputs.9 = selector
#  if selector is 0: inputs a selected
#  if selector is 1: inputs b selected

# output.1 ~ output.4 = output

define not not 1 1
connect not.i1 $inputs.9

define a1-1 and 2 1
connect a1-1.i1 $inputs.1
connect a1-1.i2 not.o1

define a1-2 and 2 1
connect a1-2.i1 $inputs.2
connect a1-2.i2 not.o1

define a1-3 and 2 1
connect a1-3.i1 $inputs.3
connect a1-3.i2 not.o1

define a1-4 and 2 1
connect a1-4.i1 $inputs.4
connect a1-4.i2 not.o1

define a2-1 and 2 1
connect a2-1.i1 $inputs.5
connect a2-1.i2 $inputs.9

define a2-2 and 2 1
connect a2-2.i1 $inputs.6
connect a2-2.i2 $inputs.9

define a2-3 and 2 1
connect a2-3.i1 $inputs.7
connect a2-3.i2 $inputs.9

define a2-4 and 2 1
connect a2-4.i1 $inputs.8
connect a2-4.i2 $inputs.9

define out1 or 2 1
connect out1.i1 a1-1.o1
connect out1.i2 a2-1.o1
alias o1 out1.o1

define out2 or 2 1
connect out2.i1 a1-2.o1
connect out2.i2 a2-2.o1
alias o2 out2.o1

define out3 or 2 1
connect out3.i1 a1-3.o1
connect out3.i2 a2-3.o1
alias o3 out3.o1

define out4 or 2 1
connect out4.i1 a1-4.o1
connect out4.i2 a2-4.o1
alias o4 out4.o1