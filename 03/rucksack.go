package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
)

func main() {
    f, err := os.Open("input.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)
    scanner.Split(bufio.ScanLines)

    var score int

    for scanner.Scan() {
        line := scanner.Text()
        size := len(line)

        compartment_1 := line[:size / 2]
        compartment_2 := line[size / 2:]

        already_checked := make(map[rune]int, 52)

        for _, item := range compartment_1 {
            priority := int((item - 'A' + 27) % 58)
            if strings.ContainsRune(compartment_2, item) {
                if already_checked[item] == 0 {
                    already_checked[item] += 1
                    score += priority
                }
            }
        }
    }

    fmt.Printf("The final score is: %d\n", score)
}
