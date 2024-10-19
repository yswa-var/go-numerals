package main

import "fmt"

func largestBelow(multipleOf int, n int) int {
	return multipleOf * ((n - 1) / multipleOf)
}

func sumOfMultiple(n int, limit int) int {
	largest := largestBelow(n, limit)
	count := largest / n
	return n * count * (count + 1) / 2
}

func main() {
	multipleOfThree := sumOfMultiple(3, 1000)
	multipleOfFive := sumOfMultiple(5, 1000)
	multipleOfFifteen := sumOfMultiple(15, 1000)
	result := (multipleOfThree + multipleOfFive) - multipleOfFifteen
	fmt.Println(result)
}
