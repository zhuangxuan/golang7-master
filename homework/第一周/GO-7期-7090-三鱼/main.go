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
		if a&c > 0 {
			sb.WriteString("1")
		} else {
			sb.WriteString("0")
		}
		c >>= 1
	}
	return sb.String()
}

func main() {
	// vit_demo()
	// format()
	fmt.Printf("%v\n", BinaryFormat(0))
	fmt.Printf("%v\n", BinaryFormat(1))
	fmt.Printf("%v\n", BinaryFormat(-1))
	fmt.Printf("%v\n", BinaryFormat(260))
	fmt.Printf("%v\n", BinaryFormat(-260))
}

// 输出打印时，可以实现下格式化输出
