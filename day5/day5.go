package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("day5_test.txt")

	if err != nil {
		fmt.Println(err)
	}

	rules := strings.Split(string(data), "\n\n")[0]
	input := strings.Split(string(data), "\n\n")[1]

	mp := make(map[int][]int)

	for _, rule := range strings.Split(rules, "\n") {
		temp := strings.Split(rule, "|")
		// fmt.Println(temp)
		key, err := strconv.Atoi(temp[0])
		if err != nil {
			fmt.Println(err)
		}
		val, err := strconv.Atoi(temp[1])
		if err != nil {
			fmt.Println(err)
		}

		arr := mp[key]
		arr = append(arr, val)
		mp[key] = arr
	}
	// fmt.Println(mp)

	// fmt.Println(rules)
	// fmt.Println(input)

	var is_after func(int, int) bool

	is_after = func(key, val int) bool {
		temp_arr := mp[key]
		for _, n := range temp_arr {
			if val == n {
				return true
			}
		}

		return false
	}

	var wrong_ones [][]string // fors second part

	sum := 0
	for _, updates := range strings.Split(input, "\n") {
		isvalid := true
		arr := strings.Split(updates, ",")
		for idx, _ := range arr {

			if idx == 0 {
				continue
			}

			key, _ := strconv.Atoi(arr[idx-1])
			val, _ := strconv.Atoi(arr[idx])

			if !is_after(key, val) {
				isvalid = false
			}

		}

		if isvalid {
			n := len(arr)
			mid, err := strconv.Atoi(arr[n/2])
			if err != nil {
				fmt.Println(err)
			}
			sum += mid
		} else { // second part starts from here
			wrong_ones = append(wrong_ones, arr)
		}
	}

	sum_wrongs := 0

	for _, arr := range wrong_ones {
		for i := 0; i < len(arr); i++ {
			for idx, _ := range arr {

				if idx == 0 {
					continue
				}

				key, _ := strconv.Atoi(arr[idx-1])
				val, _ := strconv.Atoi(arr[idx])

				if !is_after(key, val) {
					arr[idx], arr[idx-1] = arr[idx-1], arr[idx]
				}

			}
		}

		// fmt.Println(arr)

		n := len(arr) / 2
		x, _ := strconv.Atoi(arr[n])
		sum_wrongs += x
	}

	fmt.Println(sum_wrongs)
	// fmt.Println(sum)
}
