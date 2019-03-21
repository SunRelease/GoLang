package main

import "fmt"

func main() {

	var str1 string
	//定义字符串
	str1 = "absoulutily"
	//字符串是双引号,还有以\0结束符
	//很类似于c语言
	fmt.Printf("str1=%s\n", str1)
	//占位符为%s

	str2 := "makefoe"
	//便捷设置string类型
	fmt.Printf("str2 type is %T\n", str2)
	//%T为type显示
	var ch byte

	ch = 'a'
	//注意字符是单引号
	fmt.Printf("ch=%c", ch)
	//占位符为%c

}
