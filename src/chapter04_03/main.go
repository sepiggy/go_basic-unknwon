package main

import "fmt"

/*
 * 常量的初始化规则与枚举
 */

const (
	a = "123"
	b = len(a)
	c

	d, e = 1, "2"
	f, g
)

func main() {
	fmt.Println(a, b, c)
	fmt.Println(d, e, f, g)
}
