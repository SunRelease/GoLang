package main

import "fmt"

func main() {

	var cm1 complex64
	//定义cm1为复数
	cm1 = 2 + 3.2i
	fmt.Printf("cm1=\n", cm1)

	cm2 := 6 + 9i
	//系统默认是128位
	fmt.Printf("cm2 type is %T\nreal is%v\nimag is %v\n", cm2, real(cm2), imag(cm2))
	//内建函数  real是表示实部,imag是表示虚部
}
