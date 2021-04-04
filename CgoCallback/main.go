package main

import (
	"fmt"
)

/*
#include <stdio.h>
extern char * GetCmd();
extern void PrintInfo(char *);

static inline void StartProc() {
	PrintInfo("Enter StartProc");
	while (1) {
		char *cmdStr;
		cmdStr = GetCmd();
		PrintInfo(cmdStr);
	}
}
*/
import "C"

//export GetCmd
func GetCmd() *C.char {
	fmt.Println("Enter GetCmd")
	cmd := <-cmdCh
	fmt.Println("Got cmd [" + cmd + "] in GetCmd")
	return C.CString(cmd)
}

//export PrintInfo
func PrintInfo(cInfo *C.char) {
	fmt.Println("Enter PrintInfo")
	info := C.GoString(cInfo)
	fmt.Println("PrintInfo:[" + info + "]")
}

var cmdCh chan string

func GetInput() string {
	var inStr string
	_, err := fmt.Scanln(&inStr)
	if nil == err {
		return inStr
	}
	return ""
}

func StartEngine() {
	fmt.Println("Enter StartEngine")
	C.StartProc()
}

func main() {
	fmt.Println("Enter Main")
	cmdCh = make(chan string)
	go StartEngine()
	for {
		cmd := GetInput()
		fmt.Println("Got input cmd [" + cmd + "]")
		cmdCh <- cmd
	}
}
