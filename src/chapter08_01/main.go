package main

import "fmt"

func main() {
	var m map[int]string
	m = map[int]string{}
	fmt.Println(m)

	m = make(map[int]string)
	fmt.Println(m)

	m1 := make(map[int]string)
	fmt.Println(m1)

	m[1] = "OK"
	fmt.Println(m)

	delete(m, 1)

	a := m[1]
	fmt.Println(a)
}
