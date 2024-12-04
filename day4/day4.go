package main

import (
	"fmt"
	"os"
	"strings"
)

func count_xmas(row, col int, arr []string) int {
	str := "MAS"
	up := 0
	down := 0
	right := 0
	left := 0
	diag1 := 0
	diag2 := 0
	diag3 := 0
	diag4 := 0
	//up
	r := row - 1
	for ; r >= row-3 && row-3 >= 0; r-- {
		if arr[r][col] != str[row-1-r] {
			break
		}
	}

	if r == row-3-1 {
		up++
	}

	//down

	r = row + 1
	for ; r <= row+3 && row+3 < len(arr); r++ {
		if arr[r][col] != str[r-row-1] {
			break
		}
	}

	if r == row+3+1 {
		down++
	}

	//right

	c := col + 1
	for ; c <= col+3 && col+3 < len(arr[0]); c++ {
		if arr[row][c] != str[c-col-1] {
			break
		}
	}

	if c == col+3+1 {
		right++
	}

	//left
	c = col - 1
	for ; c >= col-3 && col-3 >= 0; c-- {
		if arr[row][c] != str[col-1-c] {
			break
		}
	}

	if c == col-3-1 {
		left++
	}

	// r- c+
	r = row - 1
	c = col + 1
	temp := ""
	for i := 0; i < 3 && r >= 0 && c < len(arr[0]); i++ {
		if arr[r][c] == str[i] {
			temp += string(arr[r][c])
			r--
			c++
		}
	}
	if str == temp {
		diag1++
	}

	// r- c-

	r = row - 1
	c = col - 1
	temp = ""
	for i := 0; i < 3 && r >= 0 && c >= 0; i++ {
		if arr[r][c] == str[i] {
			temp += string(arr[r][c])
			r--
			c--
		}
	}
	if str == temp {
		diag2++
	}

	// r+ c-
	r = row + 1
	c = col - 1
	temp = ""
	for i := 0; i < 3 && r < len(arr) && c >= 0; i++ {
		if arr[r][c] == str[i] {
			temp += string(arr[r][c])
			r++
			c--
		}
	}
	if str == temp {
		diag3++
	}

	// r+ c+
	r = row + 1
	c = col + 1
	temp = ""
	for i := 0; i < 3 && r < len(arr) && c < len(arr); i++ {
		if arr[r][c] == str[i] {
			temp += string(arr[r][c])
			r++
			c++
		}
	}
	if str == temp {
		diag4++
	}

	return up + down + left + right + diag1 + diag2 + diag3 + diag4

}

func is_x_mas(r, c int, arr []string) bool {
	val11 := arr[r-1][c-1]
	val12 := arr[r-1][c+1]
	val21 := arr[r+1][c-1]
	val22 := arr[r+1][c+1]

	// fmt.Println(val11, val12, val21, val22)

	str1 := string(val11) + string(val12)
	str2 := string(val21) + string(val22)

	// fmt.Println(str1, str2)

	if str1 == str2 && (str1 == "MS" || str1 == "SM") {
		return true
	}

	if str1 != str2 && ((str1 == "MM" && str2 == "SS") || (str1 == "SS" && str2 == "MM")) {
		return true
	}

	return false
}

func main() {
	data, err := os.ReadFile("day4_test.txt")

	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(string(data))

	input := string(data)

	arr := strings.Split(input, "\n")
	// fmt.Println(arr)

	res := 0
	for r, row := range arr {

		for col, c := range row {
			if c == 'A' && (r > 0 && r < len(arr)-1) && (col > 0 && col < len(arr[0])-1) {
				if is_x_mas(r, col, arr) {
					res++
				}
			}
		}

	}

	fmt.Println(res)

}
