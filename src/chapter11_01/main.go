package main

import "fmt"

type A struct {
	Name string
}

type B struct {
	Name string
}

func main() {
	a := A{Name: "Name A"}
	a.Print()

	b := B{Name: "Name B"}
	b.Print()
}

// 方法与类型相绑定, 即使方法名相同，只要类型名不同即可
func (a A) Print() {
	fmt.Println(a.Name)
}

func (b B) Print() {
	fmt.Println(b.Name)
}
