package main

import (
	"fmt"
)

func main() {

	kerli := func(n int) (int, bool) {
		return n / 2, n%2 == 0
	}
	fmt.Println(kerli(1))
	fmt.Println(kerli(2))
	fmt.Println(kerli(10))
}

// func main() {
// 	marek := half
// 	fmt.Println(marek(1))
// 	fmt.Println(marek(2))
// 	fmt.Println(marek(10))
// }
// func half(i int) (int, bool) {
// 	return i / 2, i%2 == 0

// }
