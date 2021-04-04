package main

/*
#include <stdio.h>
#include <stdlib.h>
#include "libhello.h"
#cgo CFLAGS: -I./
#cgo LDFLAGS: -lhello -Wl,-rpath=./
*/
import "C"
import "fmt"

func main() {
	C.Test()
	str := C.Hello()
	fmt.Println(str)
}
