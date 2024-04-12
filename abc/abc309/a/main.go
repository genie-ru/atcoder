package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	ans := "No"
	if a%3 != 0 && a+1 == b {
		ans = "Yes"
	}

	fmt.Println(ans)
}
