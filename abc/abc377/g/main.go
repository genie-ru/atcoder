package main

import (
    "bufio"
    "fmt"
    "math"
    "os"
)

type TrieNode struct {
    Children map[byte]*TrieNode
    MinValue int64
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1<<20)
    var N int
    fmt.Fscanln(reader, &N)
    S := make([][]byte, N)
    lengths := make([]int64, N)
    for i := 0; i < N; i++ {
        var str string
        fmt.Fscanln(reader, &str)
        S[i] = []byte(str)
        lengths[i] = int64(len(S[i]))
    }

    root := &TrieNode{
        Children: make(map[byte]*TrieNode),
        MinValue: math.MaxInt64,
    }

    for k := 0; k < N; k++ {
        Lk := lengths[k]
        minTotalCost := Lk

        currentNode := root
        depth := int64(0)

        for i := 0; i < len(S[k]); i++ {
            c := S[k][i]
            if nextNode, exists := currentNode.Children[c]; exists {
                currentNode = nextNode
                depth++
                if currentNode.MinValue != math.MaxInt64 {
                    candidateCost := Lk + currentNode.MinValue
                    if candidateCost < minTotalCost {
                        minTotalCost = candidateCost
                    }
                }
            } else {
                break
            }
        }

        fmt.Println(minTotalCost)

        currentNode = root
        depth = int64(0)
        for i := 0; i < len(S[k]); i++ {
            c := S[k][i]
            if _, exists := currentNode.Children[c]; !exists {
                currentNode.Children[c] = &TrieNode{
                    Children: make(map[byte]*TrieNode),
                    MinValue: math.MaxInt64,
                }
            }
            currentNode = currentNode.Children[c]
            depth++
            computedValue := lengths[k] - 2*depth
            if computedValue < currentNode.MinValue {
                currentNode.MinValue = computedValue
            }
        }
    }
}
