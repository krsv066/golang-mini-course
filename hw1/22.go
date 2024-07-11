package main

import "fmt"

func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	return arr
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

	fmt.Println(BubbleSort(arr))
}
