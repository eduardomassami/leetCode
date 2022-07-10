package main

import "fmt"

func main() {
	result := romanToInt("MCMXCIV")
	fmt.Println(result)
}

var romans = map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}

func romanToInt(s string) int {
	var sum int

	for i, char := range s {
		current := romans[string(char)]

		if i+1 < len(s) {
			if current < romans[string(s[i+1])] {
				current = current * -1
			}
		}

		sum += current
	}

	return sum
}
