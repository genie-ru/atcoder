package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "strconv"
    "strings"
)

type Interval struct {
    L int
    R int
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1<<20)
    N, M := readTwoInts(reader)
    intervals := make([]Interval, N)
    for i := 0; i < N; i++ {
        L, R := readTwoInts(reader)
        intervals[i] = Interval{L, R}
    }

    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i].L < intervals[j].L
    })

    MPlus1 := M + 1
    minRiFromLi := make([]int, M+2)
    for i := 0; i <= M+1; i++ {
        minRiFromLi[i] = MPlus1
    }

    idx := len(intervals) - 1
    minRi := MPlus1
    for l := M; l >= 1; l-- {
        for idx >= 0 && intervals[idx].L >= l {
            if intervals[idx].R < minRi {
                minRi = intervals[idx].R
            }
            idx--
        }
        minRiFromLi[l] = minRi
    }

    totalIntervals := M * (M + 1) / 2
    totalInvalid := 0

    for l := 1; l <= M; l++ {
        if minRiFromLi[l] <= M {
            invalidR := M - minRiFromLi[l] + 1
            totalInvalid += invalidR
        }
    }

    answer := totalIntervals - totalInvalid
    fmt.Println(answer)
}

func readTwoInts(reader *bufio.Reader) (int, int) {
    for {
        line, err := reader.ReadString('\n')
        if err != nil {
            break
        }
        fields := strings.Fields(line)
        if len(fields) >= 2 {
            a, _ := strconv.Atoi(fields[0])
            b, _ := strconv.Atoi(fields[1])
            return a, b
        }
    }
    return 0, 0
}
