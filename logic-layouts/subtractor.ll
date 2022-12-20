inputs 8

define a1 full-adder 3 2
define a2 full-adder 3 2
define a3 full-adder 3 2
define a4 full-adder 3 2

# 여기서 빅 엔디안? 리틀 엔디안? 이거 확실하게 정해줄 것
# 이 회로는 A-B인가? B-A인가? 확실하게 정할 것
connect a1.i1 $inputs.1
connect a2.i1 $inputs.2
connect a3.i1 $inputs.3
connect a4.i1 $inputs.4

connect a1.i3 1
connect a2.i3 a1.carry
connect a3.i3 a2.carry
connect a4.i3 a3.carry

define x1 xor 2 1
define x2 xor 2 1
define x3 xor 2 1
define x4 xor 2 1

connect x1.i1 $inputs.5
connect x1.i2 1

connect x2.i1 $inputs.6
connect x2.i2 1

connect x3.i1 $inputs.7
connect x3.i2 1

connect x4.i1 $inputs.8
connect x4.i2 1

connect a1.i2 x1.o1
connect a2.i2 x2.o1
connect a3.i2 x3.o1
connect a4.i2 x4.o1

alias o1 a1.sum
alias o2 a2.sum
alias o3 a3.sum
alias o4 a4.sum
alias o5 a4.carry
alias carry a4.carry
