package main

import "fmt"

/*
 * 作业
 */

type TZ int

func main() {
	var a TZ
	a.Increase(100)
	fmt.Println(a)
}

func (self *TZ) Increase(num int) {
	*self += TZ(num)
}
