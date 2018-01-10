package main

import "fmt"

/*
 * 枚举
 */

const (
	a = 'A'
	b
	c = iota
	d
)

const (
	e = iota
)

func main() {
	fmt.Println(a, b, c, d)
	fmt.Println(e)
}
