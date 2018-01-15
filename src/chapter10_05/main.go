package main

import "fmt"

/*
 * 结构类型之间的相互赋值和比较
 */

type person struct {
	Name string
	Age  int
}

func main() {
	p1 := person{Name: "joe", Age: 19}
	var p2 person
	p2 = p1
	fmt.Println(p2)
	fmt.Println(p1 == p2)
}
