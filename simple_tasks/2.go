package main

import "fmt"

func add(lhs int, rhs int) int {
	return lhs + rhs
}

func main() {
	s := add(-2, 5)
	fmt.Println(s)
}
