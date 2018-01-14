package main

import (
	"fmt"
)

// defer: 后定义先调用, 逆序向上调用
func main() {
	fmt.Println("a")
	defer fmt.Println("b")
	defer fmt.Println("c")
}
