package main

import "fmt"

// A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

// Find the largest palindrome made from the product of two 3-digit numbers.

func main() {

	// max 999*999 = 998001
	// min 100*100 = 10000

	var s string

	for num := 998001; num > 10000; num-- {
		s = string(num)
	}

	numberOne := 100
	numberTwo := 999

	for numberOne <= 999 && numberTwo >= 100 {

		product := numberOne * numberTwo

		numberOne++
		numberTwo--

		fmt.Println(product)

	}

}
