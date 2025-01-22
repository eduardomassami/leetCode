package main

import "fmt"

func main() {

	input := []int{2, 7, 11, 15}
	fmt.Println(twoSum(input, 9))

	input2 := []int{3, 2, 4}
	fmt.Println(twoSum(input2, 6))

	input3 := []int{3, 3}
	fmt.Println(twoSum(input3, 6))
}

func twoSum(nums []int, target int) []int {
    	mapNums := map[int]int{}

	if len(nums) < 3 {
		return []int{0, 1}
	}

	for i, val := range nums {
		c := target - val
		if idx, ok := mapNums[c]; ok {
			return []int{idx, i}
		}
		mapNums[val] = i
	}

	return []int{}
}
