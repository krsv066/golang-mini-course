package main

import "fmt"

func MaxOfThree(a, b, c int) int {
	if a > b {
		if a > c {
			return a
		}
	} else {
		if b > c {
			return b
		}
	}
	return c
}

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	fmt.Println(MaxOfThree(a, b, c))
}
