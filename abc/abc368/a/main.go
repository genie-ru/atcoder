package main

import "fmt"

func main() {
	var n,k int
	fmt.Scan(&n,&k)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		var t int
		fmt.Scan(&t)
		a[i] = t
	}
	
	for _, b := range a[n-k:] {
		fmt.Print(b)	
		fmt.Print(" ")
	}
	for i, b := range a[:n-k] {
		fmt.Print(b)
		if i < n-k-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
