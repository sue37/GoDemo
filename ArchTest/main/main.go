package main

import (
	"Test/ArchTest/Factory"
)
func main() {
	f := Factory.Inst()
	f.IfClassA.TestA1()
	f.IfClassA.TestA2()
}

