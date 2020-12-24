#!/bin/bash

export GOPATH=/mnt/e/WorkSpace/golang
GOOS=linux GOARCH=arm64 go build -o /mnt/e/WorkSpace/golang/CmdParserMain/input input.go