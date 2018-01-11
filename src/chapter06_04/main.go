package main

import "fmt"

func main() {
	a := [10]int{}
	a[1] = 2
	fmt.Println(a)

	// 数组指针支持索引取值
	p := new([10]int)
	p[1] = 2
	fmt.Println(*p)
}
