package main

import "fmt"

func Question_1() {

	vst := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 9, 8, 6, 5, 4, 3, 2, 1}

	countMap := make(map[int]int)
	for _, v := range vst {
		countMap[v]++
	}
	fmt.Println(countMap)
	for k, v := range countMap {
		if v == 1 {
			fmt.Println(k)
		}
	}

}
