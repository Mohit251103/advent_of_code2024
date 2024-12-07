package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isValid(test string, eq []string, idx, sum int) bool {
	if idx == len(eq) {
		chk, err := strconv.Atoi(test)
		if err != nil {
			fmt.Println(err)
		}

		if sum == chk {
			return true
		}
		return false
	}

	leave := isValid(test, eq, idx+1, sum)

	el, err := strconv.Atoi(eq[idx])
	if err != nil {
		fmt.Println(err)
	}
	take_sum := isValid(test, eq, idx+1, sum+el)
	take_mul := isValid(test, eq, idx+1, sum*el)

	return leave || take_sum || take_mul
}

func main() {
	data, err := os.ReadFile("day7_test.txt")
	if err != nil {
		fmt.Println(err)
	}

	temp := strings.Split(string(data), "\n")

	var test []string
	var equations [][]string
	for _, r := range temp {
		arr := strings.Split(r, ":")
		test = append(test, arr[0])
		equations = append(equations, strings.Split(strings.Trim(arr[1], " "), " "))
	}

	sum_part1 := 0
	for i, eq := range equations {
		start, err := strconv.Atoi(eq[0])
		if err != nil {
			fmt.Println(err)
		}
		if isValid(test[i], eq, 1, start) {
			n, err := strconv.Atoi(test[i])
			if err != nil {
				fmt.Println(err)
			}
			sum_part1 += n
		}
	}

	fmt.Println(sum_part1)
}
