package main

import (
	"fmt"
	"math"
	"strings"
)

func BinaryFormat(n int32) string {
	a := uint32(n)
	sb := strings.Builder{}
	c := uint32(math.Pow(2, 31))
	for i := 0; i < 32; i++ {
		if a&c != 0 {
			sb.WriteString("1")
		} else {
			sb.WriteString("0")
		}
		c >>= 1
	}
	return sb.String()
}

func main() {
	fmt.Println(BinaryFormat(-1))
	fmt.Println(BinaryFormat(1))
	fmt.Println(BinaryFormat(0))
	fmt.Println(BinaryFormat(260))
	fmt.Println(BinaryFormat(-260))

}

// 考察进制的转换，知识点涉及到位运算和逻辑运算，可以格式化输出下
