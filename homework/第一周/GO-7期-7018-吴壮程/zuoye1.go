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
		c = c >> 1
	}
	return sb.String()
}

func main() {
	fmt.Println(BinaryFormat(3))
	fmt.Println(BinaryFormat(4))
	fmt.Println(BinaryFormat(-120))

}

// 1、可以考虑格式化输出
// 2、注意下代码层级
// 3、逻辑正确，建议只提交作业要求的即可
