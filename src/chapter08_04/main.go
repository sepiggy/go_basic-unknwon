package main

import (
	"fmt"
	"sort"
)

// 实现按 key 的顺序从 map 中取值, 借助引用一个 slice 实现
func main() {
	m := map[int]string{
		1: "a",
		2: "b",
		3: "c",
		4: "d",
		5: "e",
	}

	s := make([]int, len(m))

	i := 0
	for k := range m {
		s[i] = k
		i++
	}
	sort.Ints(s)
	fmt.Println(s)
}
