package main

import (
    "os"
    "fmt"
    "bufio"
)

const (
    ROCK byte = 0
    PAPER byte = 1
    SCISSORS byte = 2
)

func main() {
    f, err := os.Open("input.txt")
    if err != nil {
        fmt.Println(err)
        return;
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)
    scanner.Split(bufio.ScanLines)

    var score int

    for scanner.Scan() {
        line := scanner.Text()

        opponent := line[0] - 'A'
        me := line[2] - 'X'

        // what I play
        score += int(me) + 1

        // determine outcome
        if opponent == me {
            score += 3
        } else if (me == ROCK && opponent == SCISSORS) ||
                  (me == PAPER && opponent == ROCK) ||
                  (me == SCISSORS && opponent == PAPER) {
                      score += 6
                  }
    }

    fmt.Printf("My score is: %d\n", score)
}
