package main

import "fmt"

func main() {
	var x float32
	fmt.Scan(&x)
	for i, f := range []float32{100.0, 38.0, 37.5, 0.0} {
		if x >= f {
			fmt.Println(i)
			break
		}
	}
}