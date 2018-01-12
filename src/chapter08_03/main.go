package main

import "fmt"

func main() {
	// 以 map[int]string 为元素的 slice
	sm := make([]map[int]string, 5)

	// 如果对被 range 对象直接操作需要使用 key 直接操作
	for i := range sm {
		sm[i] = make(map[int]string, 1)
		sm[i][1] = "OK"
		fmt.Println(sm[i])
	}

	fmt.Println(sm)
}
