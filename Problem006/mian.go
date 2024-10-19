package main

import (
	"fmt"
	"math"
)

func ap(n int) int {
	a := n * (n + 1) / 2
	return a
}

func sumOfSquaredNumbers(n int) int {
	return int((n * (n + 1) * (2*n + 1)) / 6)
}

func main() {
	apSquared := int(math.Pow(float64(ap(100)), 2))
	fmt.Println(apSquared - sumOfSquaredNumbers(100))
}
