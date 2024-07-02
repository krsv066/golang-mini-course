package main

import "fmt"

func Factorial(num int) int {
	factorial := 1
	for i := num; i > 1; i-- {
		factorial *= i
	}
	return factorial
}

func main() {
	var num int
	fmt.Scan(&num)
	fmt.Println(Factorial(num))
}
