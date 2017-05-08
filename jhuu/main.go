package main

import (
	"fmt"
)

func main() {

	c := vesi([]int{1, 25, 29, 24, 21, 11, 100}...)
	fmt.Println(c)
	fmt.Println(vesi(1, 2, 5, 78, 33))
	z := kuum(3)
	fmt.Println(z)
}
func vesi(max ...int) []int {
	var tuli []int
	for _, v := range max {
		if v%3 == 0 {
			tuli = append(tuli, v)
		}
	}
	return tuli

}
func kuum(boo int) int {

	var total int
	for i := 0; i <= boo; i++ {
		total += i
	}
	return total

}
func kulm() int {
	return 6
}
