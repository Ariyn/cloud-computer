inputs 2

# inputs.1 = A
# inputs.2 = B
# output eq = not(A xor B)
# output gt means A is greater then B
#   gt = (A and not B)
# output lt means A is less then B
#   lt = (not A and B)

### Equal (eq)
define eq_xor xor 2 1 no-optimization
connect eq_xor.i1 $inputs.1
connect eq_xor.i2 $inputs.2

define eq_not not 1 1 no-optimization
connect eq_not.i1 eq_xor.o1

alias eq eq_not.o1

### Great Then (gt)
define gt_not not 1 1 no-optimization
connect gt_not.i1 $inputs.2

define gt_and and 2 1 no-optimization
connect gt_and.i1 $inputs.1
connect gt_and.i2 gt_not.o1

alias gt gt_and.o1

### Less Then (lt)
define lt_not not 1 1 no-optimization
connect lt_not.i1 $inputs.1

define lt_and and 2 1 no-optimization
connect lt_and.i1 lt_not.o1
connect lt_and.i2 $inputs.2

alias lt lt_and.o1