package main

import "fmt"

func main() {
LABEL1:
	for i := 0; i < 10; i++ {
		for {
			continue LABEL1
			fmt.Println("IN")
		}
		fmt.Println("OUT")
	}
	fmt.Println("OK")
}
