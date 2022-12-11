package main

import (
	"bufio"
	"fmt"
	"os"
)

func read_file(fn string) ([][]int, error) {
	f, err := os.Open(fn)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	line := scanner.Text()

	topo := make([][]int, len(line))
	for i := range topo {
		topo[i] = make([]int, len(line))
	}

	i := 0
	for {
		for j, r := range line {
			topo[i][j] = int(r - '0')
		}
		if !scanner.Scan() {
			break
		}
		i++
		line = scanner.Text()
	}

	return topo, nil
}

func is_visible(i, j int, topo *[][]int) bool {
	val := (*topo)[i][j]

	left := (*topo)[i][:j]
	right := (*topo)[i][j+1:]

	f := flip_2d_array(topo)
	top := f[j][:i]
	bottom := f[j][i+1:]

	if val > max_slice(&left) ||
		val > max_slice(&right) ||
		val > max_slice(&top) ||
		val > max_slice(&bottom) {
		return true
	}

	return false
}

func slice_scenic_score(s *[]int) int {
	var dist int

	if len(*s) == 1 {
		return dist
	}

	for _, v := range (*s)[1:] {
		dist++
		if v >= (*s)[0] {
			break
		}
	}

	return dist
}

func get_scenic_score(i, j int, topo *[][]int) int {
	left := (*topo)[i][:j+1]
	reverse_slice(&left)
	score := slice_scenic_score(&left)

	right := (*topo)[i][j:]
	score *= slice_scenic_score(&right)

	f := flip_2d_array(topo)

	top := f[j][:i+1]
	reverse_slice(&top)
	score *= slice_scenic_score(&top)

	bottom := f[j][i:]
	score *= slice_scenic_score(&bottom)

	return score
}

func max_slice(s *[]int) int {
	if len(*s) == 0 {
		panic(fmt.Errorf("Cannot take max of empty slice.\n"))
	}

	max := (*s)[0]

	if len(*s) == 1 {
		return max
	}

	for _, v := range (*s)[1:] {
		if v > max {
			max = v
		}
	}

	return max
}

func reverse_slice(s *[]int) {
	l := len(*s)
	reverse := make([]int, l)

	for i, v := range *s {
		reverse[l-1-i] = v
	}

	*s = reverse
}

func flip_2d_array(a *[][]int) [][]int {
	flip := make([][]int, len(*a))
	for i := range flip {
		flip[i] = make([]int, len((*a)[0]))
	}

	for i := 0; i < len(flip); i++ {
		for j := 0; j < len(flip[0]); j++ {
			flip[i][j] = (*a)[j][i]
		}
	}

	return flip
}

func main() {
	file := "input.txt"
	topo, err := read_file(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	tree_count := 2*(len(topo)+len(topo[0])) - 4 // trees on edge

	for i := 1; i < len(topo)-1; i++ {
		for j := 1; j < len(topo[0])-1; j++ {
			if is_visible(i, j, &topo) {
				tree_count++
			}
		}
	}

	fmt.Printf("Number of visible trees is: %d\n", tree_count)

	max_score := -1

	for i := 0; i < len(topo); i++ {
		for j := 0; j < len(topo[0]); j++ {
			v := get_scenic_score(i, j, &topo)
			if v > max_score {
				max_score = v
			}
		}
	}

	fmt.Printf("The maximum scenic score is %d\n", max_score)
}
