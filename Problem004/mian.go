package main

import "fmt"

func isPalindromic(n int) bool {
	var temp int = n
	var reverse int = 0
	for temp > 0 {
		dig := temp % 10
		reverse *= 10
		reverse += dig
		temp /= 10
	}
	return reverse == n
}

func main() {
	largestPalindrome := 0

	for i := 999; i >= 100; i-- {
		for j := 999; j >= 100; j-- {
			product := i * j
			if isPalindromic(product) && product > largestPalindrome {
				largestPalindrome = product
			}
		}
	}
	fmt.Println(largestPalindrome)
}
