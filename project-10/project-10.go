package main

import "fmt"

// The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

// Find the sum of all the primes below two million.

func main() {

	input := 2000000

	firstBatch := input / 4

	// parentSlice := []int{}

	sum := 0

	for num := input; num > 1; num-- {

		// trying to skip numbers for efficiency

		if (num != 2 && num%2 == 0) || (num != 3 && num%3 == 0) || (num != 5 && num%5 == 0) || (num != 7 && num%7 == 0) {
			continue
		} else if checkPrime(num) {
			sum = sum + num
		}
	}

	fmt.Println(sum)

}

func checkPrime(input int) bool {

	isPrime := true

	for i := 2; i < input; i++ {
		if input%i == 0 {
			isPrime = false
		}
	}

	return isPrime

	// output = 142913828922
	// note: this took forever to run - would like to return to this and see if this can be improved

}
