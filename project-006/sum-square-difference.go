package main

import (
	"fmt"
	"math"
)

func main() {

	var sumSquare, sumNum float64

	for i := 1; i <= 100; i++ {
		sumSquare = float64(sumSquare) + math.Pow(float64(i), float64(2))
	}

	for i := 1; i <= 100; i++ {
		sumNum = sumNum + float64(i)
	}

	sumNum = math.Pow(float64(sumNum), float64(2))

	diffNum := sumNum - sumSquare

	s := fmt.Sprintf("%f", diffNum)
	fmt.Println(s)

	// output = 25164150.000000

}
