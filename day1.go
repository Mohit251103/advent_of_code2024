package main

import (
	"fmt"
	"sort"
)

func prob1() {

	var list1 []int
	var list2 []int

	mp := make(map[int]int)

	for i := 0; i < 1000; i++ {
		var x, y int
		fmt.Scanf("%d %d", &x, &y)
		list1 = append(list1, x)
		list2 = append(list2, y)
		mp[y]++
	}

	sort.Sort(sort.IntSlice(list1))
	sort.Sort(sort.IntSlice(list2))

	total := 0
	for i := 0; i < 1000; i++ {
		// total += int(math.Abs(float64(list1[i]) - float64(list2[i])))
		total += list1[i] * mp[list1[i]]
	}

	fmt.Println(total)

}
