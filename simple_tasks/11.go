package main

import "fmt"

func CelsiusToFahrenheit(temp float64) float64 {
	return temp*9/5 + 32
}

func main() {
	var temp float64
	fmt.Scan(&temp)
	fmt.Println(CelsiusToFahrenheit(temp))
}
