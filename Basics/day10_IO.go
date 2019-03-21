package main

import "fmt"

func main() {
	var a int
	//需要先定义一个东西,否则无法判断信息
	fmt.Scanf("input is \n", &a)
	//推荐上面,因为有信息显示好些,但是下面比较简便
	fmt.Scan(&a)

	fmt.Printf("input is %v\n", a)

}
