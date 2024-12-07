package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isValid(test string, eq []string, idx int, sum string) bool {
	if idx == len(eq) {
		// chk, err := strconv.Atoi(test)
		// if err != nil {
		// 	fmt.Println(err)
		// }

		if sum == test {
			return true
		}
		return false
	}

	leave := isValid(test, eq, idx+1, sum)

	el, err := strconv.Atoi(eq[idx])
	if err != nil {
		fmt.Println(err)
	}

	temp_sum, err := strconv.Atoi(sum)
	if err != nil {
		fmt.Println(sum)
		fmt.Println(err, "here")
	}

	take_sum := isValid(test, eq, idx+1, strconv.Itoa(temp_sum+el))
	take_mul := isValid(test, eq, idx+1, strconv.Itoa(temp_sum*el))
	take_conc := isValid(test, eq, idx+1, sum+eq[idx])

	return leave || take_sum || take_mul || take_conc
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

	sum := 0
	for i, eq := range equations {

		// --- part 1 --

		// start, err := strconv.Atoi(eq[0])
		// if err != nil {
		// 	fmt.Println(err)
		// }

		// --- part 1 end ---

		if isValid(test[i], eq, 1, eq[0]) {
			n, err := strconv.Atoi(test[i])
			if err != nil {
				fmt.Println(err)
			}
			sum += n
		}
	}

	fmt.Println(sum)
}
