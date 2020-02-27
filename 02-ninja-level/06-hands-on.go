package main

import "fmt"

const (
	x = iota + 2019
	y
	z
)

func main() {

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
