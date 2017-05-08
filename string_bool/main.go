package main

import (
	"fmt"
)

func main() {

	a := func(h string) (bool, int) {
		return h == "yes", len(h)
	}
	marek()
	fmt.Println(answer("yes"))
	fmt.Println(a("yes"))
	fmt.Println(yesno(10))
}
func yesno(j int) bool {

	return j != 8

}

func marek() {
	fmt.Println("yes")
}

func answer(h string) (bool, int) {
	return h == "yes", len(h)
}
