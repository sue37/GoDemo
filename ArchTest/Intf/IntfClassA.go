package Intf

type IFactory interface {
	Register(interface{})
	GetClassA() IClassA
	GetClassB() IClassB
}


type IClassA interface {
	TestA1()
	TestA2()
}

type IClassB interface {
	TestB1()
	TestB2()
}
