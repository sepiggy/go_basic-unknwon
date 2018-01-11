package main

import "fmt"

func main() {
	b := [...]int{5, 2, 6, 3, 9}
	fmt.Println(b)

	num := len(b)
	for i := 0; i < num; i++ {
		for j := i + 1; j < num; j++ {
			if b[i] < b[j] {
				temp := b[i]
				b[i] = b[j]
				b[j] = temp
			}
		}
	}
	fmt.Println(b)
}
