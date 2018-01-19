package main

import (
	"reflect"
	"fmt"
)

type User struct {
	Id   int
	Name string
	Age  int
}

type Manager struct {
	User
	title string
}

func main() {
	m := Manager{User: User{1, "Maria", 12}, title: "123"}
	t := reflect.TypeOf(m)

	// 通过反射获得匿名字段
	fmt.Printf("%#v\n", t.Field(1))
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 0}))
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 1}))
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 2}))
}
