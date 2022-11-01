package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// The four adjacent digits in the 1000-digit number that have the greatest product are 9 × 9 × 8 × 9 = 5832.

// Find the thirteen adjacent digits in the 1000-digit number that have the greatest product. What is the value of this product?

func main() {

	numbersSlice := []int{}

	f, _ := os.Open("./numbers.txt")

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanRunes)

	for scanner.Scan() {
		number := scanner.Text()
		num, _ := strconv.Atoi(number)
		numbersSlice = append(numbersSlice, num)
	}

	// fmt.Print(numbersSlice)
	// fmt.Println(numbersSlice[0])

	sumSlice := []int{}
	width := 13

	for end := 12; end < len(numbersSlice); end++ {

		sum := 1

		for iterate := 0; iterate < width; iterate++ {
			// fmt.Println(numbersSlice[i-n])
			sum = sum * numbersSlice[end-iterate]
		}

		sumSlice = append(sumSlice, sum)
	}

	//  fmt.Println(sumSlice)

	max := 0

	for find := 1; find < len(sumSlice); find++ {
		if sumSlice[find] > max {
			max = sumSlice[find]
		}

	}

	fmt.Println(max)

	// output = 23514624000

}
