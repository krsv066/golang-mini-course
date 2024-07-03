package main

import "fmt"

func ReverseCount(n int) {
	for i := n; i > 0; i-- {
		fmt.Println(i)
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	ReverseCount(n)
}
