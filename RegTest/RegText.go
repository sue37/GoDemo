package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	echo := "Value 0x00000000 (0) read from address 0x00008100."
	var lowBit uint64 = 0
	var fldLen uint64 = 0
	var updVal uint64 = 0
	var err error
	lowBit, err = strconv.ParseUint("4", 10, 32)
	fmt.Println("lowBit")
	fmt.Println(lowBit)
	if nil != err { }
	fldLen, err = strconv.ParseUint("2", 10, 32)
	fmt.Println(fldLen)
	if nil != err { }
	updVal, err = strconv.ParseUint("3", 10, 32)
	fmt.Println(updVal)
	if nil != err { }
	updVal = (updVal & ((1 << fldLen) - 1)) << lowBit
	mrkVal := uint32(((1 << fldLen) - 1) << lowBit)
	mrkVal = ^mrkVal
	oldVal := uint32(74565)
	oldVal = oldVal & mrkVal
	newVal := oldVal | uint32(updVal)
	fmt.Println(newVal)

	for {
		valExtReg := regexp.MustCompile(`Value\s+0x(?P<VAL_HEX>[\da-fA-F]{8})\s\((?P<VAL_DEC>\d+)\)\sread\sfrom\saddress\s0x(?P<ADD_HEX>[\da-fA-F]{8})`)
		if valExtReg.MatchString(echo) {
			fmt.Println("Match")
		}
	}
}


