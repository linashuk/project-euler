package main

import "fmt"

func main() {

	fibonacciSlice := []int{2}
	fmt.Println(fibonacciSlice)

	firstInt := 1
	secondInt := 2
	var thirdInt int

	for firstInt+secondInt < 4000000 {
		thirdInt = firstInt + secondInt
		if thirdInt%2 == 0 {
			fibonacciSlice = append(fibonacciSlice, thirdInt)
		}
		firstInt = secondInt
		secondInt = thirdInt
	}

	fmt.Println(fibonacciSlice)

}
