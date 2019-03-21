package main

import "fmt"

//导入包main
//一个文件只能有一个主函数main,多了会报错,而且package main一定要用上

func main() { //建立函数ADD,建议用大写,好调用
	fmt.Println("hello.world!")
	//自动换行
	fmt.Print("hello")
	//类似于c语言,换行符要添加才会换行
}
