package main

import "fmt"

func MinAndMax(arr []int) (int, int) {
	mi, ma := arr[0], arr[0]

	for _, v := range arr {
		if v < mi {
			mi = v
		}
		if v > ma {
			ma = v
		}
	}

	return mi, ma
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

	fmt.Println(MinAndMax(arr))
}
