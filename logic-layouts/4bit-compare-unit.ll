inputs 8

# inputs.1~4 = A
#  inputs.1 is LSB, inputs.4 is MSB
# inputs.5~8 = B
#  inputs.5 is LSB, inputs.8 is MSB

define cu1 1bit-compare-unit 2 3
connect cu1.i1 $inputs.1
connect cu1.i2 $inputs.5

define cu2 1bit-compare-unit 2 3
connect cu2.i1 $inputs.2
connect cu2.i2 $inputs.6

define cu3 1bit-compare-unit 2 3
connect cu3.i1 $inputs.3
connect cu3.i2 $inputs.7

define cu4 1bit-compare-unit 2 3
connect cu4.i1 $inputs.4
connect cu4.i2 $inputs.8

### GT
define gt_and1 and 4 1
connect gt_and1.i1 cu2.eq
connect gt_and1.i2 cu3.eq
connect gt_and1.i3 cu4.eq
connect gt_and1.i4 cu1.gt

define gt_and2 and 3 1
connect gt_and2.i1 cu3.eq
connect gt_and2.i2 cu4.eq
connect gt_and2.i3 cu2.gt

define gt_and3 and 2 1
connect gt_and3.i1 cu4.eq
connect gt_and3.i2 cu3.gt

define gt_or or 4 1
connect gt_or.i1 cu4.gt
connect gt_or.i2 gt_and1.o1
connect gt_or.i3 gt_and2.o1
connect gt_or.i4 gt_and3.o1

alias gt gt_or.o1

### LT
define lt_and1 and 4 1
connect lt_and1.i1 cu2.eq
connect lt_and1.i2 cu3.eq
connect lt_and1.i3 cu4.eq
connect lt_and1.i4 cu1.lt
### TODO: 위 부분에서 cu2.eq, cu3.eq, cu4.eq는 아래의 eq와 결합해서 어떻게 최적화 할 수 있을 것 같다.

define lt_and2 and 3 1
connect lt_and2.i1 cu3.eq
connect lt_and2.i2 cu4.eq
connect lt_and2.i3 cu2.lt

define lt_and3 and 2 1
connect lt_and3.i1 cu4.eq
connect lt_and3.i2 cu3.lt

define lt_or or 4 1
connect lt_or.i1 cu4.lt
connect lt_or.i2 lt_and1.o1
connect lt_or.i3 lt_and2.o1
connect lt_or.i4 lt_and3.o1

alias lt lt_or.o1

### EQ
define eq_and and 4 1
connect eq_and.i1 cu1.eq
connect eq_and.i2 cu2.eq
connect eq_and.i3 cu3.eq
connect eq_and.i4 cu4.eq

alias eq eq_and.o1