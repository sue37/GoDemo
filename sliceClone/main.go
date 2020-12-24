package main

import "fmt"

type Person struct {
	Name string
	Age  uint32
}

type Family struct {
	dad map[string]*Person
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	var b []int
	b = a[:]
	a[2] = 100
	b = append(b[:1], b[3:]...)
	println(b)

	mapA := make(map[string]*Person)
	mapA["WN"] = &Person{"WN", 40}

	f := Family{
		dad: mapA,
	}

	mapA["WN"].Age = 41
	mapA = make(map[string]*Person)
	fmt.Println(f)
}
