package main

import "fmt"

func Add(lhs, rhs int) int {
	return lhs + rhs
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(Add(a, b))
}
