package main

import "fmt"

func main() {
	a := 1
	switch {
	case a >= 0:
		fmt.Println("a>=0")
		fallthrough // 继续检查下面的条件
	case a >= 1:
		fmt.Println("a>=1")
	default:
		fmt.Println("None")
	}
}
