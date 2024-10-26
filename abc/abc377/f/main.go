package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
)

type Interval struct {
    start int64
    end   int64
    size  int64
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    var N int64
    var M int
    fmt.Fscanf(reader, "%d %d\n", &N, &M)

    R_map := make(map[int64]bool)
    C_map := make(map[int64]bool)
    S_set := make(map[int64]struct{})
    D_set := make(map[int64]struct{})

    R_list := make([]int64, 0, M)
    C_list := make([]int64, 0, M)

    for i := 0; i < M; i++ {
        var a, b int64
        fmt.Fscanf(reader, "%d %d\n", &a, &b)
        if !R_map[a] {
            R_map[a] = true
            R_list = append(R_list, a)
        }
        if !C_map[b] {
            C_map[b] = true
            C_list = append(C_list, b)
        }
        S_set[a+b] = struct{}{}
        D_set[a-b] = struct{}{}
    }

    sort.Slice(R_list, func(i, j int) bool { return R_list[i] < R_list[j] })
    sort.Slice(C_list, func(i, j int) bool { return C_list[i] < C_list[j] })

    rows := getIntervals(1, N, R_list)
    cols := getIntervals(1, N, C_list)

    for i := range rows {
        rows[i].size = rows[i].end - rows[i].start + 1
    }
    for i := range cols {
        cols[i].size = cols[i].end - cols[i].start + 1
    }

    total_positions := int64(0)
    for _, row := range rows {
        total_positions += row.size * sumIntervalSizes(cols)
    }

    bad_positions := int64(0)

    for s := range S_set {
        count := countDiagonal(rows, cols, s)
        bad_positions += count
    }

    for d := range D_set {
        count := countAntiDiagonal(rows, cols, d)
        bad_positions += count
    }

    overlap := int64(0)
    for s := range S_set {
        for d := range D_set {
            sum := s
            diff := d
            if (sum+diff)%2 != 0 || (sum-diff)%2 != 0 {
                continue
            }
            i := (sum + diff) / 2
            j := (sum - diff) / 2
            if i < 1 || i > N || j < 1 || j > N {
                continue
            }
            if R_map[i] || C_map[j] {
                continue
            }
            if !inIntervals(rows, i) || !inIntervals(cols, j) {
                continue
            }
            overlap++
        }
    }

    bad_positions -= overlap

    total_good_positions := total_positions - bad_positions
    fmt.Println(total_good_positions)
}

func getIntervals(start, end int64, exclude []int64) []Interval {
    intervals := []Interval{}
    currStart := start
    sort.Slice(exclude, func(i, j int) bool { return exclude[i] < exclude[j] })
    for _, ex := range exclude {
        if ex > currStart {
            intervals = append(intervals, Interval{currStart, ex - 1, 0})
        }
        currStart = ex + 1
    }
    if currStart <= end {
        intervals = append(intervals, Interval{currStart, end, 0})
    }
    return intervals
}

func sumIntervalSizes(intervals []Interval) int64 {
    total := int64(0)
    for _, interval := range intervals {
        total += interval.size
    }
    return total
}

func countDiagonal(rows, cols []Interval, s int64) int64 {
    count := int64(0)
    for _, row := range rows {
        i_min := row.start
        i_max := row.end
        j_min := s - i_max
        j_max := s - i_min

        if j_min > j_max {
            continue
        }

        left := lowerBound(cols, j_min)
        right := upperBound(cols, j_max)

        for i := left; i < right; i++ {
            col := cols[i]
            overlap_start := max(j_min, col.start)
            overlap_end := min(j_max, col.end)
            if overlap_start > overlap_end {
                continue
            }
            count += overlap_end - overlap_start + 1
        }
    }
    return count
}

func countAntiDiagonal(rows, cols []Interval, d int64) int64 {
    count := int64(0)
    for _, row := range rows {
        i_min := row.start
        i_max := row.end
        j_min := i_min - d
        j_max := i_max - d

        if j_min > j_max {
            continue
        }

        left := lowerBound(cols, j_min)
        right := upperBound(cols, j_max)

        for i := left; i < right; i++ {
            col := cols[i]
            overlap_start := max(j_min, col.start)
            overlap_end := min(j_max, col.end)
            if overlap_start > overlap_end {
                continue
            }
            count += overlap_end - overlap_start + 1
        }
    }
    return count
}

func lowerBound(intervals []Interval, x int64) int {
    return sort.Search(len(intervals), func(i int) bool {
        return intervals[i].end >= x
    })
}

func upperBound(intervals []Interval, x int64) int {
    return sort.Search(len(intervals), func(i int) bool {
        return intervals[i].start > x
    })
}

func inIntervals(intervals []Interval, x int64) bool {
    idx := sort.Search(len(intervals), func(i int) bool {
        return intervals[i].end >= x
    })
    if idx < len(intervals) && intervals[idx].start <= x {
        return true
    }
    return false
}

func max(a, b int64) int64 {
    if a > b {
        return a
    }
    return b
}

func min(a, b int64) int64 {
    if a < b {
        return a
    }
    return b
}
