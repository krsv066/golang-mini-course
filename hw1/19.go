package main

import "fmt"

func DeleteByInd(arr []int, ind int) []int {
	return append(arr[:ind], arr[ind+1:]...)
}

func main() {
	var n, ind int
	fmt.Scan(&n)

	arr := make([]int, 0, n)
	for i := 0; i < n; i++ {
		var elem int
		fmt.Scan(&elem)
		arr = append(arr, elem)
	}
	fmt.Scan(&ind)

	fmt.Println(DeleteByInd(arr, ind))
}
