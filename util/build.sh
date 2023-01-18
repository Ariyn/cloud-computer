#!/bin/bash

go build -o bin/not ./app/not
go build -o bin/or ./app/or
go build -o bin/and ./app/and
go build -o bin/xor ./app/xor
go build -o "bin/alias" "./app/alias"
go build -o bin/flipflop ./app/flipflop
go build -o bin/input ./app/input
go build -o bin/parser ./parser
