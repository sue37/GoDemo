#!/bin/bash

CGO_ENABLED=1 go build -buildmode=c-shared ./Lib/lib.go
go build ./main/main.go

rm ./lib.dll
