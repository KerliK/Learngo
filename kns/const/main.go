package main

import (
	"fmt"
)

const kerli = 22

func main() {
	fmt.Println(kerli)
	var a int
	a = 10
	fmt.Println(a)
	b := 8
	kuu(&b)
	fmt.Println(b)

}
func kuu(x *int) {
	*x = 5
}
