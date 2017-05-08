package main

import (
	"fmt"
)

func main() {
	a := 2
	b := 3
	fmt.Println(a, " * ", b, " = ", a*b)
	fmt.Println(multiply(1, 2))
	var c int
	multiplyPointer(2, 5, &c)
	fmt.Println(c)
	ohtu := func(x, y int) int {
		return x * y
	}
	fmt.Println(ohtu(4, 5))
}
func multiplyPointer(a, b int, c *int) {
	*c = a * b
}
func multiply(x, y int) int {
	return x * y
}
