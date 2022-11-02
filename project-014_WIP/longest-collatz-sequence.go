package main

import "fmt"

func main() {

	length := 0

	for i := 13; i > 12; i-- {
		first := collatzSeq(i)
		second := collatzSeq(i - 2)

		if first > second {
			length = first
		}
	}

	fmt.Println("Max length is:", length)

}

func collatzSeq(input int) int {

	seqLength := 1

	fmt.Println("Input is:", input)

	for i := 1; ; i++ {

		if input%2 == 0 {
			output := input / 2
			seqLength++
			input = output
		} else if input%2 != 0 && input != 1 {
			output := (3 * input) + 1
			seqLength++
			input = output
		} else {
			break
		}

	}
	fmt.Println("Slice length is:", seqLength)
	return seqLength

}
