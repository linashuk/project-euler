package main

import "fmt"

// A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

// Find the largest palindrome made from the product of two 3-digit numbers.

func main() {

	// max 999*999 = 998001
	// min 100*100 = 10000

	start := 998001
	sliceInput := []int{}
	sliceA := []int{}
	sliceB := []int{}

	for i := start; i > 10000; i-- {
		a, b := calculator(i)
		if a > 100 && b > 100 {
			sliceInput = append(sliceInput, i)
			sliceA = append(sliceA, a)
			sliceB = append(sliceB, b)
		}
	}

	for i := 1; i < len(sliceInput); i++ {
		if isPalindrome(sliceInput[i]) {
			fmt.Println(sliceA[i], "*", sliceB[i], "=", sliceInput[i])
			break
		}
	}

	// output = 913 * 993 = 906609

}

func calculator(input int) (int, int) {

	var a, b int

	for i := 100; i < 1000; i++ {
		if input%i != 0 {
			continue
		} else if input != 0 {
			n := input / i
			if n > 100 && n < 1000 {
				a = i
				b = n
				break
			}
		} else {
			continue
		}
	}

	return a, b

}

func isPalindrome(n int) bool {

	isPalindrome := false

	if n > 10000 && n < 100000 {
		one := n % 10
		two := (n / 10) % 10
		// three := (n / 100) % 10
		four := (n / 1000) % 10
		five := (n / 10000) % 10
		if one == five && two == four {
			// fmt.Println(one, two, three, four, five)
			isPalindrome = true
		}
	} else if n >= 100000 {
		one := n % 10
		two := (n / 10) % 10
		three := (n / 100) % 10
		four := (n / 1000) % 10
		five := (n / 10000) % 10
		six := (n / 100000) % 10
		if one == six && two == five && three == four {
			isPalindrome = true
		}
	}
	return isPalindrome
}
