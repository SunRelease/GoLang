package main

import "fmt"

func main() {
	const (
		a = iota
		//iota的枚举  从0开始
		b = iota
		//依次递增+1
		c = iota

		d = iota
	)
	fmt.Printf("a=%d\nb=%d\nc=%d\nd=%d\n", a, b, c, d)

	const (
		e = iota

		f

		g1, g2, g3 = iota, iota, iota
		//同行iota值不变,都是同一个值
	)
	fmt.Printf("e=%d\nf=%d\ng1=%d\ng2=%d\ng3=%d\n", e, f, g1, g2, g3)

}
