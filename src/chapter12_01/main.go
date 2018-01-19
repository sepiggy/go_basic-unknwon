package main

import "fmt"

// 接口只有方法的声明
type USB interface {
	Name() string
	Connect()
}

// Golang 中结构的属性和方法是分开定义的, 方法不写到结构里面，需要单独定义
// 结构中只包含字段属性的定义
type PhoneConnecter struct {
	name string
}

// PhoneConnecter 结构的方法定义
func (self PhoneConnecter) Name() string {
	return self.name
}

func (self PhoneConnecter) Connect() {
	fmt.Println("Connect:", self.name)
}

func Disconnect(usb USB) {
	fmt.Println("Disconnected ...")
}

func main() {
	u := PhoneConnecter{}
	u.name = "SONY"
	u.Connect()
	Disconnect(u)
}
