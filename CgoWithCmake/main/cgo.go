package main

import (
	"fmt"
)

/*
#cgo CFLAGS: -g -Wall
#cgo CFLAGS: -I ${SRCDIR}/../CPyEngine/include
#cgo LDFLAGS: -L ${SRCDIR}/../CPyEngine/libs -ldl -lCPyEngine_library
#include "CPyEngineIntf.h"
*/
import "C"

//export GetCmd
func GetCmd() *C.char {
	cmd := <-cmdCh
	return C.CString(cmd)
}

//export GoPrint
func GoPrint(cStr *C.cchar_t) {
	fmt.Println(C.GoString(cStr))
}

var cmdCh chan string
// #cgo CFLAGS: -I ${SRCDIR}/../CPyEngine/extern/pybind11/include/pybind11
//#cgo LDFLAGS: -lstdc++ -lc++
//#cgo LDFLAGS: -ldl
func GetInput() string {
	var inStr string
	_, err := fmt.Scanln(&inStr)
	if nil == err {
		return inStr
	}
	return ""
}

func StartEngine() {
	C.StartPyEngine()
}

func main() {
	cmdCh = make(chan string)
	go StartEngine()
	for {
		cmd := GetInput()
		cmdCh <- cmd
	}
}