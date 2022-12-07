#!/bin/bash

bin=/home/ariyn/go/src/github.com/ariyn/cloud-computer/bin

(bin/xor -i1 input_1 -i2 input_2 && wait & bin/and -i1 input_1 -i2 input_2 && wait)
