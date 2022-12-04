package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
    "io"
)

func parse_input(inpt string) ([]int, error) {
    split := strings.Split(inpt, ",")
    res := make([]int, 4)

    // elf 1
    e1_str := strings.Split(split[0], "-")
    for i, v := range e1_str {
        num, err := strconv.Atoi(v)
        if err != nil { return nil, err }
        res[i] = num
    }

    // elf 2
    e2_str := strings.Split(split[1], "-")
    for i, v := range e2_str {
        num, err := strconv.Atoi(v)
        if err != nil { return nil, err }
        res[i + 2] = num
    }

    return res, nil
}

func main() {
    f, err := os.Open("input.txt")
    if err != nil {
        fmt.Println(err)
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)
    scanner.Split(bufio.ScanLines)

    // part one
    var count int

    for scanner.Scan() {
        line := scanner.Text()
        elves, err := parse_input(line)
        if err != nil {
            fmt.Println(err)
            return
        }

        if (elves[0] <= elves[2] && elves[1] >= elves[3]) ||
           (elves[0] >= elves[2] && elves[1] <= elves[3]) {
            count++
        }
    }

    fmt.Printf("Number of completely overlapping pairs: %d\n", count)

    // part two
    f.Seek(0, io.SeekStart)
    scanner = bufio.NewScanner(f)
    count = 0

    for scanner.Scan() {
        line := scanner.Text()
        elves, err := parse_input(line)
        if err != nil {
            fmt.Println(err)
            return
        }

        if !(elves[1] < elves[2] || elves[3] < elves[0]) {
            count++
        }
    }

    fmt.Printf("Number of overlapping pairs: %d\n", count)
}
