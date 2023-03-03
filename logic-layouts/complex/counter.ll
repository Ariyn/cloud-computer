inputs 1

define trigger experiments/rising-edge-trigger 1 1
connect trigger.i1 $inputs.1

define r complex/4bit-enable-register 7 4
connect r.i6 trigger.o1
connect r.i7 1

define sa 4bit-adder-subtractor 9 1
connect sa.i1 r.o1
connect sa.i2 r.o2
connect sa.i3 r.o3
connect sa.i4 r.o4
connect sa.i5 1
connect sa.i6 0
connect sa.i7 0
connect sa.i8 0
connect sa.i9 0

connect r.i1 sa.o1
connect r.i2 sa.o2
connect r.i3 sa.o3
connect r.i4 sa.o4

alias o1 r.o1
alias o2 r.o2
alias o3 r.o3
alias o4 r.o4
