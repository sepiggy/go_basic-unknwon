package main

import (
	"fmt"
)

/*
 * 对象赋值给接口类似于函数传参，有一个【拷贝】的过程
 */

type Object interface{}

type Connecter interface {
	Connect()
}

type USB interface {
	Connecter
	Name() string
}

type PhoneConnecter struct {
	name string
}

func (self PhoneConnecter) Connect() {
	fmt.Println(self.name, "is connected.")
}

func (self PhoneConnecter) Name() string {
	return self.name
}

type TVConnecter struct {
	name string
}

func (self TVConnecter) Connect() {
	fmt.Println(self.name, "is connected.")
}

func main() {
	pc := PhoneConnecter{"PhoneConnecter"}
	var a Connecter
	// 对象赋值给接口
	a = Connecter(pc)
	a.Connect()

	pc.name = "xxxxxxxxxxx"
	a.Connect() // 无视对原对象的修改
}
