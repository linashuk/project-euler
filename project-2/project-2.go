package main

import "fmt"

func main() {

	fibonacciSlice := []int{}

	// declare starting integers

	a, b := 0, 1
	var c int

	for a+b < 4000000 {

		c = a + b

		if c%2 == 0 {

			// append to slice
			fibonacciSlice = append(fibonacciSlice, c)

		}

		// reassign a and b integers
		a = b
		b = c
	}

	var sumFibonacci int

	// interate through the slice
	for i := 0; i < len(fibonacciSlice); i++ {
		sumFibonacci = sumFibonacci + fibonacciSlice[i]
	}

	fmt.Println(sumFibonacci)

	// output = 4613732

}
