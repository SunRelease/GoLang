package main

import (
	"fmt"
	//"time"
)

func main() {
	//for var1,var2 :=range string_name{
	//实现遍历

	var str string
	str = "abcdefg"
	for i, data := range str {
		fmt.Printf("下标号%d 字母切片%c\n", i, data)
	}

	//for循环 变量;判断条件;增量 {}
	sum := 0

	for j := 1; j <= 100; j++ {
		sum = sum + j
		go fmt.Println("next is ", j)

	}
	fmt.Println("sum =", sum)
}
