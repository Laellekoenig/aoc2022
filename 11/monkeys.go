package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type Monkey struct {
	num              int
	items            []int
	operation        func(int, int) int
	snd_param        int
	snd_param_is_old bool
	mod              int
	true_mokey       int
	false_monkey     int
}

func parse_monkey(monkey_str []string) (*Monkey, error) {
	// parse monkey number
	num := int(strings.Split(monkey_str[0], " ")[1][0] - '0')

	// parse starting items_str
	items_str := strings.Split(monkey_str[1], " ")[4:]
	items := make([]int, len(items_str), 20)

	for i, item := range items_str {
		item = strings.Split(item, ",")[0]
		item_conv, err := strconv.Atoi(item)
		if err != nil {
			return nil, err
		}
		items[i] = item_conv
	}

	// parse operation
	op_str := strings.Split(monkey_str[2], "= ")[1]
	op_arr := strings.Split(op_str, " ")

	var f func(int, int) int
	if op_arr[1] == "*" {
		f = func(x, y int) int { return x * y }
	} else if op_arr[1] == "+" {
		f = func(x, y int) int { return x + y }
	}

	snd_param_is_old := false
	var snd_param int
	if op_arr[2] == "old" {
		snd_param_is_old = true
	} else {
		var err error
		snd_param, err = strconv.Atoi(op_arr[2])
		if err != nil {
			return nil, err
		}
	}

	// parse divisible by
	mod, err := strconv.Atoi(strings.Split(monkey_str[3], " ")[5])
	if err != nil {
		return nil, err
	}

	// parse if true monkey
	true_monkey, err := strconv.Atoi(strings.Split(monkey_str[4], " ")[9])
	if err != nil {
		return nil, err
	}

	// parse if false monkey
	false_monkey, err := strconv.Atoi(strings.Split(monkey_str[5], " ")[9])
	if err != nil {
		return nil, err
	}

	monkey := Monkey{
		num:              num,
		items:            items,
		operation:        f,
		snd_param:        snd_param,
		snd_param_is_old: snd_param_is_old,
		mod:              mod,
		true_mokey:       true_monkey,
		false_monkey:     false_monkey,
	}

	return &monkey, nil
}

func gcd(x, y int) int {
	for y != 0 {
		t := y
		y = x % y
		x = t
	}

	return x
}

func lcm(x, y int) int {
	return int(math.Abs(float64(x*y))) / gcd(x, y)
}

func get_lcm_arr(a *[]int) int {
	result := 1

	for _, item := range *a {
		result = lcm(result, item)
	}

	return result
}

func main() {
	input_bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	input := string(input_bytes)
	lines := strings.Split(input, "\n")
	size := 7 // number of lines per monkey
	num_monkeys := len(lines) / 7

	monkeys := make([]*Monkey, num_monkeys)

	for i := 0; i < num_monkeys; i++ {
		m, err := parse_monkey(lines[i*size : (i+1)*size-1])
		if err != nil {
			fmt.Println(err)
			return
		}

		monkeys[i] = m
	}

	// part one
	//num_rounds := 20

	// part two
	mods := make([]int, len(monkeys))
	for i, m := range monkeys {
		mods[i] = m.mod
	}

	// determine lcm of all mods
	lcm := get_lcm_arr(&mods)

	// part two
	num_rounds := 10000

	num_inspected := make(map[*Monkey]int)

	for i := 0; i < num_rounds; i++ {
		for _, monkey := range monkeys {
			// for every monkey
			for _, item := range monkey.items {
				// for every item
				num_inspected[monkey]++

				// apply operation
				if monkey.snd_param_is_old {
					item = monkey.operation(item, item)
				} else {
					item = monkey.operation(item, monkey.snd_param)
				}

				// part one
				// div by 3
				// item /= 3

				// part two
				item %= lcm

				// if else -> throw to next monkey
				var next_monkey int
				if item%monkey.mod == 0 {
					next_monkey = monkey.true_mokey
				} else {
					next_monkey = monkey.false_monkey
				}

				for _, m := range monkeys {
					if m.num == next_monkey {
						m.items = append(m.items, item)
					}
				}
			}

			// monkey threw all items -> clear
			monkey.items = make([]int, 0, 20)
		}
	}

	fst := -1
	snd := -1

	for key := range num_inspected {
		val := num_inspected[key]

		if val > fst {
			snd = fst
			fst = val
		} else if val > snd {
			snd = val
		}
	}

	monkey_business := fst * snd

	fmt.Printf("Level of monkey business after %d rounds: %d\n", num_rounds, monkey_business)
}
