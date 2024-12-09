package main

import (
	"fmt"
	"os"
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

	// fmt.Println(dec_input)

	start := 0
	last := len(dec_input) - 1

	for start < last {
		for dec_input[start] != "." {
			start++
		}
		for dec_input[last] == "." {
			last--
		}
		if start >= last {
			break
		}
		dec_input[start] = dec_input[last]
		dec_input[last] = "."
	}

	// fmt.Println(dec_input)
	output := 0
	for pos, val := range dec_input {
		if val == "." {
			break
		}
		id, _ := strconv.Atoi(val)
		if id < 0 {
			continue
		}
		output += id * pos
	}

	fmt.Println(output)

}
