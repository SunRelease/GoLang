package main

import "fmt"

func main() {

	a := 1
	fmt.Printf("a is %T\n", a)

	var d, b, c int
	d, b, _ = 1, 2, 4
	fmt.Printf("a=%d\nb=%d\nc=%d\n", d, b, c)

	var (
		e int
		f int
		g float64
	)
	e = 1
	f = 90
	g = 23.5

	fmt.Printf("e=%d\nf=%d\ng=%T\n", e, f, g)

	const (
		i = 1
		j = 2.8
	)

	fmt.Printf("i=%d\n", i)
	fmt.Printf("j=%T", j)

}
