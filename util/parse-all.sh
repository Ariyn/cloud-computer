#!/bin/bash

for i in logic-layouts/*; do
echo $i;
util/parse $i;
done
