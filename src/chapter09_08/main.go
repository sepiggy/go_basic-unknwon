package main

import "fmt"

// defer 与匿名函数结合, 三次打印都是 3
func main() {
	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
}
