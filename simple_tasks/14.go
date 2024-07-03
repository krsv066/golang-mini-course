package main

import "fmt"

func Contains(arr []int, elem int) bool {
	for _, v := range arr {
		if v == elem {
			return true
		}
	}
	return false
}

func main() {
	var n, elem int
	var arr []int

	fmt.Scan(&n)
	arr = make([]int, 0, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&elem)
		arr = append(arr, elem)
	}
	fmt.Scan(&elem)

	fmt.Println(Contains(arr, elem))
}
