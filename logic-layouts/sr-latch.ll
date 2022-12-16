inputs 2

define nor1 nor 2 1
define nor2 nor 2 1

connect nor1.i1 $inputs.1
connect nor1.i2 nor2.output # 아웃풋 이름이 항상 달라지는 문제가 있음. 수정할 것

connect nor2.i1 $inputs.2
connect nor2.i2 nor1.output

alias q nor1.output
alias q-bar nor2.output