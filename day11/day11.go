package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func removeLeadingZeroes(str string) string {
	res := ""
	ok := false
	for _, val := range str {
		if val != '0' {
			ok = true
		}
		if ok {
			res += string(val)
		}
	}

	if len(res) == 0 {
		return "0"
	}
	return res
}

func solve(el string, result chan int, mp map[string]int, wg *sync.WaitGroup) {
	defer wg.Done()
	res := []string{el}
	for i := 0; i < 75; i++ {
		var arr []string
		for _, val := range res {
			if len(val)%2 == 0 {
				fh := val[0 : len(val)/2]
				sh := val[len(val)/2:]
				fh = removeLeadingZeroes(fh)
				sh = removeLeadingZeroes(sh)
				arr = append(arr, fh, sh)
			} else if val == "0" {
				arr = append(arr, "1")
			} else {
				n, _ := strconv.Atoi(val)
				arr = append(arr, strconv.Itoa(n*2024))
			}
		}
		res = arr
	}
	result <- len(res)
}

func recursive(val string, mp map[string]int, turn int) int {
	if turn == 75 {
		return 1
	}

	tstr := strconv.Itoa(turn)
	if mp[val+"#"+tstr] != 0 {
		return mp[val+"#"+tstr]
	}
	res := 0
	if len(val)%2 == 0 {
		fh := val[0 : len(val)/2]
		sh := val[len(val)/2:]
		fh = removeLeadingZeroes(fh)
		sh = removeLeadingZeroes(sh)
		res = recursive(fh, mp, turn+1) + recursive(sh, mp, turn+1)
	} else if val == "0" {
		res = recursive("1", mp, turn+1)
	} else {
		n, _ := strconv.Atoi(val)
		s := strconv.Itoa(n * 2024)
		res += recursive(s, mp, turn+1)
	}

	mp[val+"#"+tstr] = res
	return mp[val+"#"+tstr]
}

func main() {
	data, _ := os.ReadFile("day11_test.txt")
	input := strings.Split(string(data), " ")

	// var wg sync.WaitGroup
	mp := make(map[string]int)

	// resultChan := make(chan int, len(input))
	res := 0
	// for _, val := range input {
	// 	wg.Add(1)
	// 	go solve(val, resultChan, mp, &wg)
	// }

	for _, val := range input {
		res += recursive(val, mp, 0)
	}

	// wg.Wait()
	// close(resultChan)

	// for result := range resultChan {
	// 	res += result
	// }
	fmt.Println(res)
}
