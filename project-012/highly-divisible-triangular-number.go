package main

import "fmt"

// The sequence of triangle numbers is generated by adding the natural numbers. So the 7th triangle number would be 1 + 2 + 3 + 4 + 5 + 6 + 7 = 28. The first ten terms would be:

// 1, 3, 6, 10, 15, 21, 28, 36, 45, 55, ...

// We can see that 28 is the first triangle number to have over five divisors.

// What is the value of the first triangle number to have over five hundred divisors?

func main() {

	// slice := []int{}

	var triNum, numDivisors int

	for i := 1; ; i++ {

		triNum, numDivisors = triangleNumberAndCount(i)

		if numDivisors <= 500 {
			continue
		} else {
			fmt.Println(triNum, "has", numDivisors, "divisors")
			break
		}

	}
}

func triangleNumberAndCount(n int) (int, int) {

	// finds the n'th triangle number

	triangleNumber := (n * (n + 1)) / 2

	// declare a count of divisors
	var countDivisors int

	// iterates to count divisors
	for i := 1; i <= triangleNumber; i++ {
		if triangleNumber%i == 0 {
			countDivisors++
		}
	}

	// returns the n'th triangle number and the count for the main function

	return triangleNumber, countDivisors
}

//output = 76576500 has 576 divisors
