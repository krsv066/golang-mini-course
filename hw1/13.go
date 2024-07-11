package main

import "fmt"

func Len(str string) int {
	l := 0
	for i, _ := range str {
		l = i
	}
	return l + 1
}

func main() {
	var str string
	fmt.Scan(&str)
	fmt.Println(Len(str))
}
