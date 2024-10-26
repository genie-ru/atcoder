package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    
    var rowMask, colMask uint8
    
    board := make([]string, 8)
    
    for i := 0; i < 8; i++ {
        line, _ := reader.ReadString('\n')
        board[i] = line[:8]
        
        for j := 0; j < 8; j++ {
            if line[j] == '#' {
                rowMask |= 1 << i
                colMask |= 1 << j
            }
        }
    }
    
    safeSpaces := 0
    emptyRows := ^rowMask & 0xFF
    emptyCols := ^colMask & 0xFF
    
    for i := 0; i < 8; i++ {
        if emptyRows&(1<<i) != 0 {
            for j := 0; j < 8; j++ {
                if emptyCols&(1<<j) != 0 && board[i][j] == '.' {
                    safeSpaces++
                }
            }
        }
    }
    
    fmt.Println(safeSpaces)
}