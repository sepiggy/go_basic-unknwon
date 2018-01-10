package main

import (
	"fmt"
	"math"
)

/*
 * 类型的零值
 */
func main() {
	var a int32
	fmt.Println(a)

	var b float64
	fmt.Println(b)

	var c bool
	fmt.Println(c)

	var d string
	fmt.Println(d)

	var e []int
	fmt.Println(e)

	var f [100]bool
	fmt.Println(f)

	var g [1]byte
	fmt.Println(g)

	fmt.Println(math.MinInt8)
	fmt.Println(math.MaxInt32)
}
