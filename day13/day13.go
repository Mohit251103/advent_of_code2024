package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
)

func clear_input() [][][]int {
	data, _ := os.ReadFile("day13_test.txt")
	str := strings.Split(string(data), "\n\n")

	var res [][][]int
	for _, each := range str {
		temp := strings.Split(each, "\n")
		btnA := strings.Split(strings.Split(temp[0], ":")[1], ",")
		btnB := strings.Split(strings.Split(temp[1], ":")[1], ",")
		price := strings.Split(strings.Split(temp[2], ":")[1], ",")

		var arr [][]int
		var el []int
		for _, btn := range btnA {
			move, _ := strconv.Atoi(strings.Split(btn, "+")[1])
			el = append(el, move)
		}
		cp := make([]int, len(el))
		copy(cp, el)
		arr = append(arr, cp)
		el = el[:0]

		for _, btn := range btnB {
			move, _ := strconv.Atoi(strings.Split(btn, "+")[1])
			el = append(el, move)
		}
		cp2 := make([]int, len(el))
		copy(cp2, el)
		arr = append(arr, cp2)
		el = el[:0]

		for _, val := range price {
			c, _ := strconv.Atoi(strings.Split(val, "=")[1])
			el = append(el, c)
		}
		cp3 := make([]int, len(el))
		copy(cp3, el)
		arr = append(arr, cp3)

		res = append(res, arr)
	}

	return res

}

func divide_and_conquer(start, limit int, machine [][]int) int {
	btnA := machine[0]
	btnB := machine[1]
	priceX := machine[2][0] + 10000000000000
	priceY := machine[2][1] + 10000000000000
	tokens := math.MaxInt

	for start <= limit {
		mid := (start + limit) / 2 // number of b presses
		fmt.Println(start, limit)
		Y_b := btnB[1] * mid
		X_b := btnB[0] * mid
		if Y_b <= priceY && X_b <= priceX {
			Y_a := priceY - Y_b
			X_a := priceX - X_b
			if Y_a%btnA[1] != 0 || X_a%btnA[0] != 0 || Y_a/btnA[1] != X_a/btnA[0] {
				var wg sync.WaitGroup
				wg.Add(1)
				go func() {
					defer wg.Done()
					tokens = min(tokens, divide_and_conquer(start, mid-1, machine))
				}()

				wg.Add(1)
				go func() {
					defer wg.Done()
					tokens = min(tokens, divide_and_conquer(mid+1, limit, machine))
				}()

				wg.Wait()
				break
			}
			tokens = min(tokens, (X_a/btnA[0])*3+mid)
			fmt.Println("outside", tokens)
			start = mid + 1
		} else {
			limit = mid - 1
		}
	}

	// fmt.Println(tokens)
	return tokens
}

func determinant(a1, b1, a2, b2 float64) float64 {
	return a1*b2 - a2*b1
}

// Function to solve the system using the matrix method
func solveMatrixMethod(a1, b1, c1, a2, b2, c2 float64) (float64, float64, error) {
	// Calculate determinant of matrix A
	det := determinant(a1, b1, a2, b2)

	if det == 0 {
		return 0, 0, errors.New("no unique solution (determinant is zero)")
	}

	// Calculate inverse of A multiplied by C
	x := (c1*b2 - c2*b1) / det
	y := (c1 - a1*x) / b1

	return x, y, nil
}

func main() {
	input := clear_input()

	totalTokens := 0
	// tokenChan := make(chan int, len(input))
	// var wg sync.WaitGroup
	for _, machine := range input {
		btnA := machine[0]
		btnB := machine[1]
		priceX := machine[2][0] + 1e13
		priceY := machine[2][1] + 1e13
		// fmt.Println(priceX, priceY)

		a, b, err := solveMatrixMethod(float64(btnA[0]), float64(btnB[0]), float64(priceX), float64(btnA[1]), float64(btnB[1]), float64(priceY))

		if err != nil {
			continue
		}

		if a != float64(int(a)) || b != float64(int(b)) {
			continue
		}

		// fmt.Println(a, b)
		totalTokens += int(a*3 + b)

		// var limit int
		// if priceX/btnB[0] >= priceY/btnB[1] {
		// 	limit = int(priceX/btnB[0]) + 1
		// } else {
		// 	limit = int(priceY/btnB[1]) + 1
		// }

		// wg.Add(1)
		// go func() {
		// 	defer wg.Done()
		// 	tokenChan <- divide_and_conquer(0, limit, machine)
		// }()
	}

	// go func() {
	// 	wg.Wait()
	// 	close(tokenChan)
	// }()

	// for tokens := range tokenChan {
	// 	if tokens != math.MaxInt {
	// 		totalTokens += tokens
	// 	}
	// }

	fmt.Println(totalTokens)

}
