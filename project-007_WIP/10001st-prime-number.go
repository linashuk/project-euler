package main

import "fmt"

// By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

// What is the 10 001st prime number?

func main() {

	slice := []int{}

	for i := 2; i < 20; i++ {
		if isPrime(i) {
			slice = append(slice, i)
		}
	}

	// fmt.Println(slice)

}

func isPrime(num int) bool {

	isPrime := false
	var result bool

	slicePrime := []bool{}

	for divisor := 2; divisor < num; divisor++ {
		if num%divisor != 0 && num%2 != 0 {
			isPrime = true
			slicePrime = append(slicePrime, isPrime)
		} else {
			slicePrime = append(slicePrime, isPrime)
		}
		if divisor == num-1 {
			for i := 0; i < len(slicePrime); i++ {
				if slicePrime[i] == false {
					result = false
					break
				} else {
					result = true
				}
			}
		}

	}
	fmt.Println(num, result)
	return result

}
