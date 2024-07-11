package main

import "fmt"

type Rectangle struct {
	a float64
	b float64
}

func (rectangle Rectangle) Square() float64 {
	return rectangle.a * rectangle.b / 2
}

func main() {
	var a, b float64
	fmt.Scan(&a, &b)

	rectangle := Rectangle{a, b}

	fmt.Println(rectangle.Square())
}
