package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("day9_test.txt")
	if err != nil {
		fmt.Println(err)
	}

	input := strings.Split(string(data), "")

	var dec_input []string
	id := 0
	for i := 0; i < len(input); i += 2 {
		fb, _ := strconv.Atoi(input[i]) // file block size
		for fb != 0 {
			dec_input = append(dec_input, strconv.Itoa(id))
			fb--
		}

		if i+1 < len(input) {
			free, _ := strconv.Atoi(input[i+1]) // free space
			for free != 0 {
				dec_input = append(dec_input, ".")
				free--
			}
		}

		id++
	}

	// <--- part 1 ---> start

	// start := 0
	// last := len(dec_input) - 1

	// free_space := 0
	// need := 0
	// for start < last {
	// 	for dec_input[start] != "." && free_space == 0 {
	// 		start++
	// 	}

	// 	for dec_input[start] == "." {
	// 		free_space++
	// 		start++
	// 	}

	// 	for dec_input[last] == "." && need == 0 {
	// 		last--
	// 	}
	// 	if start >= last {
	// 		break
	// 	}

	// 	ref := dec_input[last]
	// 	for start < last && dec_input[last] != "." && dec_input[last] == ref {
	// 		need++
	// 		last--
	// 	}

	// 	if start >= last {
	// 		break
	// 	}

	// 	if free_space >= need {
	// 		for need != 0 {
	// 			dec_input[start-free_space] = ref
	// 			dec_input[last+need] = "."
	// 			free_space--
	// 			need--
	// 		}
	// 	} else {
	// 		free_space = 0
	// 	}
	// }
	// fmt.Println(dec_input)

	// <--- part 1 ---> end

	// <--- part 2 ---> start

	var sorted_id []int
	mp := make(map[int]int)
	for i := len(dec_input) - 1; i >= 0; i-- {
		if dec_input[i] == "." {
			continue
		}
		val, _ := strconv.Atoi(dec_input[i])
		if mp[val] == 0 {
			sorted_id = append(sorted_id, val)
		}
		mp[val]++
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sorted_id)))

	for idx, id := range sorted_id {
		need := mp[sorted_id[idx]]
		free_space := 0
		for i := 0; i < len(dec_input); i++ {
			if strconv.Itoa(id) == dec_input[i] {
				break
			}
			if dec_input[i] == "." {
				free_space++
			} else {
				free_space = 0
			}

			if free_space == need {
				// fmt.Println(id, i, i-free_space+1, free_space)
				for j := len(dec_input) - 1; j >= 0; j-- {
					if dec_input[j] == strconv.Itoa(id) {
						dec_input[j] = "."
					}
				}
				for free_space != 0 {
					dec_input[i-free_space+1] = strconv.Itoa(id)
					free_space--
				}
				// fmt.Println(dec_input)
				// fmt.Println()

				break
			}
		}

	}

	// <--- part 2 ---> end

	output := 0
	for pos, val := range dec_input {
		if val == "." {
			continue
		}
		id, _ := strconv.Atoi(val)
		if id < 0 {
			continue
		}
		output += id * pos
	}

	fmt.Println(output)

}
