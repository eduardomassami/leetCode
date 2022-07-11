package main

import "fmt"

func main() {

	fmt.Println(numberOfSteps(8))

	fmt.Println(numberOfStepsRecursion(8))

}

var counterRecursion int = 0

func numberOfStepsRecursion(num int) int {

	if num == 0 {
		return counterRecursion
	}

	counterRecursion++

	if num%2 == 0 {
		num = num / 2
	} else {
		num = num - 1
	}

	numberOfStepsRecursion(num)

	return counterRecursion
}

func numberOfSteps(num int) int {
	var counter int

	for num > 0 {
		if num%2 == 0 {
			num = num / 2
		} else {
			num = num - 1
		}
		counter++
	}

	return counter
}
