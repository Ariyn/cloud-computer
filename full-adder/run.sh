#!/bin/bash

bin=/home/ariyn/go/src/github.com/ariyn/cloud-computer/bin

(bin/xor -name xor1 -i1 input_1 -i2 input_2 && wait) & \
(bin/and -name and1 -i1 input_1 -i2 input_2 && wait) & \
(bin/xor -name xor2 -i1 xor1_output_1 -i2 input_3 && wait) & \
(bin/and -name and2 -i1 xor1_output_1 -i2 input_3 && wait) & \
(bin/or -name or1 -i1 and1_output_1 -i2 and2_output_1 && wait)
# xor2 = s
# or1 = c