package main

import "fmt"

// chapter 8 作业
// 将 map[int]string 的 key value 对调，变为 map[string]int 形式

func main() {
	m1 := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}
	fmt.Println(m1)
	m2 := make(map[string]int, 5)

	for k, v := range m1 {
		m2[v] = k
	}
	fmt.Println(m2)
}
