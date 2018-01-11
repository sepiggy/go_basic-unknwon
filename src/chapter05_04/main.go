package main

import "fmt"

func main() {
	a := "string"
	len := len(a)

	for i := 0; i < len; i++ {
		fmt.Println(string(a[i]))
	}

	for _, v := range a {
		fmt.Println(string(v))
	}
}
