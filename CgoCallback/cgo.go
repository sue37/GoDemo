package main

import (
	"fmt"
)

/*
extern char * GetCmd();

static inline void ProcCmd(char * cMsg) {

}

static inline void StartProc() {
   printf("Enter StartProc");
   while (true) {
      char cmdStr[100];
      cmdStr = GetCmd();
      printf("%s", cmdStr);
   }
}
*/
import "C"

//export GetCmd
func GetCmd() *C.char {
	cmd := <-cmdCh
	return C.CString(cmd)
}

//func GetCmd() string {
// cmd := <-cmdCh
// return cmd
//}

var cmdCh chan string

//type PyEngine struct {
// cmdCh chan string
//}
//
//func (engine *PyEngine) Init() {
//
//}
//
//func (engine *PyEngine) StartEngine() {
// C.StartProc()
//}
//
//func (engine *PyEngine) ProcCmd() {
// var cmd string
// for {
//    cmd = <-engine.cmdCh
//    var cMsg *C.char = C.CString(cmd)
//    C.ProcCmd(cMsg)
//    C.free(unsafe.Pointer(cMsg))
// }
//}

func GetInput() string {
	var inStr string
	_, err := fmt.Scanln(&inStr)
	if nil == err {
		return inStr
	}
	return ""
}

func StartEngine() {
	C.StartProc()
}

func main() {
	cmdCh = make(chan string)
	go StartEngine()
	for {
		cmd := GetInput()
		cmdCh <- cmd
	}
}