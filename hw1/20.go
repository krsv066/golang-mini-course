package main

import "fmt"

func Find(arr []int, num int) int {
	for i, v := range arr {
		if v == num {
			return i
		}
	}
	return -1
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

	fmt.Println(Find(arr, num))
}
