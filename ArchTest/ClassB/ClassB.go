package ClassB

import (
	"fmt"
	"Test/ArchTest/Intf"
)

type ClassB struct {
	fac Intf.IFactory
}

var __instance *ClassB

func Create(factory Intf.IFactory) *ClassB {
	__instance = new(ClassB)
	__instance.fac = factory
	factory.Register(__instance)
	return __instance
}

func (in *ClassB) TestB1() {
	fmt.Println("TestB1")
}

func (in *ClassB) TestB2() {
	fmt.Println("TestB2")
}