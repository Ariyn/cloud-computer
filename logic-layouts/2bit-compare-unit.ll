inputs 4

# inputs.1~2 = A
#  inputs.1 is LSB, inputs.2 is MSB
# inputs.3~4 = B
#  inputs.3 is LSB, inputs.4 is MSB

define cu1 1bit-compare-unit 2 3
connect cu1.i1 $inputs.1
connect cu1.i2 $inputs.3

define cu2 1bit-compare-unit 2 3
connect cu2.i1 $inputs.2
connect cu2.i2 $inputs.4

### GT
define gt_and and 2 1
connect gt_and.i1 cu2.eq
connect gt_and.i2 cu1.gt

define gt_or or 2 1
connect gt_or.i1 cu2.gt
connect gt_or.i2 gt_and.o1

alias gt gt_or.o1

### LT
define lt_and and 2 1
connect lt_and.i1 cu2.eq
connect lt_and.i2 cu1.lt

define lt_or or 2 1
connect lt_or.i1 cu2.lt
connect lt_or.i2 lt_and.o1

alias lt lt_or.o1

### EQ
define eq_and and 2 1
connect eq_and.i1 cu1.eq
connect eq_and.i2 cu2.eq

alias eq eq_and.o1
