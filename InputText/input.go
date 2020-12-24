package main

import (
	"bufio"
	"fmt"
	"os"
)

type test struct {
	exitFlg bool
	ctrlFlg bool
	cmdBuf []byte
}

func (inst test) Test () {
	inst.cmdBuf = make([]byte, 0)
	reader := bufio.NewReader(os.Stdin)
	inst.exitFlg = false
	byt := make([]byte, 1)
	for !inst.exitFlg {
		//_, err := os.Stdin.Read(byt)
		_, err := reader.Read(byt)
		if err != nil {
			return
		}
		if '\b' == byt[0] {
			fmt.Println('\b')

		} else if '\n' == byt[0] {

		}
		inst.cmdBuf = append(inst.cmdBuf, byt[0])
		fmt.Println(byt)
		fmt.Print(inst.cmdBuf)
	}
}

func main() {
	var inst test
	inst.Test()
}
