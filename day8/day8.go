package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func euclidean(x, y, x2, y2 int) int {
	d1 := x2 - x
	d2 := y2 - y
	return int(math.Sqrt(float64(d1*d1) + float64(d2*d2)))
}

func calcSlope(x, y, x1, y1 int) float64 {
	return float64(y1-y) / float64(x1-x)
}

func findAntinodes(key rune, coords [][]int, field []string, set map[interface{}]bool) int {
	count := 0
	for i := 0; i < len(coords); i++ {
		for j := i + 1; j < len(coords); j++ {
			target := euclidean(coords[i][0], coords[i][1], coords[j][0], coords[j][1])
			slope1 := calcSlope(coords[i][0], coords[i][1], coords[j][0], coords[j][1])
			// slope2 := math.Abs(slope1)
			// fmt.Println(coords[i], coords[j], target, slope1, slope2)

			t := 0
			for r, row := range field {
				for c, _ := range row {
					// fmt.Println(r, c)
					// if field[r][c] == 'A' {
					// 	// 	fmt.Println(string(key), r, c, coords[i][0], coords[i][1], coords[j][0], coords[j][1], "slope: ", calcSlope(r, c, coords[i][0], coords[i][1]), slope1)
					// 	fmt.Println(r, c)
					// }
					if field[r][c] == '#' || (r == coords[i][0] && c == coords[i][1]) || (r == coords[j][0] && c == coords[j][1]) {
						continue
					}
					k := strconv.Itoa(r) + "#" + strconv.Itoa(c)

					if (((calcSlope(r, c, coords[i][0], coords[i][1]) == slope1) && euclidean(coords[i][0], coords[i][1], r, c) == target) || ((calcSlope(coords[j][0], coords[j][1], r, c) == slope1) && euclidean(coords[j][0], coords[j][1], r, c) == target)) && !set[k] {
						t++
						set[k] = true
						// fmt.Println("slope1", r, c)
					}

					// if (((calcSlope(r, c, coords[i][0], coords[i][1]) == slope2) && euclidean(coords[i][0], coords[i][1], r, c) == target) || ((calcSlope(coords[j][0], coords[j][1], r, c) == slope2) && euclidean(coords[j][0], coords[j][1], r, c) == target)) && !set[k] {
					// 	t++
					// 	set[k] = true
					// 	fmt.Println("slope2", r, c)
					// }
				}

				if t == 2 {
					break
				}
			}
			count += t
		}
	}

	return count
}

func main() {
	data, err := os.ReadFile("day8_test.txt")
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println(string(data))

	// going to use euclidean distance for this one , involves some math

	mp := make(map[rune][][]int)

	input := strings.Split(string(data), "\n")
	for r, row := range input {
		for c, ch := range row {
			if ch == '.' {
				continue
			}
			coord := []int{r, c}
			mp[ch] = append(mp[ch], coord)
		}
	}

	// fmt.Println(mp)
	set := make(map[interface{}]bool)
	count := 0
	for key, val := range mp {
		n := findAntinodes(key, val, input, set)
		count += n
	}

	fmt.Println(count)

}
