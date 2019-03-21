package main

import "fmt"

func main() {
	//%d 数字的使用
	var a int
	a = 78
	//%f 浮点型的使用
	var b float64
	b = 2.4
	//%s 字符串的使用
	var str string
	str = "abcdefg"
	//%c 字符的使用
	var ch byte
	ch = 'a'
	//%v 智能识别类型并输出,对%c字符可能不友好,对float比较友好
	var bl bool
	bl = false
	//%t  bool型的输出
	fmt.Printf("ch is %v\nstr is %v\n b is %v\na is %v\nb1=%t", ch, str, b, a, bl)
}
