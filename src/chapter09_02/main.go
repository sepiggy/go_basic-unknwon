package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3, 4}
	A(s[0], s[1], s[2], s[3])
	fmt.Println(s)
}

// 引用类型和值类型作为形参的区别,
// 变长参数的本质是切片
func A(s ...int) {
	s[0] = 999
	fmt.Println(s)
}
