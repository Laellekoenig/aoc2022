package main

import (
    "fmt"
    "os"
    "bufio"
)

const (
    ROCK byte = 0
    PAPER byte = 1
    SCISSORS byte = 2
    LOOSE byte = 0
    DRAW byte = 1
    WIN byte = 2
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

        opp := line[0] - 'A'
        outcome := line[2] - 'X'

        // add outcome to score
        if outcome == WIN {
            score += 6
            if opp == ROCK {
                score += 2
            } else if opp == PAPER {
                score += 3
            } else {
                score += 1
            }
        } else if outcome == DRAW {
            score += 3 + 1 + int(opp)
        } else {
            if opp == ROCK {
                score += 3
            } else if opp == PAPER {
                score += 1
            } else {
                score += 2
            }
        }
    }

    fmt.Printf("My score is: %d\n", score)
}
