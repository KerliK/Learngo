package main

import (
	"fmt"
)

func main() {
	Tigepuks(1, 3)

}

// func main() {
// 	for i := 1; i <= 10; i++ {
// 		if i%2 == 0 {
// 			fmt.Println(i)
// 		} else if i%3 == 0 {
// 			fmt.Println(i)

// 		}
// 	}
// }
func Tigepuks(min, max int) {

	for i := min; i <= max; i++ {
		fmt.Println(i)
	}
}
