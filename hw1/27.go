package main

import "fmt"

func Merge(arr1, arr2 []int) []int {
	var result []int
	i, j := 0, 0

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			result = append(result, arr1[i])
			i++
		} else {
			result = append(result, arr2[j])
			j++
		}
	}

	for i < len(arr1) {
		result = append(result, arr1[i])
		i++
	}
	for j < len(arr2) {
		result = append(result, arr2[j])
		j++
	}

	return result
}

func main() {
	var n1, n2 int

	fmt.Scan(&n1)
	arr1 := make([]int, 0, n1)
	for i := 0; i < n1; i++ {
		var elem int
		fmt.Scan(&elem)
		arr1 = append(arr1, elem)
	}

	fmt.Scan(&n2)
	arr2 := make([]int, 0, n2)
	for i := 0; i < n2; i++ {
		var elem int
		fmt.Scan(&elem)
		arr2 = append(arr2, elem)
	}

	fmt.Println(Merge(arr1, arr2))
}
