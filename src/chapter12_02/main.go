package main

import "fmt"

/*
 * 嵌入接口实现 JAVA 中接口的继承
 */

type Connecter interface {
	Connect()
}

type USB interface {
	Name() string
	// 嵌入接口
	Connecter
}

type PhoneConnecter struct {
	name string
}

func (self PhoneConnecter) Name() string {
	return self.name
}

func (self PhoneConnecter) Connect() {
	fmt.Println("Connected ...")
}

func main() {
	phoneConnecter := PhoneConnecter{"iPhone"}
	phoneConnecter.Connect()
	fmt.Println(phoneConnecter.Name())
	Disconnect(phoneConnecter)
}

// 空接口
func Disconnect(connecter interface{}) {
	// 类型断言
	// if pc, ok := connecter.(PhoneConnecter); ok {
	//	fmt.Println("Disconnected ...", pc.name)
	//	return
	//}
	//fmt.Println("Unknown device ...")

	// type-switch
	switch v := connecter.(type) {
	case PhoneConnecter:
		fmt.Println("Disconnected", v.name)
	default:
		fmt.Println("unknown device.")
	}
}
