package main

// Work out the first ten digits of the sum of the following one-hundred 50-digit numbers.

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	numbersSlice := []float64{}

	f, _ := os.Open("./numbers.txt")

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		num, _ := strconv.ParseFloat(line, 64)
		// fmt.Println(num)
		numbersSlice = append(numbersSlice, num)
	}

	var sum float64

	// fmt.Println(len(numbersSlice))

	for i := 0; i < len(numbersSlice); i++ {
		sum = sum + numbersSlice[i]
	}

	sumToString := fmt.Sprintf("%f", sum)

	fmt.Println(sumToString)

	// output = 5537376230390877287140145935443224959721771787878400.000000
	// answer = 55373762303

}
