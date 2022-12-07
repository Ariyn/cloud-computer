#!/bin/bash

(bin/nor -name r_nor -i1 input_r -i2 s_nor_output_1 && wait) & \
(bin/nor -name s_nor -i1 input_s -i2 r_nor_output_1 && wait)