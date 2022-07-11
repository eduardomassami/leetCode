package main

import "fmt"

func main() {
	// accounts := [][]int{
	// 	{1, 2, 3},
	// 	{3, 2, 1},
	// }

	// fmt.Println(maximumWealth(accounts))

	accounts := [][]int{
		{1, 5},
		{7, 3},
		{3, 5},
	}

	fmt.Println(maximumWealth(accounts))
}

func maximumWealth(accounts [][]int) int {

	var customersWealth int

	for _, customer := range accounts {
		var total int

		for _, wealth := range customer {
			total += wealth
		}

		if total > customersWealth {
			customersWealth = total
		}
	}

	return customersWealth
}
