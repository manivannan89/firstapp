package main

import "fmt"

func main() {

	a := 100
	b := 100

	if a == b {
		fmt.Println("a is equal to b")
	}

	if a <= b {
		fmt.Println("a is less than or equal to b")
	}

	if a >= b {
		fmt.Println("a is greater than or equal to b")
	}

	if a != b {
		fmt.Println("a is not equal to b")
	}

	if a < b {
		fmt.Println("a is less than b")
	}

	if a > b {
		fmt.Println("a is greater than b")
	}

}
