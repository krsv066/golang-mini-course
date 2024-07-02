package main

import "fmt"

func Primes(num int) {
	var primes = make([]bool, num+1)
	for i := 0; i < num+1; i++ {
		primes[i] = true
	}

	for prime := 2; prime <= num; prime++ {
		if primes[prime] {
			fmt.Println(prime)
			for j := 2 * prime; j <= num; j += prime {
				primes[j] = false
			}
		}
	}
}

func main() {
	var number int
	fmt.Scan(&number)
	Primes(number)
}
