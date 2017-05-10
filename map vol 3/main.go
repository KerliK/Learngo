package main

import "fmt"

func main() {
	vihm := make(map[int]string)
	vihm[7] = "onneseen"
	vihm[13] = "ebaonn"
	fmt.Println(vihm)
	fmt.Println(vihm[7])
	aike := vihm[13]
	fmt.Println(aike)
	paike, torm := vihm[5]
	fmt.Println(paike, torm)
	udu(7, vihm)
	sugis := pilv(13, vihm)
	fmt.Println(sugis)
	kevad, kevadtwo := leht(8, vihm)
	fmt.Println(kevad, kevadtwo)

	for k, v := range vihm {
		fmt.Println(k, v)

	}
}

func udu(k int, y map[int]string) {
	fmt.Println(k, " - ", y[k])

}
func pilv(g int, p map[int]string) string {
	return p[g]

}
func leht(f int, w map[int]string) (int, string) {
	return f, w[f]
}
