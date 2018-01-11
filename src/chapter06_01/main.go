package main

import "fmt"

func main() {
	var a [2]int // 长度为 2 的 int 数组
	fmt.Println(a)

	var b [2]int
	b = a
	fmt.Println(b)

	var c [2]string
	fmt.Println(c)

	d := [2]int{1, 1}
	fmt.Println(d)

	e := [2]int{1}
	fmt.Println(e)

	f := [20]int{19: 1}
	fmt.Println(f)

	g := [...]int{1, 2, 3, 4, 5}
	fmt.Println(len(g))

	h := [...]int{0: 1, 1: 2, 2: 3}
	fmt.Println(h)

	i := [...]int{
		19: 1,
	}
	fmt.Println(i)

	j := [...]int{99: 1}
	fmt.Println(j)

}
