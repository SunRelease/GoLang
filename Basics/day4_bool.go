package main

import "fmt"

func main() {

	var a bool
	//定义bool型

	fmt.Println("a1=", a)
	//未初始化前  默认a 是False>>0
	a = false

	fmt.Println("a=", a)

	b := false
	//便捷型定义b的值为bool型
	fmt.Println("b=", b)
}
