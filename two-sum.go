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
	hashMap := make(map[int]int)
	var response []int

	for i, num := range nums {
		hashMap[num] = i
	}

	for i, num := range nums {
		y := target - num
		n, ok := hashMap[y]

		if ok && i != n {
			response = append(response, i)
			response = append(response, n)
			break
		}
	}

	return response
}
