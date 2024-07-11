package main

import "fmt"

func Average(arr []int) float64 {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return float64(sum) / float64(len(arr))
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

	fmt.Println(Average(arr))
}
