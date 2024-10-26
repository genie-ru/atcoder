package main

import (
    "fmt"
)

func main() {
    var n, c int
    fmt.Scan(&n, &c)

    t := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&t[i])
    }

    ans := 0
    pre := -999999
    for i := 0; i < n; i++ {
        if t[i]-pre < c {
            continue
        }
        pre = t[i]
        ans++
    }

    fmt.Println(ans)
}
