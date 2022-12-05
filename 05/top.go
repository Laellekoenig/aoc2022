package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	BUFF             int = 40 // number of empty cells above highest crate
	CRATE_MOVER_9000 int = 9000
	CRATE_MOVER_9001 int = 9001
)

func input2Array(input string) [][]byte {
	lines := strings.Split(input, "\n")

	// get number of cols
	nums := lines[len(lines)-2]
	num_cols := int(nums[len(nums)-1] - 48)
	num_rows := len(lines) - 2 // ignore row containing numbers of cols and blank row

	// construct final 2D array
	result := make([][]byte, num_rows+BUFF)
	for i := range result {
		result[i] = make([]byte, num_cols)
	}

	// fill it with values
	for i, v := range lines[:num_rows] {
		for j := 0; j < num_cols; j++ {
			r := v[1+j*4]
			if r != ' ' {
				result[i+BUFF][j] = r
			}
		}
	}

	return result
}

func makeMove(move string, stack *[][]byte, moverType int) error {
	split := strings.Split(move, " ")

	// parse instruction
	num, err := strconv.Atoi(split[1])
	if err != nil {
		return err
	}
	from, err := strconv.Atoi(split[3])
	if err != nil {
		return err
	}
	to, err := strconv.Atoi(split[5])
	if err != nil {
		return err
	}

	// execute
	if moverType == CRATE_MOVER_9000 {
		for i := 0; i < num; i++ {
			err := moveFIFO(from, to, stack)
			if err != nil {
				return err
			}
		}
	} else if moverType == CRATE_MOVER_9001 {
		err := moveLIFO(num, from, to, stack)
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("Unknown mover type %d\n", moverType)
	}

	return nil
}

func moveLIFO(num, from, to int, stack *[][]byte) error {
	// adjust for 0 index
	from -= 1
	to -= 1

	var from_i int = -1
	for i := range *stack {
		curr := (*stack)[i][from]
		if curr != 0 {
			from_i = i
			break
		}
	}

	if from_i < 0 {
		return fmt.Errorf("Invalid move, empty col.\n")
	}
	if from_i+num-1 == len(*stack) {
		return fmt.Errorf("Not enough crates in col.\n")
	}

	var to_i int = len(*stack) - 1
	for i := range *stack {
		curr := (*stack)[i][to]
		if curr != 0 {
			to_i = i - 1
			break
		}
	}

	if to_i < 0 {
		return fmt.Errorf("Invalid move, empty col.\n")
	}
	if to_i-num+1 < 0 {
		return fmt.Errorf("Not enough space in col.\n")
	}

	for i := 0; i < num; i++ {
		(*stack)[to_i-i][to] = (*stack)[from_i+num-1-i][from]
		(*stack)[from_i+num-1-i][from] = 0
	}

	return nil
}

func moveFIFO(from, to int, stack *[][]byte) error {
	// adjust for 0 index
	from -= 1
	to -= 1

	var crate byte
	for i := range *stack {
		crate = (*stack)[i][from]
		if crate != 0 {
			(*stack)[i][from] = 0
			break
		}

		if i == len(*stack)-1 {
			return fmt.Errorf("Cannot move crate from empty stack.\n")
		}
	}

	for i := range *stack {
		curr := (*stack)[i][to]
		if curr != 0 {
			if i-1 < 0 {
				return fmt.Errorf("Index %d out of bounds for length %d\n", i-1, len(*stack))
			}
			(*stack)[i-1][to] = crate
			break
		}

		// empty col
		if i == len(*stack)-1 {
			(*stack)[i][to] = crate
		}
	}

	return nil
}

func readTop(stack *[][]byte) string {
	result := ""

	for j := range (*stack)[0] {
		for i := range *stack {
			curr := (*stack)[i][j]
			if curr != 0 {
				result += fmt.Sprintf("%c", curr)
				break
			}

			// empty col
			if i == len(*stack)-1 {
				result += " "
			}
		}
	}

	return result
}

func printStack(stack *[][]byte) {
	fmt.Println()
	for i := range *stack {
		for j := range (*stack)[i] {
			c := (*stack)[i][j]
			if c == 0 {
				c = ' '
			}
			fmt.Printf("%c ", c)
		}
		fmt.Println()
	}
}

func main() {
	//crateMover := CRATE_MOVER_9000
	crateMover := CRATE_MOVER_9001

	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	input := ""

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		input += line + "\n"
	}

	stack := input2Array(input)

	for scanner.Scan() {
		line := scanner.Text()
		err := makeMove(line, &stack, crateMover)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	top_crates := readTop(&stack)
	fmt.Printf("Top of stack: %s\n", top_crates)
}
