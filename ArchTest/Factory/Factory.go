package Factory

import (
	"Test/ArchTest/Intf"
	"Test/ArchTest/ClassA"
	"Test/ArchTest/ClassB"
)

type Factory struct {
	IfClassA Intf.IClassA
	IfClassB Intf.IClassB
}

var __factory *Factory

func init() {
	__factory = new(Factory)
	__factory.IfClassA = ClassA.Create(__factory)
	__factory.IfClassB = ClassB.Create(__factory)
}

func Inst() *Factory {
	return __factory
}

//func (f *Factory)ClassA() Intf.IClassA {
//	return
//}

func (f *Factory) Register(imp interface{}) {
	switch imp.(type) {
	case ClassA.ClassA:
		f.IfClassA = imp.(Intf.IClassA)
	case ClassB.ClassB:
		f.IfClassB = imp.(Intf.IClassB)
	}
}

func (f *Factory) GetClassA() Intf.IClassA{
	return f.IfClassA
}

func (f *Factory) GetClassB() Intf.IClassB{
	return f.IfClassB
}
