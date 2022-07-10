package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(fizzBuzz(15))
}

func fizzBuzz(input int) []string {
	var response []string

	for i := 1; i <= input; i++ {
		curr := ""

		if i%3 == 0 {
			curr += "Fizz"
		}

		if i%5 == 0 {
			curr += "Buzz"
		}

		if curr == "" {
			curr += strconv.Itoa(i)
		}

		response = append(response, curr)
	}

	return response
}
