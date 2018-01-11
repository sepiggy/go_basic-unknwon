package main

import "fmt"

func main() {
	if 1 < 2 {
		fmt.Println("hello world")
	}

	a := 10
	if a := 1; a >= 1 {
		fmt.Println(a)
	}
	fmt.Println(a)
}
