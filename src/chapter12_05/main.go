package main

import "fmt"

//【注意！！！】
// 结构的"方法调用"会做 receiver 的自动转换，
// 但是接口的"方法调用"不会做 receiver 的自动转换！

type Object interface{}

type Runner interface {
	Run()
}

type People struct {
	name string
	age  int
	city string
}

func (self People) Run() {
	fmt.Println(self.name, "is running...")
}

func main() {
	p := People{"Maria", 18, "NY"}
	r := Runner(p)

	p.Run()
	(&p).Run()
	r.Run()
}
