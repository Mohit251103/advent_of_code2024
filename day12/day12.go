package main

import (
	"fmt"
	"os"
	"strings"
)

func cal_area_per(r, c int, crop byte, field []string, visited [][]byte) (int, int) {

	if visited[r][c] == '1' || field[r][c] != crop {
		return 0, 0
	}

	area := 0
	perimeter := 4

	if field[r][c] == crop {
		area += 1
	}

	visited[r][c] = '1'

	if r-1 >= 0 && field[r-1][c] == crop {
		perimeter--
		top, pt := cal_area_per(r-1, c, crop, field, visited)
		area += top
		perimeter += pt
	}
	if r+1 < len(field) && field[r+1][c] == crop {
		perimeter--
		down, pd := cal_area_per(r+1, c, crop, field, visited)
		area += down
		perimeter += pd
	}
	if c-1 >= 0 && field[r][c-1] == crop {
		perimeter--
		left, pl := cal_area_per(r, c-1, crop, field, visited)
		area += left
		perimeter += pl
	}
	if c+1 < len(field) && field[r][c+1] == crop {
		perimeter--
		right, pr := cal_area_per(r, c+1, crop, field, visited)
		area += right
		perimeter += pr
	}

	return area, perimeter

}

func calc_sides(up, down, left, right [][]int) int {
	res := 0
	// up
	for r, _ := range up {
		newside := true
		sides := 0
		for c := 0; c < len(up[0]); c++ {
			if newside && up[r][c] == 1 {
				newside = false
				sides++
			} else if !newside && up[r][c] == 0 {
				newside = true
			}
		}
		res += sides
	}

	// down
	for r, _ := range down {
		newside := true
		sides := 0
		for c := 0; c < len(down[0]); c++ {
			if newside && down[r][c] == 1 {
				newside = false
				sides++
			} else if !newside && down[r][c] == 0 {
				newside = true
			}
		}
		res += sides
	}

	// left
	for c, _ := range left[0] {
		newside := true
		sides := 0
		for r := 0; r < len(left); r++ {
			if newside && left[r][c] == 1 {
				newside = false
				sides++
			} else if !newside && left[r][c] == 0 {
				newside = true
			}
		}
		res += sides
	}

	// right
	for c, _ := range left[0] {
		newside := true
		sides := 0
		for r := 0; r < len(right); r++ {
			if newside && right[r][c] == 1 {
				newside = false
				sides++
			} else if !newside && right[r][c] == 0 {
				newside = true
			}
		}
		res += sides
	}

	return res
}

func clear_sides(up, down, left, right [][]int) {
	for r, row := range up {
		for c, _ := range row {
			up[r][c] = 0
		}
	}
	for r, row := range up {
		for c, _ := range row {
			down[r][c] = 0
		}
	}
	for r, row := range up {
		for c, _ := range row {
			left[r][c] = 0
		}
	}
	for r, row := range up {
		for c, _ := range row {
			right[r][c] = 0
		}
	}
}

func main() {
	data, _ := os.ReadFile("day12_test.txt")
	input := strings.Split(string(data), "\n")
	price := 0
	visited := make([][]byte, len(input))
	for i := 0; i < len(input); i++ {
		visited[i] = make([]byte, len(input[0]))
	}
	// fmt.Println(visited)
	up := make([][]int, len(input))
	down := make([][]int, len(input))
	left := make([][]int, len(input))
	right := make([][]int, len(input))

	for i := 0; i < len(input); i++ {
		up[i] = make([]int, len(input[0]))
		down[i] = make([]int, len(input[0]))
		left[i] = make([]int, len(input[0]))
		right[i] = make([]int, len(input[0]))
	}

	var calc_area_side func(int, int, byte, []string, [][]byte) int

	calc_area_side = func(r, c int, crop byte, field []string, visited [][]byte) int {
		if r >= len(field) || r < 0 || c < 0 || c >= len(field) || field[r][c] != crop || visited[r][c] == '1' {
			return 0
		}

		area := 0

		if field[r][c] == crop {
			area += 1
		}

		if r-1 < 0 || field[r-1][c] != crop {
			up[r][c] = 1
		}
		if r+1 >= len(input) || field[r+1][c] != crop {
			down[r][c] = 1
		}
		if c-1 < 0 || field[r][c-1] != crop {
			left[r][c] = 1
		}
		if c+1 >= len(input) || field[r][c+1] != crop {
			right[r][c] = 1
		}

		visited[r][c] = '1'

		top := calc_area_side(r-1, c, crop, field, visited)
		down := calc_area_side(r+1, c, crop, field, visited)
		left := calc_area_side(r, c-1, crop, field, visited)
		right := calc_area_side(r, c+1, crop, field, visited)

		area += top + down + left + right
		return area
	}

	for r, row := range input {
		for c, _ := range row {
			if input[r][c] != '1' {
				// area, perimeter := cal_area_per(r, c, input[r][c], input, visited)
				// price += area * perimeter
				area := calc_area_side(r, c, input[r][c], input, visited)
				sides := calc_sides(up, down, left, right)
				price += area * sides
				clear_sides(up, down, left, right)
			}
		}
	}

	fmt.Println(price)
}
