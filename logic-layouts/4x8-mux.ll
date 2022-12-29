inputs 5

# TODO: Rename to mux-bus?
# pass input tp selected output.
# inputs.1 ~ inputs.4 = inputs
#   1 = LSB, 4 = MSB
# inputs.5 = selector
#  if selector is 0: inputs a selected
#  if selector is 1: inputs b selected

# output.1 ~ output.4 = output a
#   1 is LSB, 4 is MSB
# output.5 ~ output.8 = output b
#   5 is LSB, 8 is MSB

define not not 1 1
connect not.i1 $inputs.5

define a1-1 and 2 1
connect a1-1.i1 $inputs.1
connect a1-1.i2 not.o1
alias o1 a1-1.o1

define a1-2 and 2 1
connect a1-2.i1 $inputs.2
connect a1-2.i2 not.o1
alias o2 a1-2.o1

define a1-3 and 2 1
connect a1-3.i1 $inputs.3
connect a1-3.i2 not.o1
alias o3 a1-3.o1

define a1-4 and 2 1
connect a1-4.i1 $inputs.4
connect a1-4.i2 not.o1
alias o4 a1-4.o1

define a2-1 and 2 1
connect a2-1.i1 $inputs.1
connect a2-1.i2 $inputs.5
alias o5 a2-1.o1

define a2-2 and 2 1
connect a2-2.i1 $inputs.2
connect a2-2.i2 $inputs.5
alias o6 a2-2.o1

define a2-3 and 2 1
connect a2-3.i1 $inputs.3
connect a2-3.i2 $inputs.5
alias o7 a2-3.o1

define a2-4 and 2 1
connect a2-4.i1 $inputs.4
connect a2-4.i2 $inputs.5
alias o8 a2-4.o1
