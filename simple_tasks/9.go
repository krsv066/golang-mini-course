package main

import "fmt"

func Sum(arr []int) int {
	var sum int
	for _, elem := range arr {
		sum += elem
	}
	return sum
}

func main() {
	var n int
	fmt.Scan(&n)

	var arr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Println(Sum(arr))
}
