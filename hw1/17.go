package main

import "fmt"

func IsPalindrome(str string) bool {
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var str string
	fmt.Scan(&str)
	fmt.Println(IsPalindrome(str))
}
