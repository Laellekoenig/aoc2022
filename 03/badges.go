package main

import (
    "fmt"
    "os"
    "bufio"
)

func update_badges(rucksack string, badges *map[rune]int) {
    tmp := make(map[rune]int, 52)
    for _, item := range rucksack {
        if tmp[item] == 0 {
            tmp[item]++
            (*badges)[item]++
        }
    }
}

func get_priority(r rune) int {
    return int((r - 'A' + 27) % 58)
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

    var score int

    for scanner.Scan() {
        badges := make(map[rune]int, 52)

        for i := 0; i < 3; i++ {
            elf := scanner.Text()
            update_badges(elf, &badges)
            if i != 2 {
                scanner.Scan()
            }
        }

        for k, v := range badges {
            if v == 3 {
                score += get_priority(k)
                break
            }
        }
    }

    fmt.Printf("The final score is: %d\n", score)
}
