package main

import "fmt"

func main() {
	f := closure(100)
	fmt.Println(f(1))
	fmt.Println(f(2))
}

// 用到匿名函数和闭包的经典例子
func closure(x int) (func(int) int) {
	fmt.Printf("%p\n", &x)
	return func(y int) int {
		fmt.Printf("%p\n", &x)
		return x + y
	}
}
