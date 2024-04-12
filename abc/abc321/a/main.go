package main

import (
	"fmt"
)

func main() {
	var N string
	fmt.Scan(&N)
	prevDigit := 'A'

	for _, digit := range N {
		if digit >= prevDigit {
			fmt.Println("No")
			return
		}
		prevDigit = digit
	}
	fmt.Println("Yes")
}
