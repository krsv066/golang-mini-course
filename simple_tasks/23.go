package main

import "fmt"

func Fibonacci(n int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		fmt.Println(a)
		a, b = b, a+b
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	Fibonacci(n)
}
