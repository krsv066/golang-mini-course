package main

import "fmt"

func MultiplicationTable(n int) {
	for i := 1; i <= 10; i++ {
		println(n, "x", i, "=", n*i)
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	MultiplicationTable(n)
}
