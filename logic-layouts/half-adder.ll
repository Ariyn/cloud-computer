# inputs 2

define x1 xor 2 1
define a1 and 2 1

connect x1.i1 inputs.1
connect x1.i2 inputs.2

connect a1.i1 inputs.1
connect a1.i2 inputs.2

# alias sum x1.o1
# alias carry a1.o1

# export sum
# export carry