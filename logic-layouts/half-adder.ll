# inputs 2

define x1 xor 2 1
define a1 and 2 1

connect x1.i1 ${i1} # 통일성을 위해서 i1이 아니라 inputs.1으로 되어야 함
connect x1.i2 ${i2}

connect a1.i1 ${i1}
connect a1.i2 ${i2}

alias sum x1.o1 # name.sum 으로 이름이 바뀌어야 함
alias carry a1.o1

# export sum
# export carry