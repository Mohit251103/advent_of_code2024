package main

import (
	"fmt"
	"os"
	"strings"
)

func count_score(r, c int, input []string, visited [][]bool, prev byte) int {
	if r >= len(input) || r < 0 || c >= len(input[0]) || c < 0 || visited[r][c] || rune(prev)+1 != rune(input[r][c]) {
		return 0
	}

	if input[r][c] == '9' {
		return 1
	}

	visited[r][c] = true
	left := count_score(r, c-1, input, visited, input[r][c])
	right := count_score(r, c+1, input, visited, input[r][c])
	up := count_score(r-1, c, input, visited, input[r][c])
	down := count_score(r+1, c, input, visited, input[r][c])
	visited[r][c] = false

	return left + right + up + down
}

func clear_all(input []string, visited [][]bool) {
	for r, row := range input {
		for c, _ := range row {
			if visited[r][c] {
				visited[r][c] = false
			}
		}
	}
}

func main() {
	data, _ := os.ReadFile("day10_test.txt")
	// fmt.Println(string(data))

	input := strings.Split(string(data), "\n")
	visited := make([][]bool, len(input))
	for i := 0; i < len(visited); i++ {
		arr := make([]bool, len(input[0]))
		visited[i] = arr
	}

	res := 0
	for r, row := range input {
		for c, _ := range row {
			if input[r][c] == '0' {
				visited[r][c] = true
				// fmt.Println("reached 1")
				left := count_score(r, c-1, input, visited, input[r][c])
				// fmt.Println("reached 2")
				right := count_score(r, c+1, input, visited, input[r][c])
				// fmt.Println("reached 3")
				up := count_score(r-1, c, input, visited, input[r][c])
				// fmt.Println("reached 4")
				down := count_score(r+1, c, input, visited, input[r][c])
				// fmt.Println("reached 5")
				visited[r][c] = false
				res += left + right + up + down
				// fmt.Println(visited)
				clear_all(input, visited)
				// fmt.Println(left + right + up + down)
			}
		}
	}

	fmt.Println(res)
}
