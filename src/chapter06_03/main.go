package main

import "fmt"

func main() {
	a := [...]int{1, 2}
	b := [...]int{1, 2}
	c := [...]int{3, 4}
	fmt.Println(a == b)
	fmt.Println(a != c)
}
