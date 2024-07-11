package main

import "fmt"

func DeleteDuplicates(arr []int) []int {
	var result []int
	var unique = make(map[int]bool)

	for _, v := range arr {
		_, contains := unique[v]
		if !contains {
			unique[v] = true
			result = append(result, v)
		}
	}

	return result
}

func main() {
	var n int
	fmt.Scan(&n)

	arr := make([]int, 0, n)
	for i := 0; i < n; i++ {
		var elem int
		fmt.Scan(&elem)
		arr = append(arr, elem)
	}

	fmt.Println(DeleteDuplicates(arr))
}
