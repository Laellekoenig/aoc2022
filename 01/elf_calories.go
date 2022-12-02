package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
)

func main() {
    f, err := os.Open("input.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)
    // read line by line
    scanner.Split(bufio.ScanLines)

    max := 0
    curr := 0

    for scanner.Scan() {
        line := scanner.Text()

        if line == "" {
            if curr > max {
                max = curr
            }
            curr = 0
        } else {
            num, err := strconv.Atoi(line)
            if err != nil {
                fmt.Println(err)
                return
            }
            curr += num
        }
    }

    fmt.Println("The maximum amount of calories is:")
    fmt.Println(max)
}
