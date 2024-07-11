package main

import "fmt"

func Intersection(arr1, arr2 []int) []int {
	var contains1 = make(map[int]bool)
	var contains2 = make(map[int]bool)
	var res []int

	for _, elem := range arr1 {
		contains1[elem] = true
	}
	for _, elem := range arr2 {
		contains2[elem] = true
	}

	for k, _ := range contains1 {
		if contains2[k] {
			res = append(res, k)
		}
	}

	return res
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

	fmt.Println(Intersection(arr1, arr2))
}
