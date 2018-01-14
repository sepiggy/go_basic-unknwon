package main

import "fmt"

func main() {
	a, b, c := A()
	fmt.Println(a, b, c)
}

// 命名返回值
func A() (a, b, c int) {
	a, b, c = 1, 2, 3
	return a,b,c
}
