package main

import "fmt"

/*
 * 匿名结构
 */

// 匿名结构嵌套在其它结构之中
type Person struct {
	Name string
	Age  int
	Contact struct {
		Phone, City string
	}
}

/*
 * 匿名结构
 */
func main() {
	/*
	a := &struct {
		Name string
		Age  int
	}{"JMS", 28}
	fmt.Println(*a)
	*/

	/*
	 * 匿名结构的初始化
	 */
	p := Person{
		Name: "Joe",
		Age:  19,
	}
	p.Contact.Phone = "12310243470"
	p.Contact.City = "Beijing"
	fmt.Println(p)

}
