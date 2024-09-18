package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "strconv"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)

    // N を読み込む
    inputN, _ := reader.ReadString('\n')
    N, _ := strconv.Atoi(strings.TrimSpace(inputN))

    // A を読み込む
    inputA, _ := reader.ReadString('\n')
    inputA = strings.TrimSpace(inputA)
    tokens := strings.Split(inputA, " ")

    A := make([]int, N)
    idx := 0

    // 必要なら追加で読み込む
    for len(tokens) < N {
        line, _ := reader.ReadString('\n')
        line = strings.TrimSpace(line)
        tokens = append(tokens, strings.Split(line, " ")...)
    }

    for i := 0; i < N; i++ {
        A[i], _ = strconv.Atoi(tokens[idx])
        idx++
    }

    operations := 0

    for countPositive(A) > 1 {
        sort.Slice(A, func(i, j int) bool {
            return A[i] > A[j]
        })
        A[0]--
        A[1]--
        operations++
    }

    fmt.Println(operations)
}

func countPositive(A []int) int {
    count := 0
    for _, a := range A {
        if a > 0 {
            count++
        }
    }
    return count
}
