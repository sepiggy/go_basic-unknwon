package main

import "fmt"

/*
 * 常量定义
 */

const a int = 1
const b = 'A'

const (
	c, d, e = a, 2, 3
)

func main() {
	fmt.Println(a, b, c, d, e)
}
