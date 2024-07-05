package main

import (
	"fmt"
	"strings"
)

func AreAnagrams(str1, str2 string) bool {
	str1 = strings.ToLower(str1)
	str2 = strings.ToLower(str2)

	if len(str1) != len(str2) {
		return false
	}

	dict := make(map[rune]int)
	for _, r := range str1 {
		dict[r]++
	}

	for _, r := range str2 {
		dict[r]--
		if dict[r] < 0 {
			return false
		}
	}

	return true
}

func main() {
	var str1, str2 string
	fmt.Scan(&str1, &str2)
	fmt.Println(AreAnagrams(str1, str2))
}
