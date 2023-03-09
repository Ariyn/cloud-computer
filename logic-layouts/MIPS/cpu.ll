inputs 1

# inputs.1 = clock

define PC MIPS/counter 35 32
connect PC.i34 $inputs.1
connect PC.i35 1

define MEMORY MIPS/ram
connect MEMORY.i35 PC.o1
connect MEMORY.i36 PC.o2
connect MEMORY.i37 PC.o3
connect MEMORY.i38 PC.o4