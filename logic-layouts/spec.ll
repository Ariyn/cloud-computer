# ll file will describe single Logical logic gates Group(LG)
# LG will work like single logic gate
# but it still has multiple gates in it.
# everything after sharp is comment, does nothing

inputs {size}
# this will describe size of input for LG
# input count starts from 1
# if you used `input 2`, there is 2 inputs; inputs.1, inputs.2
# you can use inputs command once for single file
# if no inputs, default size is 2

define {name} {logic gate} {input size} {output size} [no-optimization]
# input_size of inputs will be created.
# each inputs will be named as i1, i2, i3...
# output is same but o1, o2, o3...
# "no-optimization" parameter is optional. When you use "no-optimization", gates will publish their outputs even it ditn't changed.

define {name} {Logical logic gates Group}
# you can define another LG, except it self.
# recursive define is not allowed.
# you can not set input size or output size for LG
# LG file has it's own input size in it. if not, it's input size is 2.
# there is no output in LG. you can export names which output or alias

# can access to specific input or output with dot
# ex) logic_gate1.i1, and1.o2, other_lg.output

connect {input or output} {other input or output or static value}

# connect will connect 2 input or output
# if first one is input, second should be output and vice versa.
# connect can occur infinite loop or oscillation, be careful.

# For default inputs, you can use "$inputs.{i}"

# If you use 0 or 1 for static value, that will be fixed value. 

alias {unique name} {input or output or static value}
# alias for specific input or output, which is just sugar syntax
# name should be unique in single ll file

# you can also access other LG's alias.
