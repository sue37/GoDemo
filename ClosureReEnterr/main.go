package main

import "fmt"

type TypeFnCallBackGen func() string

func GenFn(name string) TypeFnCallBackGen {
	greet := "Hello " + name
	return func() string {
		return greet + ", this is GenFn"
	}
}

func main() {
	f1 := GenFn("Dad")
	f2 := GenFn("Mod")
	fmt.Println(f1())
	fmt.Println(f2())
}
