package main

import "fmt"

type person struct {
	Name string
	Age  int
}

func main() {
	a := person{}
	a.Name = "Joe"
	a.Age = 19
	fmt.Println(a)

	// 使用字面值初始化，类似于其他语言中的构造函数
	b := person{
		Name: "Mary",
		Age:  22,
	}
	fmt.Println(b)
}
