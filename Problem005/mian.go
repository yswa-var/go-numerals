package main

import "fmt"

func smallestMultiple() int {
	smallest := 20
	for {
		isDivisible := true
		for i := 20; i >= 1; i-- {
			if smallest%i != 0 {
				isDivisible = false
				break
			}
		}
		if isDivisible {
			return smallest
		}
		smallest++
	}
}

func main() {
	fmt.Println(smallestMultiple())
}
