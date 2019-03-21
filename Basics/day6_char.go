package main

import "fmt"

func main() {

	var (
		ch1 byte
		ch2 byte
		ch3 byte
	)
	//定义ch 为ASCII码
	ch1 = 'a'

	ch2 = 98

	ch3 = '\n'
	//ASCII码为小写a
	//ascii码为数字
	//ASCII码为反义字符

	fmt.Printf("ch1=%c\nch2=%c\nch3=%c\n", ch1, ch2, ch3)

	fmt.Printf("e=%c", '\n')
	//'\n'表示是反义换行符\n,是替换%c,也是ASCII码的一部分
	fmt.Printf("hello,world!")
}
