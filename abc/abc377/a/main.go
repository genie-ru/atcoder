package main

import (
    "fmt"
    "sort"
    "strings"
)

func main() {
    var s string
    fmt.Scan(&s)

    letters := strings.Split(s, "")
    sort.Strings(letters)

    if strings.Join(letters, "") == "ABC" {
        fmt.Println("Yes")
    } else {
        fmt.Println("No")
    }
}
