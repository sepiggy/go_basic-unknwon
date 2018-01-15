package main

import "fmt"

type A struct {
	Name string
}

type B struct {
	Name string
}

func main() {
	a := A{}
	a.Name = "A"
	fmt.Println(a)
	// 不管是值调用还是指针调用的形式都是相同的
	a.change()
	fmt.Println(a)

	b := B{}
	b.Name = "B"
	fmt.Println(b)
	b.change()
	fmt.Println(b)
}

/*
 * Receiver 也是参数，既然是参数就涉及是以值传递还是以指针传递
 * 与普通参数的用法一致
 */
func (self *A) change() {
	self.Name = "AA"
}

func (self B) change() {
	self.Name = "BB"
}
