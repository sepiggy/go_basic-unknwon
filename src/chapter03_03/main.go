package main

import "fmt"

/*
 * 类型别名
 */
type (
	byte uint8
	rune int32
	文本 string
	ByteSize int64
)

func main() {
	var s 文本
	s = "中文类型名"
	fmt.Println(s)
}
