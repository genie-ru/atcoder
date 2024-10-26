package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1<<20)
    line, _ := reader.ReadString('\n')
    line = strings.TrimSpace(line)
    parts := strings.Split(line, " ")
    N, _ := strconv.Atoi(parts[0])
    K, _ := strconv.ParseInt(parts[1], 10, 64)

    line, _ = reader.ReadString('\n')
    line = strings.TrimSpace(line)
    parts = strings.Split(line, " ")

    P := make([]int, N)
    for i := 0; i < N; i++ {
        x, _ := strconv.Atoi(parts[i])
        P[i] = x - 1
    }

    visited := make([]bool, N)
    result := make([]int, N)

    for i := 0; i < N; i++ {
        if !visited[i] {
            cycle := []int{}
            curr := i
            for !visited[curr] {
                visited[curr] = true
                cycle = append(cycle, curr)
                curr = P[curr]
            }

            L := len(cycle)
            m := modPow(2, K, int64(L))

            for j := 0; j < L; j++ {
                newPos := (j + int(m)) % L
                result[cycle[j]] = cycle[newPos]
            }
        }
    }

    output := make([]string, N)
    for i := 0; i < N; i++ {
        output[i] = strconv.Itoa(result[i]+1)
    }
    fmt.Println(strings.Join(output, " "))
}

func modPow(a, b int64, mod int64) int64 {
    result := int64(1)
    a = a % mod
    for b > 0 {
        if b%2 == 1 {
            result = (result * a) % mod
        }
        a = (a * a) % mod
        b /= 2
    }
    return result
}
