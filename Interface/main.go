package main

import "fmt"

type ITest interface {
	GetName() string
}

type ClsA struct {
	name string
	age  int
}

func (a *ClsA) GetName() string {
	return a.name
}

type ClsB struct {
	name   string
	gender string
}

func (b *ClsB) GetName() string {
	return b.name
}

func main() {
	intfMap := make(map[string]interface{})
	clsA := new(ClsA)
	clsA.name = "A"
	clsA.age = 12
	clsB := new(ClsB)
	clsB.name = "B"
	clsB.gender = "male"
	intfMap["A"] = clsA
	intfMap["B"] = clsB
	for _, obj := range intfMap {
		switch obj.(type) {
		case *ClsA:
			ptr := obj.(*ClsA)
			fmt.Println(ptr.age)
		case *ClsB:
			ptr := obj.(*ClsB)
			fmt.Println(ptr.gender)
		default:
			fmt.Println(obj)
		}
	}
}
