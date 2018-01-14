package main

import "fmt"

func main() {
	// a 是 func 类型，b 是 int 类型
	a := A
	b := A()
	fmt.Println(b)
	a()
}

func A() int {
	fmt.Println("func A")
	return 1
}
