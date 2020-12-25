package main

import "fmt"

type Person struct {
	Name string
	Age  uint32
}

type Family struct {
	dad map[string]*Person
	lst []*Person
}

func Test(lst []*Person) {
	lstNew := lst
	lstNew = append(lstNew[:0], lstNew[1:]...)
	fmt.Println(lst)
}

func Test1(lst []*Person) {
	var lstNew []*Person
	lstNew = append(lstNew, lst...)
	lstNew = append(lstNew[:0], lstNew[1:]...)
	fmt.Println(lst)
	fmt.Println(lstNew)
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	var b []int
	b = a[:]
	a[2] = 100
	b = append(b[:1], b[3:]...)
	println(b)

	f1 := new(Family)
	p1 := &Person{"Andy", 40}
	p2 := &Person{"Susan", 12}
	p3 := &Person{"Bob", 20}
	f1.lst = append(f1.lst, p1)
	f1.lst = append(f1.lst, p2)
	f1.lst = append(f1.lst, p3)

	var lst []*Person = nil
	lst = append(lst, f1.lst[0])
	lst = append(lst, f1.lst[1])

	lstNew := lst
	lstNew = append(lstNew[:0], lstNew[1:]...)
	fmt.Println(lst)

	Test1(lst)
	Test(lst)
	lst = append(lst, f1.lst[2])

	mapA := make(map[string]*Person)
	mapA["WN"] = p1

	f := Family{
		dad: mapA,
	}

	mapA["WN"].Age = 41
	mapA = make(map[string]*Person)
	fmt.Println(f)
}
