package main

import "fmt"

func Test(parLst... interface{}) {
	fmt.Println("len of input par is:", len(parLst))
	for _, par := range parLst{
		fmt.Println(par)
	}
}

func Test1(par... interface{}) {
	Test(par...)
}

func main() {
	Test1(123, "hello")
}
