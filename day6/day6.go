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

func main() {
	data, err := os.ReadFile("day6_test.txt")
	if err != nil {
		fmt.Println(err)
	}

	input := strings.Split(string(data), "\n")

	x, y := get_start_coords(input)

	res_part1 := count_unique_path_points(x, y, input)

	fmt.Println(res_part1)
}
