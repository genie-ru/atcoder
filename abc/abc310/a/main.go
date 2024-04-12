package main

import (
	"fmt"
	"math"
)

func main() {
	var n, p, q int
	fmt.Scan(&n, &p, &q)

	dMin := math.MaxInt32
	var d int
	for i := 1; i <= n; i++ {
		fmt.Scan(&d)
		dMin = int(math.Min(float64(d), float64(dMin)))
	}

	fmt.Println(min(p, q+dMin))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
