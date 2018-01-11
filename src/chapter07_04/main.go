package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	//fmt.Println(len(a))
	//fmt.Println(cap(a))

	s1 := a[2:5]
	s2 := a[1:3]
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s2))
	s2 = append(s2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1)

	s1[0] = 111
	fmt.Println(s1)
	fmt.Println(s2)
}
