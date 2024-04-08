<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [Go语言基础语法](#go%E8%AF%AD%E8%A8%80%E5%9F%BA%E7%A1%80%E8%AF%AD%E6%B3%95)
  - [标识符与关键字](#%E6%A0%87%E8%AF%86%E7%AC%A6%E4%B8%8E%E5%85%B3%E9%94%AE%E5%AD%97)
  - [操作符与表达式](#%E6%93%8D%E4%BD%9C%E7%AC%A6%E4%B8%8E%E8%A1%A8%E8%BE%BE%E5%BC%8F)
  - [变量、常量、字面量](#%E5%8F%98%E9%87%8F%E5%B8%B8%E9%87%8F%E5%AD%97%E9%9D%A2%E9%87%8F)
    - [变量类型](#%E5%8F%98%E9%87%8F%E7%B1%BB%E5%9E%8B)
    - [变量声明](#%E5%8F%98%E9%87%8F%E5%A3%B0%E6%98%8E)
    - [变量初始化](#%E5%8F%98%E9%87%8F%E5%88%9D%E5%A7%8B%E5%8C%96)
    - [常量](#%E5%B8%B8%E9%87%8F)
    - [字面量](#%E5%AD%97%E9%9D%A2%E9%87%8F)
  - [变量作用域](#%E5%8F%98%E9%87%8F%E4%BD%9C%E7%94%A8%E5%9F%9F)
  - [注释与godoc](#%E6%B3%A8%E9%87%8A%E4%B8%8Egodoc)
    - [注释的形式](#%E6%B3%A8%E9%87%8A%E7%9A%84%E5%BD%A2%E5%BC%8F)
    - [注释的位置](#%E6%B3%A8%E9%87%8A%E7%9A%84%E4%BD%8D%E7%BD%AE)
    - [go doc](#go-doc)
    - [godoc](#godoc)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# Go语言基础语法
<!-- 自动生成目录的方法：doctoc Go语言基础语法.md -->
## 标识符与关键字
&#8195;&#8195;go变量、常量、自定义类型、包、函数的命名方式必须遵循以下规则：
1. 首字符可以是任意Unicode字符或下划线。
2. 首字符之外的部分可以是Unicode字符、下划线或数字。
3. 名字的长度无限制。
> 理论上名字里可以有汉字，甚至可以全是汉字，但实际中不要这么做。

Go语言关键字
```
break  default  func  interface  select  case  defer  go  map  struct  chan  else  goto  package  switch  const  if  range  type  continue  for  import  return  fallthrough  var
```
常量
```
true  false  iota  nil   
```
数据类型
```
int  int8  int16  int32  int64  uint  uint8  uint16  uint32  uint64  uintptr  float32  float64  complex128  complex64  bool  byte  rune  string  error
```
函数
```
make  len  cap  new  append  copy  close  delete  complex  real  imag  panic  recover
```
## 操作符与表达式
算法术运算符
| 运算符 | 描述 |
| :---: | :---: |
| + | 相加 |
|-|相减|
|*|相乘|
|/|相除|
|%|求余|  
```Go
//arithmetic 算术运算
func arithmetic() {
	var a float32 = 8
	var b float32 = 3
	var c float32 = a + b
	var d float32 = a - b
	var e float32 = a * b
	var f float32 = a / b
	fmt.Printf("a=%.3f, b=%.3f, c=%.3f, d=%.3f, e=%.3f, f=%.3f\n", a, b, c, d, e, f)
}
```
关系运算符  

|运算符|描述|
| :---: | :--- |
|==|检查两个值是否相等，如果相等返回 True 否则返回 False|
|!=|检查两个值是否不相等，如果不相等返回 True 否则返回 False|
|>|检查左边值是否大于右边值，如果是返回 True 否则返回 False|
|>=|检查左边值是否大于等于右边值，如果是返回 True 否则返回 False|
|<|检查左边值是否小于右边值，如果是返回 True 否则返回 False|
|<=|检查左边值是否小于等于右边值，如果是返回 True 否则返回 False| 
```Go
//relational 关系运算符
func relational() {
	var a float32 = 8
	var b float32 = 3
	var c float32 = 8
	fmt.Printf("a==b吗 %t\n", a == b)
	fmt.Printf("a!=b吗 %t\n", a != b)
	fmt.Printf("a>b吗 %t\n", a > b)
	fmt.Printf("a>=b吗 %t\n", a >= b)
	fmt.Printf("a<c吗 %t\n", a < b)
	fmt.Printf("a<=c吗 %t\n", a <= c)
}
```
逻辑运算符  

|运算符|描述|
| :---: | :--- |
|&|逻辑 AND 运算符。 如果两边的操作数都是 True，则为 True，否则为 False|
| \|\| |逻辑 OR 运算符。 如果两边的操作数有一个 True，则为 True，否则为 False|
|!|逻辑 NOT 运算符。 如果条件为 True，则为 False，否则为 True|
```Go
//logistic 逻辑运算符
func logistic() {
	var a float32 = 8
	var b float32 = 3
	var c float32 = 8
	fmt.Printf("a>b && b>c吗 %t\n", a > b && b > c)
	fmt.Printf("a>b || b>c吗 %t\n", a > b || b > c)
	fmt.Printf("a>b不成立，对吗 %t\n", !(a > b))
	fmt.Printf("b>c不成立，对吗 %t\n", !(b > c))
}
```
位运算符

|运算符|描述|
| :---: | :--- |
|&|参与运算的两数各对应的二进位相与（两位均为1才为1）|
|\||参与运算的两数各对应的二进位相或（两位有一个为1就为1）|
|^|参与运算的两数各对应的二进位相异或，当两对应的二进位相同时为0，不同时为1。作为一元运算符时表示按位取反，，符号位也跟着变|
|<<|左移n位就是乘以2的n次方。a<<b是把a的各二进位全部左移b位，高位丢弃，低位补0。通过左移，符号位可能会变|
|>>|右移n位就是除以2的n次方。a>>b是把a的各二进位全部右移b位，正数高位补0，负数高位补1|
```Go
//bit_op 位运算
func bit_op() {
	fmt.Printf("os arch %s, int size %d\n", runtime.GOARCH, strconv.IntSize) //int是4字节还是8字节，取决于操作系统是32位还是64位
	var a int32 = 260
	fmt.Printf("260     %s\n", util.BinaryFormat(a))
	fmt.Printf("-260    %s\n", util.BinaryFormat(-a)) //负数用补码表示。在对应正数二进制表示的基础上，按拉取反，再末位加1
	fmt.Printf("260&4   %s\n", util.BinaryFormat(a&4))
	fmt.Printf("260|3   %s\n", util.BinaryFormat(a|3))
	fmt.Printf("260^7   %s\n", util.BinaryFormat(a^7))     //^作为二元运算符时表示异或
	fmt.Printf("^-260   %s\n", util.BinaryFormat(^-a))     //^作为一元运算符时表示按位取反，符号位也跟着变
	fmt.Printf("-260>>10 %s\n", util.BinaryFormat(-a>>10)) //正数高位补0，负数高位补1
	fmt.Printf("-260<<3 %s\n", util.BinaryFormat(-a<<3))   //负数左移，可能变成正数
	//go语言没有循环（无符号）左/右移符号   >>>  <<<
}
```
赋值运算符

|运算符|描述|
| :---: | :--- |
|=|简单的赋值运算符，将一个表达式的值赋给一个左值|
|+=|相加后再赋值|
|-=|相减后再赋值|
|*=|相乘后再赋值|
|/=|相除后再赋值|
|%=|求余后再赋值|
|<<=|左移后赋值|
|>>=|右移后赋值|
|&=|按位与后赋值|
|\|=|按位或后赋值|
|^=|按位异或后赋值|
```Go
//assignment 赋值运算
func assignment() {
	var a, b int = 8, 3
	a += b
	fmt.Printf("a+=b %d\n", a)
	a, b = 8, 3
	a -= b
	fmt.Printf("a-=b %d\n", a)
	a, b = 8, 3
	a *= b
	fmt.Printf("a*=b %d\n", a)
	a, b = 8, 3
	a /= b
	fmt.Printf("a/=b %d\n", a)
	a, b = 8, 3
	a %= b
	fmt.Printf("a%%=b %d\n", a) //%在fmt里有特殊含意，所以需要前面再加个%转义一下
	a, b = 8, 3
	a <<= b
	fmt.Printf("a<<=b %d\n", a)
	a, b = 8, 3
	a >>= b
	fmt.Printf("a>>=b %d\n", a)
	a, b = 8, 3
	a &= b
	fmt.Printf("a&=b %d\n", a)
	a, b = 8, 3
	a |= b
	fmt.Printf("a|=b %d\n", a)
	a, b = 8, 3
	a ^= b
	fmt.Printf("a^=b %d\n", a)
}
```
## 变量、常量、字面量
### 变量类型
|类型|go变量类型|fmt输出|
| :---: | :---: | :---: |
|整型|int int8 int16 int32 int64 uint uint8 uint16 uint32 uint64 |%d|
|浮点型|float32 float64 |%f %e %g|复数|complex128 complex64 |%v|
|布尔型|bool|%t|
|指针|uintptr|%p|
|引用|map slice channel|%v|
|字节|byte|%c|
|任意字符|rune|%c|
|字符串|string|%s|
|错误|error|%v|
### 变量声明
&#8195;&#8195;Go语言变量必须先声明再使用，所谓使用指读取或修改。  
标题声明
```Go
var name string 
var age int 
var isOk bool
```
批量声明
```Go
var ( 
	name string 
	age int 
	isOk bool 
)
```
### 变量初始化
&#8195;&#8195;如果声明后未显式初始化，数值型初始化0，字符串初始化为空字符串，布尔型初始化为false，引用类型、函数、指针、接口初始化为nil。
```Go
var a string="china"  //初始化一个变量
var a="china"  //类型推断为string
var a,b int=3,7  //初始化多个变量
var a,b="china",7  //初始化多个变量，每个变量都单独地执行类型推断     
```
&#8195;&#8195;函数内部的变量(非全局变量)可以通过:=声明并初始化。
```Go
a:=3
```
&#8195;&#8195;下划线表示匿名变量。匿名变量不占命名空间，不会分配内存，因此可以重复使用。
```Go
_=2+4
```
### 常量
&#8195;&#8195;常量在定义时必须赋值，且程序运行期间其值不能改变。
```Go
const PI float32=3.14

const(
    PI=3.14
    E=2.71
)

const(
    a=100
    b	//100，跟上一行的值相同
    c	//100，跟上一行的值相同
)
```
iota
```Go
const(
    a=iota	//0
    b		//1
    c		//2
    d		//3
)

const(
    a=iota 	//0
    b		//1
    _		//2
    d		//3
)

const(
    a=iota 	//0
    b=30    
    c=iota 	//2
    d		//3
)

const(
    _=iota		// iota =0
    KB=1<<(10* iota) 	// iota =1
    MB=1<<(10* iota) 	// iota =2
    GB=1<<(10* iota) 	// iota =3
    TB=1<<(10* iota) 	// iota =4
)

const(
    a,b=iota+1, iota+2	//1,2  iota =0
     c,d			//2,3  iota =1
     e,f			//3,4  iota =2
)
```
### 字面量
&#8195;&#8195;字面量--没有出现变量名，直接出现了值。基础类型的字面量相当于是常量。
```Go
fmt.Printf("%t\n", 04 == 4.00) //用到了整型字面量和浮点型字面量
fmt.Printf("%v\n", .4i) //虚数字面量 0.4i
fmt.Printf("%t\n", '\u4f17' == '众') //Unicode和rune字面量
fmt.Printf("Hello\nWorld\n!\n") //字符串字面量
```
## 变量作用域
&#8195;&#8195;对于全局变量，如果以大写字母开头，所有地方都可以访问，跨package访问时需要带上package名称；如果以小写字母开头，则本package内都可以访问。  
&#8195;&#8195;函数内部的局部变量，仅本函数内可以访问。{}可以固定一个作用域。内部声明的变量可以跟外部声明的变量有冲突，以内部的为准--就近原则。
```Go
var (
    A=3	//所有地方都可以访问
    b=4	//本package内可以访问
)

func foo(){
    b:=5  //本函数内可以访问
    {
        b:=6  //本作用域内可以访问
    }
}
```
## 注释与godoc
### 注释的形式
- 单行注释，以//打头。
- 多行注释有2种形式：
    1. 连续多行以//打头，注意多行注释之间不能出现空行。
    2. 在段前使用/\*，段尾使用*/。
- 注释行前加缩进即可写go代码。
- 注释中给定的关键词。NOTE: 引人注意，TODO: 将来需要优化，Deprecated: 变量或函数强烈建议不要再使用。
```Go
//Add 2个整数相加
//返回和。
//
//NOTE: 注释可以有多行，但中间不能出现空行（仅有//不算空行）。
func Add(a, b int) int {
	return a + b
}

/*
Sub 函数使用示例：
  for i:=0;i<3;i++{
	  Sub(i+1, i)
  }
看到了吗？只需要行前缩进，注释里就可以写go代码，是不是很简单。
*/
func Sub(a, b int) int {
	return a - b
}

//TODO: Prod 该函数不能并发调用，需要优化
func Prod(a, b int) int {
	return a * b
}

//Deprecated: Div 不要再调用了
func Div(a, b int) int {
	return a / b
}
```
### 注释的位置
&#8195;&#8195;针对行的注释在行上方或右侧。函数的上方在func xxx()上方。结构体的注释在type xxx struct上方。包注释在package xxx的上方。一个包只需要在一个地方写包注释，通常会专门写一个doc.go，里面只有一行package xxx和关于包的注释。
```Go
// FormatBool, FormatFloat, FormatInt, and FormatUint convert values to strings:
//
//	s := strconv.FormatBool(true)
//	s := strconv.FormatFloat(3.1415, 'E', -1, 64)
//	s := strconv.FormatInt(-42, 16)
//	s := strconv.FormatUint(42, 16)
package fmt
```
### go doc
&#8195;&#8195;go doc是go自带的命令。
```Shell
go doc entrance_class/util
```
上述命令查看entrance_class/util包的注释。
### godoc
&#8195;&#8195;godoc是第三方工具，可以为项目代码导出网页版的注释文档。安装godoc命令如下
```Shell
go get -u golang.org/x/tools/cmd/godoc
go install golang.org/x/tools/cmd/godoc@latest
```
启动http服务：
```Shell
godoc -http=:6060
```
用浏览器访问http://127.0.0.1:6060 ，可以查看go标准库的文档。
