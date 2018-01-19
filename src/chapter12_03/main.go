package main

import "fmt"

/*
 * 接口转换
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

type TVConnecter struct {
	name string
}

func (self TVConnecter) Connect() {
	fmt.Println(self.name, "is connected ...")
}

func (self PhoneConnecter) Connect() {
	fmt.Println(self.name, "is connected ...")
}

func (self PhoneConnecter) Name() string {
	return self.name
}

func main() {
	pc := PhoneConnecter{"PhoneConnecter"} // pc is a struct
	var a Connecter                        // a is a interface
	a = Connecter(pc)                      // convert a struct to interface
	a.Connect()

	// tv:=TVConnecter{"TVConnecter"}

}

func Disconnect(object Object) {
	switch v := object.(type) {
	case PhoneConnecter:
		fmt.Println("Disconnected:", v.name)
	default:
		fmt.Println("Unknown device.")
	}
}
