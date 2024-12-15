package main

import (
	"fmt"
	"os"
	"strings"
)

func getMap(file string) []string {
	data, _ := os.ReadFile(file)
	var res []string
	for _, row := range strings.Split(strings.Split(string(data), "\n\n")[0], "\n") {
		curr_row := ""
		for _, ch := range row {
			if ch == '.' || ch == '#' {
				curr_row += string(ch) + string(ch)
			} else if ch == '@' {
				curr_row += "@."
			} else if ch == 'O' {
				curr_row += "[]"
			}
		}

		res = append(res, curr_row)
	}

	return res
}

func getMoves(file string) string {
	data, _ := os.ReadFile(file)
	str := ""
	for _, val := range strings.Split(strings.Split(string(data), "\n\n")[1], "\n") {
		str += val
	}

	return str
}

func get_robot_loc(mp []string) (int, int) {
	for r := 0; r < len(mp); r++ {
		for c := 0; c < len(mp[0]); c++ {
			if mp[r][c] == '@' {
				return r, c
			}
		}
	}

	return 0, 0
}

func completeMove(mp []string, sx, sy, cr, cy int) (int, int) {
	nx := cr + sx
	ny := cy + sy
	if mp[nx][ny] == '#' {
		return cr, cy
	} else if mp[nx][ny] == 'O' {
		// count_O := 0
		for mp[nx][ny] == 'O' {
			// count_O++
			nx += sx
			ny += sy
		}
		if mp[nx][ny] == '#' {
			return cr, cy
		} else if mp[nx][ny] == '.' {
			str_prev := make([]byte, len(mp[nx]))
			str_new := make([]byte, len(mp[nx]))
			copy(str_new, mp[nx])
			copy(str_prev, mp[cr])
			str_new[ny] = 'O'
			if nx == cr {
				str_new[cy+sy] = '@'
				str_new[cy] = '.'
				mp[nx] = string(str_new)
			} else {
				str_new[cy+sy] = '@'
				str_prev[cy] = '.'
				mp[nx] = string(str_new)
				mp[cr] = string(str_prev)
			}
			return cr + sx, cy + sy
		}
	}
	str_prev := make([]byte, len(mp[nx]))
	str_new := make([]byte, len(mp[nx]))
	copy(str_new, mp[nx])
	copy(str_prev, mp[cr])
	if nx == cr {
		str_new[cy+sy] = '@'
		str_new[cy] = '.'
		mp[nx] = string(str_new)
	} else {
		str_new[cy+sy] = '@'
		str_prev[cy] = '.'
		mp[nx] = string(str_new)
		mp[cr] = string(str_prev)
	}
	return cr + sx, cy + sy
}

func calc_res(mp []string) int {
	res := 0
	for r := 0; r < len(mp); r++ {
		for c := 0; c < len(mp[0]); c++ {
			if mp[r][c] == '[' {
				res += 100*r + c
			}
		}
	}

	return res
}

func display(mp []string) {
	for _, row := range mp {
		fmt.Println(row)
	}
}

