package main

import (
	"fmt"
)

func main() {
	marek(10, 10)
}
func marek(uks, kaks int) {
	for i := 0; i <= uks; i++ {
		fmt.Println("i", i)
		for j := 0; j <= kaks; j++ {
			fmt.Println("j", j)
			kook := kerli(i, j)
			if kook%3 == 0 {
				continue
			}
			fmt.Println(kook)

		}

	}
}
func kerli(kolm, neli int) int {
	return kolm * neli
}

// func main() {
// 	var uus int
// 	uus = kaksarvu(4, 5)
// 	fmt.Println(uus)

// }
// func kaksarvu(uks, kaks int) int {

// 	return uks * kaks
// }
