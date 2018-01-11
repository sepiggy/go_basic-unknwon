package main

import "fmt"

func main() {
	a := 1
	a++
	a--
	p := &a
	fmt.Println(*p)
}
