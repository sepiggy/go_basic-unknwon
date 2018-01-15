package main

import "fmt"

/*
 * struct 是值类型
 */
type Person struct {
	Name string
	Age  int
}

func main() {
	// 初始化时可以直接保存为指针
	//【最佳实践】
	// 在对结构进行初始化的时候，通常前面使用取地址符号 &，将其赋值给变量；
	// 通过结构指针对结构进行操作，是最简洁高效的方式
	// 在 GO 中可以直接通过 . 操作符直接操作指针，不需要通过 * 解引用
	p := &Person{
		Name: "JMS",
		Age:  28,
	}
	fmt.Println(*p)
	A(p)
	fmt.Println(*p)
}

func A(person *Person) {
	person.Age = 100
	fmt.Println("In func A", *person)
}
