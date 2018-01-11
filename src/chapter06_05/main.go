package main

import "fmt"

func main() {
	a := [2][3]int{
		{1, 1, 1},
		{2, 2, 2},
	}
	fmt.Println(a)

	b := [2][3]int{
		{1: 1},
		{2: 2},
	}
	fmt.Println(b)

	c := [...][3]int{
		{1, 2, 3}, {4, 5, 6},
	}
	fmt.Println(c)
}
