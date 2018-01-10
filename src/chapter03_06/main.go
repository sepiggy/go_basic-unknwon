package main

import "fmt"

/*
 * 变量的类型转换
 */
func main() {
	var a float32 = 100.1
	fmt.Println(a)
	b := int(a)
	fmt.Println(b)
}
