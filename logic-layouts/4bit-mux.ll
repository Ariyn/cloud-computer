inputs 7

define a1 and 4 1
define a2 and 4 1
define a3 and 4 1
define a4 and 4 1
define o1 or 4 1

define n1 not 1 1
define n2 not 1 1
define n3 not 1 1
define n4 not 1 1
define n5 not 1 1

connect a1.i1 $inputs.1
connect a2.i1 $inputs.2
connect a3.i1 $inputs.3
connect a4.i1 $inputs.4

connect n1.i1 $inputs.5
connect n3.i1 $inputs.6
connect n5.i1 $inputs.7

connect n2.i1 n1.o1
connect n4.i1 n2.o1

connect a1.i2 n1.o1
connect a2.i2 n1.o1
connect a3.i2 n2.o1
connect a4.i2 n2.o1

connect a1.i3 n3.o1
connect a3.i3 n3.o1

connect a2.i2 n4.o1
connect a4.i2 n4.o1

connect a1.i4 n5.o1
connect a2.i4 n5.o1
connect a3.i4 n5.o1
connect a4.i4 n5.o1

connect o1.i1 a1.o1
connect o1.i2 a2.o1
connect o1.i3 a3.o1
connect o1.i4 a4.o1

alias output o1.o1