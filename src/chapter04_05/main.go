package main

import (
	"fmt"
)

func main() {
	fmt.Println(1 << 10 << 10 << 10 >> 10 >> 10)
	fmt.Println(6 & 11)
	fmt.Println(6 | 11)
	fmt.Println(6 ^ 11)
	fmt.Println(6 &^ 11)
}
