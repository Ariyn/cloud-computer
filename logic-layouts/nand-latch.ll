inputs 2

define nand1 nand 2 1 no-optimization
define nand2 nand 2 1 no-optimization

connect nand1.i1 $inputs.1
connect nand1.i2 nand2.output # 아웃풋 이름이 항상 달라지는 문제가 있음. 수정할 것

connect nand2.i1 $inputs.2
connect nand2.i2 nand1.output

alias q nand1.output
alias q-bar nand2.output