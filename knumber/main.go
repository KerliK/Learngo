package main

import "fmt"

func main() {
	fmt.Print("Enter big number: ")
	var big int
	fmt.Scan(&big)
	fmt.Print("Enter small number: ")
	var small int
	fmt.Scan(&small)
	fmt.Println(big % small)
}
