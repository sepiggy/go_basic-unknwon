package main

import "fmt"

/*
 * 匿名函数
 */
func main() {
	a := func() {
		fmt.Println("匿名函数1。。。")
	}
	a()

	func() {
		fmt.Println("匿名函数2。。。")
	}()

	r := test(func(s string) string {
		return s
	})
	fmt.Println(r)
}

func test(f func(s string) string) string {
	return f("匿名函数3。。。")
}
