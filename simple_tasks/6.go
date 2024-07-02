package main

import (
	"fmt"
	"slices"
)

func IsVowel(letter rune) bool {
	vowels := []rune{'a', 'e', 'i', 'o', 'u', 'y', 'а', 'е', 'ё', 'и', 'о', 'у', 'ы', 'э', 'ю', 'я'}
	return slices.Contains(vowels, letter)
}

func main() {
	var letter rune
	fmt.Scanf("%c", &letter)
	fmt.Println(IsVowel(letter))
}
