package main

import "fmt"

type TZ int

func main() {
	var a TZ = 1
	a.Print()
}

// 可以为任意一种类型添加方法
func (self *TZ) Print() {
	fmt.Println("TZ")
}
