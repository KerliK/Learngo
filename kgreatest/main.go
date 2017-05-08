package main

import (
	"fmt"
	"math"
)

func main() {
	b2 := greatest([]int{-1, -2, -3, 4})
	fmt.Println(b2)
}

func greatest(a []int) int {
	m := math.MinInt64
	for _, v := range a {
		if v > m {
			m = v
		}
	}

	return m
}

// ALGVÄÄRTUS
// jooksed üle elementide (for)
// kui element on suurem kui ALGVÄÄRTUS
// siis uus ALGVÄÄRTUS on element
// <=
