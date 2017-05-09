package main

import "fmt"

func main() {
	a := uks("lollpea")
	fmt.Println(a)
}

func uks(uk string) map[string]int {
	kuu := make(map[string]int)
	for _, v := range uk {
		kuu[string(v)]++
	}

	return kuu
}
