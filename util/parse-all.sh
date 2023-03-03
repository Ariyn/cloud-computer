#!/bin/bash

for i in logic-layouts/*; do

	if [[ $i == "logic-layouts/spec.ll" ]]; then
		continue
	fi

	util/parse $i;
done
