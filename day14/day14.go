package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func getInput(file string) ([]string, [][]string, [][]string) {
	data, _ := os.ReadFile(file)
	sep := strings.Split(string(data), "\n\n")
	dimensions := strings.Split(sep[1], ",")
	var p [][]string
	var v [][]string
	for _, row := range strings.Split(sep[0], "\n") {
		p_v := strings.Split(row, " ")
		p = append(p, strings.Split(strings.Split(p_v[0], "=")[1], ","))
		v = append(v, strings.Split(strings.Split(p_v[1], "=")[1], ","))
	}

	return dimensions, p, v
}

func main() {
	dimensions, p, v := getInput("day14_test.txt")
	// fmt.Println(v, p)
	// fmt.Println(p, v, dimensions)
	x, _ := strconv.Atoi(dimensions[0])
	y, _ := strconv.Atoi(dimensions[1])

	// fmt.Println(x, y)

	midx := x / 2
	midy := y / 2
	// fmt.Println(midx, midy)

	var coords [][]int

	min_sf := math.MaxInt
	res := 0
	for sec := 0; sec <= x*y; sec++ {
		room := make([][]int, y)
		quad1 := 0
		quad2 := 0
		quad3 := 0
		quad4 := 0
		for i := 0; i < y; i++ {
			arr := make([]int, x)
			room[i] = arr
		}
		for i := 0; i < len(p); i++ {
			vx, _ := strconv.Atoi(v[i][0])
			vy, _ := strconv.Atoi(v[i][1])

			// fmt.Println(vx, vy)

			sx, _ := strconv.Atoi(p[i][0])
			sy, _ := strconv.Atoi(p[i][1])

			distx := (sx + (vx * sec)) % x
			disty := (sy + (vy * sec)) % y

			// fmt.Println(distx, disty)

			var coord []int
			if distx < 0 {
				coord = append(coord, x+distx)
			} else {
				coord = append(coord, distx)
			}

			if disty < 0 {
				coord = append(coord, y+disty)
			} else {
				coord = append(coord, disty)
			}

			coords = append(coords, coord)

			cx := coord[0]
			cy := coord[1]

			room[cy][cx] = 1

			// fmt.Println(cx, cy)
			if cx == midx || cy == midy {
				continue
			}

			if cx > midx && cy > midy {
				quad4++
			} else if cx < midx && cy > midy {
				quad3++
			} else if cx < midx && cy < midy {
				quad2++
			} else if cx > midx && cy < midy {
				quad1++
			}
		}

		// for r := 0; r < y; r++ {
		// 	fmt.Println(room[r])
		// 	// fmt.Println()
		// }
		// fmt.Println()

		// fmt.Println(room, "\n")
		sf := quad1 * quad2 * quad3 * quad4
		if sf < min_sf {
			min_sf = sf
			res = sec
		}
	}
	fmt.Println(min_sf, res)

	// fmt.Println(quad1, quad2, quad3, quad4)
}
