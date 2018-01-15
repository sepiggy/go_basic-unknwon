package main

import "fmt"

/*
 * 匿名字段
 */

type Person struct {
	string
	int
}

func main() {
	p := Person{"Jms", 19}
	fmt.Println(p)
}
