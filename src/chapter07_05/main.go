package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := []int{7, 8, 9, 10, 11, 12, 13, 14}
	//copy(s1, s2)
	copy(s2[5:8], s1[1:3])
	//fmt.Println(s1)
	fmt.Println(s2)
}
