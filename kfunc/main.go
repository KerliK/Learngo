package main

import (
	"fmt"
)

func main() {
	a := hei(10)
	fmt.Println(a)

}

func hei(max int) []int {
	var jo []int
	for i := 0; i <= max; i++ {

		jo = append(jo, i)

	}

	return jo
	return []int{}
}
