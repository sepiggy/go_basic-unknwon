package main

import "fmt"

type USB interface {
	Name() string
	Connect()
}

type PhoneConnecter struct {
	name string
}

func main() {
	u := PhoneConnecter{"SONY"}
	fmt.Println(u.Name())
	u.Connect()
	Disconnect(u)
}

func (self PhoneConnecter) Name() string {
	return self.name
}

func (self PhoneConnecter) Connect() {
	fmt.Println("Connect phone ...")
}

func Disconnect(usb USB) {
	fmt.Println("Disconnected ...")
}
