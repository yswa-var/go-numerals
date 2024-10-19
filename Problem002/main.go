package main

import "fmt"

func main() {
	a := 1
	b := 2
	result := 0
	for b <= 213456 {
		if b%2 == 0 {
			result += b
		}
		c := a + b
		a = b
		b = c
	}
	fmt.Println(result)
}
