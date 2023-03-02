inputs 1

# circuit looks like
# input.1 ------------------ nor
#         |                   |
#         not1 - not2 - not3--

# inputs 1 = clock, pulse or whatever 1bit signal
# when inputs.1 is 0
#   not 3 is 1
#   nor has (0, 1) inputs, so it does not work

# when inputs.1 changed into 0 from 1
#   `nor` has (0,0) after receive inputs.1 input
#   after 3 not gate receive and transmit signal, `nor` has (1, 0) input.
#   only 3 not gate duration, nor has 1 output

# when inputs.1 changed into 1 from 0
#   `nor` has (1, 1) after receive inputs.1 input
#   so nothing changed

define not1 not 1 1
connect not1.i1 $inputs.1

define not2 not 1 1
connect not2.i1 not1.o1

define not3 not 1 1
connect not3.i1 not2.o1

define nor and 2 1
connect nor.i1 $inputs.1
connect nor.i2 not3.o1

alias o1 nor.o1
