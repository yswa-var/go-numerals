package main

import "fmt"

var answer = 1

func main() {
	n := 600851475143
	for n%2 == 0 {
		answer = 2
		n /= 2
	}
	for i := 3; i*i <= n; i += 2 {
		for n%i == 0 {
			if i > answer {
				answer = i
			}
			n /= i
		}
	}
	if n > 2 {
		answer = n
	}
	fmt.Println(answer)
	fmt.Println(n / 1471)
}
