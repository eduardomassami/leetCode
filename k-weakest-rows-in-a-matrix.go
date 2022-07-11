package main

import (
	"fmt"
	"sort"
)

func main() {

	// mat := [][]int{
	// 	{1, 1, 0, 0, 0},
	// 	{1, 1, 1, 1, 0},
	// 	{1, 0, 0, 0, 0},
	// 	{1, 1, 0, 0, 0},
	// 	{1, 1, 1, 1, 1},
	// }
	// response := kWeakestRows(mat, 3)

	mat := [][]int{
		{1, 1, 0},
		{1, 1, 0},
		{1, 1, 1},
		{1, 1, 1},
		{0, 0, 0},
		{1, 1, 1},
		{1, 0, 0},
	}
	response := kWeakestRows(mat, 6)

	fmt.Println(response)
}

func kWeakestRows(mat [][]int, k int) []int {
	var response []int
	mapSoldiers := make(map[int]int)

	for i, row := range mat {
		mapSoldiers[i] = 0
		response = append(response, i)
		for _, soldiers := range row {
			if soldiers != 0 {
				mapSoldiers[i]++
				continue
			}
			break
		}
	}

	sort.SliceStable(response, func(i, j int) bool {
		return mapSoldiers[response[i]] < mapSoldiers[response[j]]
	})

	return response[0:k]
}
