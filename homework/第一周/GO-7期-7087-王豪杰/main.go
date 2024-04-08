package main

import (
	"fmt"
	"math"
	"strings"
)

//32位运算
func binaryFormat(n int32) string {
	Source := uint32(n)
	strBuiuld := strings.Builder{}
	Dest := uint32(math.Pow(2, 31))

	for i := 0; i < 32; i++ {
		if Source&Dest != 0 {
			strBuiuld.WriteString("1")
		} else {
			strBuiuld.WriteString("0")
		}
		Dest >>= 1
	}
	return strBuiuld.String()

}

func main() {

	binary := binaryFormat(7)
	binary = binaryFormat(-7)
	binary = binaryFormat(260)
	fmt.Println(binary)
	binary = binaryFormat(-260)
	fmt.Println(binary)
	binary = binaryFormat(^260)
	fmt.Println(binary)
}

// 输出打印时，可以实现格式化输出
