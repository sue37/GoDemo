package ClassA

import (
	"fmt"
	"Test/ArchTest/Intf"
)

type ClassA struct {
	fac Intf.IFactory
}

var __instance *ClassA

func Create(factory Intf.IFactory) *ClassA {
	__instance = new(ClassA)
	__instance.fac = factory
	factory.Register(__instance)
	return __instance
}

func (in *ClassA) TestA1() {
	fmt.Println("TestA1")
}

func (in *ClassA) TestA2() {
	fmt.Println("TestA2")
	in.fac.GetClassB().TestB1()
}

