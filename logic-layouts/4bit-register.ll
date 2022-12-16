inputs 4

define latch1 sr-latch 2 1
define latch2 sr-latch 2 1
define latch3 sr-latch 2 1
define latch4 sr-latch 2 1

connect latch1.i2 $inputs.1
connect latch2.i2 $inputs.2
connect latch3.i2 $inputs.3
connect latch4.i2 $inputs.4

