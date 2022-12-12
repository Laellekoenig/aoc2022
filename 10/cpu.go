package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type CPU struct {
	cycle               int
	regx                int
	signal_strength_sum int
	pixels              [][]int
}

func new_cpu() *CPU {
	pixels := make([][]int, 6)
	for i := range pixels {
		pixels[i] = make([]int, 40)
	}

	cpu := CPU{
		cycle:  1,
		regx:   1,
		pixels: pixels,
	}

	return &cpu
}

func (cpu *CPU) noop() {
	cpu.draw_pixel()
	cpu.cycle++
	cpu.check_cycle_num()
}

func (cpu *CPU) addx(x int) {
	// first cycle
	cpu.draw_pixel()
	cpu.cycle++
	cpu.check_cycle_num()

	// second cycle
	cpu.draw_pixel()
	cpu.cycle++
	cpu.regx += x
	cpu.check_cycle_num()
}

func (cpu *CPU) draw_pixel() {
	j := (cpu.cycle - 1) % 40
	var i int = (cpu.cycle - 1) / 40

	if cpu.regx-1 == j ||
		cpu.regx == j ||
		cpu.regx+1 == j {
		cpu.pixels[i][j] = 1
	}
}

func (cpu *CPU) check_cycle_num() {
	if cpu.cycle == 20 ||
		cpu.cycle == 60 ||
		cpu.cycle == 100 ||
		cpu.cycle == 140 ||
		cpu.cycle == 180 ||
		cpu.cycle == 220 {
		cpu.signal_strength_sum += cpu.cycle * cpu.regx
	}
}

func (cpu *CPU) render() string {
	result := ""
	for _, row := range cpu.pixels {
		for _, pixel := range row {
			if pixel == 0 {
				result += "."
			} else {
				result += "#"
			}
		}
		result += "\n"
	}

	return result
}

func (cpu *CPU) exec(cmd string) error {
	split := strings.Split(cmd, " ")

	op := split[0]

	if op == "" {
		return nil
	}

	if op == "noop" {
		cpu.noop()
		return nil
	}

	arg, err := strconv.Atoi(split[1])
	if err != nil {
		return err
	}

	cpu.addx(arg)
	return nil
}

func main() {
	input_bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	cmds := strings.Split(string(input_bytes), "\n")
	cpu := new_cpu()

	for _, cmd := range cmds {
		err := cpu.exec(cmd)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	fmt.Printf("The sum of all signal strengths is %d\n", cpu.signal_strength_sum)
	fmt.Print(cpu.render())
}
