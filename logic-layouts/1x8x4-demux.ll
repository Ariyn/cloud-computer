inputs 9

define not-selector not 1 1
connect not-selector.i1 $inputs.9

define and1 and 2 1
connect and1.i1 $inputs.1
connect and1.i2 $inputs.9

define and2 and 2 1
connect and2.i1 $inputs.2
connect and2.i2 $inputs.9

define and3 and 2 1
connect and3.i1 $inputs.3
connect and3.i2 $inputs.9

define and4 and 2 1
connect and4.i1 $inputs.4
connect and4.i2 $inputs.9

define and5 and 2 1
connect and5.i1 $inputs.5
connect and5.i2 not-selector.o1

define and6 and 2 1
connect and6.i1 $inputs.6
connect and6.i2 not-selector.o1

define and7 and 2 1
connect and7.i1 $inputs.7
connect and7.i2 not-selector.o1

define and8 and 2 1
connect and8.i1 $inputs.8
connect and8.i2 not-selector.o1

define or1 or 2 1
connect or1.i1 and1.o1
connect or1.i2 and5.o1

define or2 or 2 1
connect or2.i1 and2.o1
connect or2.i2 and6.o1

define or3 or 2 1
connect or3.i1 and3.o1
connect or3.i2 and7.o1

define or4 or 2 1
connect or4.i1 and4.o1
connect or4.i2 and8.o1

alias o1 or1.o1
alias o2 or2.o1
alias o3 or3.o1
alias o4 or4.o1