package main

import (
	"fmt"
	"os"
)

func main() {
	var n int
	var s []byte
	fmt.Scan(&n, &s)

	m := make(map[byte]bool, 3)
	for i, v := range s {
		m[v] = true
		if len(m) == 3 {
			fmt.Println(i + 1)
			os.Exit(0)
		}
	}
}
