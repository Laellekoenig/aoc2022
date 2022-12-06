package main

import (
    "os"
    "bufio"
    "fmt"
)

func isStart(r rune, i int, list *[]rune) int {
    (*list)[i % len(*list)] = r

    different := true

    // check if all are different
    for i, v := range *list {
        if v == 0 { different = false }

        for j := i + 1; j < len(*list); j++ {
            if v == (*list)[j] { different = false }
        }
    }

    if different { return i + 1 }  // adjust for 0 index
    return -1
}

func main() {
    f, err := os.Open("input.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)
    scanner.Split(bufio.ScanLines)
    scanner.Scan()

    input := scanner.Text()
    list := make([]rune, 4)
    var res int = -1

    for i, v := range input {
        res = isStart(v, i, &list)
        if res != -1 { break }
    }

    fmt.Printf("The packet starts at %d.\n", res)

    // part two
    list2 := make([]rune, 14)
    res = -1

    for i, v := range input {
        res = isStart(v, i, &list2)
        if res != -1 { break }
    }

    fmt.Printf("The message startes at %d.\n", res)
}
