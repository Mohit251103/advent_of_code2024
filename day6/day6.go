package main

import (
	"fmt"
	"os"
	"strings"
)

func get_start_coords(input []string) (int, int) {
	x, y := 0, 0
	for r, row := range input {
		for c, _ := range row {
			if input[r][c] == '^' {
				x, y = r, c
			}
		}
	}

	return x, y
}

func count_unique_path_points(x, y int, input []string) int {
	moves := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	res := 0
	curr := 0

	for true {

		if input[x][y] != 'X' {
			str := []byte(input[x])
			str[y] = 'X'
			input[x] = string(str)
			res++
		}
		mv := moves[curr]
		nx := x + mv[0]
		ny := y + mv[1]

		if nx < 0 || ny < 0 || nx == len(input) || ny == len(input[0]) {
			break
		}

		if input[nx][ny] == '#' {
			curr = (curr + 1) % 4
			continue
		}

		x = nx
		y = ny

	}

	return res

}

func is_cycle_present(x, y int, input []string) bool {
	// count_obs_visit := 0

	moves := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	curr := 0

	rec := make([][]int, len(input))
	for i := 0; i < len(rec); i++ {
		rec[i] = make([]int, len(input[0]))
	}

	for true {
		if input[x][y] != 'X' {
			str := []byte(input[x])
			str[y] = 'X'
			input[x] = string(str)

		}

		rec[x][y]++
		if rec[x][y] > 10 {
			break
		}

		mv := moves[curr]
		nx := x + mv[0]
		ny := y + mv[1]

		if nx < 0 || ny < 0 || nx == len(input) || ny == len(input[0]) {
			break
		}

		for input[nx][ny] == '#' {
			curr = (curr + 1) % 4
			mv = moves[curr]
			nx = x + mv[0]
			ny = y + mv[1]
		}

		x = nx
		y = ny
	}

	return rec[x][y] > 10
}

func main() {
	data, err := os.ReadFile("day6_test.txt")
	if err != nil {
		fmt.Println(err)
	}

	input := strings.Split(string(data), "\n")

	x, y := get_start_coords(input)

	// res_part1 := count_unique_path_points(x, y, input)
	res_part2 := 0
	for obsx, row := range input {
		for obsy, _ := range row {
			if input[obsx][obsy] == '#' || (x == obsx && y == obsy) {
				continue
			}
			str := make([]byte, len(input[obsx]))
			copy(str, []byte(input[obsx]))
			// temp := make([]byte, len(str))
			// copy(temp, str)
			str[obsy] = '#'
			temp_input := make([]string, len(input))
			copy(temp_input, input)
			temp_input[obsx] = string(str)
			// fmt.Println(input)``
			// fmt.Print(temp_input)
			if is_cycle_present(x, y, temp_input) {
				res_part2++
				// fmt.Println(true)
			}

			// input[obsx] = string(temp)

			// fmt.Println(input)
		}
	}

	fmt.Println(res_part2)
}
