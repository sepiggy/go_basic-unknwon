package main

import "fmt"

func main() {
	a := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k'}
	sa := a[2:5]
	fmt.Println(len(sa), cap(sa))

	sb := sa[3:5]
	fmt.Println(len(sb), cap(sb))

	fmt.Println(string(sa))
	fmt.Println(string(sb))
}
