package main

import "fmt"

/*
 * 嵌入结构模拟 Java 面向对象中的继承
 * 在 GO 中用【组合】实现 Java 中的【继承】功能
 * 借助匿名字段实现上述功能
 */

type human struct {
	Sex int
}

type teacher struct {
	human
	Name string
	Age  int
}

type student struct {
	human
	Name string
	Age  int
}

func main() {
	a := teacher{Name: "teacher_wang", Age: 66}
	b := student{Name: "student_zhou", Age: 22}
	fmt.Println(a, b)

	a = teacher{Name: "teacher_wang", Age: 66, human: human{Sex: 0}}
	b = student{Name: "student_zhou", Age: 22, human: human{Sex: 1}}
	fmt.Println(a, b)

	a.Name = "teacher_liu"
	a.Age = 99
	a.Sex = 1
	//a.human.Sex = 0
	fmt.Println(a)
}
