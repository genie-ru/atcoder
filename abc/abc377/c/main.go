package main

import (
    "bufio"
    "fmt"
    "os"
)

type Point struct {
    x, y int64
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    var n, m int64
    fmt.Sscanf(scanner.Text(), "%d %d", &n, &m)

    knights := make([]Point, m)
    for i := int64(0); i < m; i++ {
        scanner.Scan()
        var x, y int64
        fmt.Sscanf(scanner.Text(), "%d %d", &x, &y)
        knights[i] = Point{x, y}
    }

    attacked := make(map[Point]bool)
    occupied := make(map[Point]bool)

    for _, k := range knights {
        occupied[k] = true
    }

    moves := []Point{
        {2, 1}, {1, 2}, {-1, 2}, {-2, 1},
        {-2, -1}, {-1, -2}, {1, -2}, {2, -1},
    }

    for _, k := range knights {
        for _, move := range moves {
            nx := k.x + move.x
            ny := k.y + move.y
            if nx >= 1 && nx <= n && ny >= 1 && ny <= n {
                attacked[Point{nx, ny}] = true
            }
        }
    }

    for pos := range occupied {
        delete(attacked, pos)
    }

    total := n * n
    unsafeCount := int64(len(attacked) + len(occupied))
    safeCount := total - unsafeCount

    fmt.Println(safeCount)
}