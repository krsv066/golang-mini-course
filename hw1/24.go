package main

import "fmt"

func Count(arr []int, num int) int {
	count := 0
	for _, elem := range arr {
		if elem == num {
			count++
		}
	}
	return count
}

func main() {
	var n, num int
	fmt.Scan(&n)

	arr := make([]int, 0, n)
	for i := 0; i < n; i++ {
		var elem int
		fmt.Scan(&elem)
		arr = append(arr, elem)
	}
	fmt.Scan(&num)

	fmt.Println(Count(arr, num))
}
