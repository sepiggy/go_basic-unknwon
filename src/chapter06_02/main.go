package main

import "fmt"

func main() {
	// 数组的指针
	a := [...]int{99: 1}
	var p *[100]int = &a
	fmt.Println(p)
	fmt.Println(*p)

	// 指针的数组
	x, y := 1, 2
	b := [...]*int{&x, &y}
	fmt.Println(b)
}
