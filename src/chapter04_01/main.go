package main

import (
	"strconv"
	"fmt"
)

func main() {
	var a int = 65

	b := string(a)
	fmt.Println(b)

	c := strconv.Itoa(a)
	fmt.Println(c)

	d, _ := strconv.Atoi(c)
	fmt.Println(d)

}
