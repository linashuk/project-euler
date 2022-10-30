package main

import "fmt"

func main() {

	// The prime factors of 13195 are 5, 7, 13 and 29.
	// What is the largest prime factor of the number 600851475143 ?

	number := 600851475143

	slice := []int{}

	for i := 1; i <= number; i++ {
		if number%i == 0 {
			slice = append(slice, i)
			number = number / i
		}
	}

	fmt.Println(slice[len(slice)-1])

	//output = 6857

}