func move_left(mp []string, r, c int) (int, int) {
	nc := c - 1
	if mp[r][nc] == '#' {
		return r, c
	} else if mp[r][nc] == '.' {
		new_row := make([]byte, len(mp[r]))
		copy(new_row, mp[r])
		new_row[nc] = '@'
		new_row[c] = '.'
		mp[r] = string(new_row)
		return r, nc
	} else {
		for mp[r][nc] == '[' || mp[r][nc] == ']' {
			nc--
		}
		if mp[r][nc] == '#' {
			return r, c
		} else {
			new_row := make([]byte, len(mp[r]))
			copy(new_row, mp[r])
			for tc := c - 2; tc >= nc; tc -= 2 {
				new_row[tc-1] = '['
				new_row[tc] = ']'
			}
			new_row[c-1] = '@'
			new_row[c] = '.'
			mp[r] = string(new_row)

			return r, c - 1
		}
	}
}
func move_right(mp []string, r, c int) (int, int) {
	nc := c + 1
	if mp[r][nc] == '#' {
		return r, c
	} else if mp[r][nc] == '.' {
		new_row := make([]byte, len(mp[r]))
		copy(new_row, mp[r])
		new_row[nc] = '@'
		new_row[c] = '.'
		mp[r] = string(new_row)
		return r, nc
	} else {
		for mp[r][nc] == '[' || mp[r][nc] == ']' {
			nc++
		}
		if mp[r][nc] == '#' {
			return r, c
		} else {
			new_row := make([]byte, len(mp[r]))
			copy(new_row, mp[r])
			for tc := c + 2; tc <= nc; tc += 2 {
				new_row[tc] = '['
				new_row[tc+1] = ']'
			}
			new_row[c+1] = '@'
			new_row[c] = '.'
			mp[r] = string(new_row)

			return r, c + 1
		}
	}
}
func move_up(mp []string, r, c int) (int, int) {

	nr := r - 1
	if mp[nr][c] == '#' {
		return r, c
	} else if mp[nr][c] == '.' {
		str1 := make([]byte, len(mp[0]))
		str2 := make([]byte, len(mp[0]))
		copy(str1, mp[nr])
		copy(str2, mp[r])
		str1[c] = '@'
		str2[c] = '.'
		mp[nr] = string(str1)
		mp[r] = string(str2)
		return nr, c
	} else {
		for (mp[nr][c] == '[' && mp[nr][c+1] == ']') || (mp[nr][c-1] == '[' && mp[nr][c] == ']') {
			nr--
		}
		if mp[nr][c] == '#' {
			return r, c
		} else {
			str1 := make([]byte, len(mp[0]))
			str2 := make([]byte, len(mp[0]))
			str3 := make([]byte, len(mp[0]))
			copy(str1, mp[nr])
			copy(str2, mp[r])
			copy(str3, mp[r-1])
			if mp[r-1][c] == '[' {
				str1[c] = '['
				str1[c+1] = ']'
				str3[c] = '@'
				str3[c+1] = '.'
			} else if mp[r-1][c] == ']' {
				str1[c] = ']'
				str1[c-1] = '['
				str3[c] = '@'
				str3[c-1] = '.'
			}
			str2[c] = '.'

			mp[nr] = string(str1)
			mp[r] = string(str2)
			mp[r-1] = string(str3)
			return r - 1, c
		}
	}

}
func move_down(mp []string, r, c int) (int, int) {

	nr := r + 1
	if mp[nr][c] == '#' {
		return r, c
	} else if mp[nr][c] == '.' {
		str1 := make([]byte, len(mp[0]))
		str2 := make([]byte, len(mp[0]))
		copy(str1, mp[nr])
		copy(str2, mp[r])
		str1[c] = '@'
		str2[c] = '.'
		mp[nr] = string(str1)
		mp[r] = string(str2)
		return nr, c
	} else {
		for (mp[nr][c] == '[' && mp[nr][c+1] == ']') || (mp[nr][c-1] == '[' && mp[nr][c] == ']') {
			nr++
		}
		if mp[nr][c] == '#' {
			return r, c
		} else {
			str1 := make([]byte, len(mp[0]))
			str2 := make([]byte, len(mp[0]))
			str3 := make([]byte, len(mp[0]))
			copy(str1, mp[nr])
			copy(str2, mp[r])
			copy(str3, mp[r+1])
			if mp[r+1][c] == '[' {
				str1[c] = '['
				str1[c+1] = ']'
				str3[c] = '@'
				str3[c+1] = '.'
			} else if mp[r+1][c] == ']' {
				str1[c] = ']'
				str1[c-1] = '['
				str3[c] = '@'
				str3[c-1] = '.'
			}
			str2[c] = '.'

			mp[nr] = string(str1)
			mp[r] = string(str2)
			mp[r+1] = string(str3)
			return r + 1, c
		}
	}

}

func main() {
	mp := getMap("day15_train.txt")
	moves := getMoves("day15_train.txt")

	display(mp)

	cr, cc := get_robot_loc(mp)
	// fmt.Println(cr, cc)

	for _, move := range moves {

		if move == '<' {
			cr, cc = move_left(mp, cr, cc)

			display(mp)
			// fmt.Println(cr, cc)
			// break
		} else if move == '>' {
			cr, cc = move_right(mp, cr, cc)

			display(mp)
			// break
			// fmt.Println(cr, cc)
		} else if move == '^' {
			cr, cc = move_up(mp, cr, cc)
			display(mp)
			break
		} else {
			cr, cc = move_down(mp, cr, cc)
			// display(mp)
			// break
			// fmt.Println(cr, cc)
		}

	}

	// display(mp)

	res := calc_res(mp)
	fmt.Println(res, len(moves))
	// fmt.Println(cr, cc)
}
