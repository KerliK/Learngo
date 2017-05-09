package main

import "fmt"

func main() {
	a := uks("lollpea")
	fmt.Println(a)
	h, k := kolm(a)
	fmt.Println(h, k)
}

func uks(uk string) map[string]int {
	kuu := make(map[string]int)
	for _, v := range uk {
		kuu[string(v)]++
	}

	return kuu
}
func kolm(neli map[string]int) ([]string, []int) {
	paike := []string{}
	maa := []int{}
	for k, v := range neli {
		paike = append(paike, k)
		maa = append(maa, v)
	}

	return paike, maa
}
