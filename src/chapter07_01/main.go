package main

import "fmt"

func main() {
	var s1 []int
	fmt.Println(s1)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))

	// 通过底层数组创建 slice (reslice)
	a := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(a)
	s2 := a[5:]
	fmt.Println(s2)
	s3 := a[:5]
	fmt.Println(s3)

	// 通过 make 函数创建 slice
	s4 := make([]int, 3)
	fmt.Println(s4)
	fmt.Println(len(s4))
	fmt.Println(cap(s4))


}
