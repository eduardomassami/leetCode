package main

import "fmt"

func main() {

	fmt.Println(canConstruct("aa", "cb"))

}

func canConstruct(ransomNote string, magazine string) bool {

	if len(ransomNote) > len(magazine) {
		return false
	}

	set := make(map[rune]int)
	for _, char := range magazine {
		set[char]++
	}

	for _, char := range ransomNote {
		s, ok := set[char]
		if !ok || s <= 0 {
			return false
		}

		set[char]--
	}

	return true
}
