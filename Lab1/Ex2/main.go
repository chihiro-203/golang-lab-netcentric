package main

import (
	"fmt"
	"strings"
)

func main() {
	score := 0

	fmt.Print("Input a word: ")
	var word string
	fmt.Scan(&word)
	for _, char := range word {
		c := strings.ToUpper(string(char))

		switch c {
		case "A", "E", "I", "O", "U", "L", "N", "R", "S", "T":
			score++
		case "D", "G":
			score += 2
		case "B", "C", "M", "P":
			score += 3
		case "F", "H", "V", "W", "Y":
			score += 4
		case "K":
			score += 5
		case "J", "X":
			score += 8
		case "Q", "Z":
			score += 10
		default:
			score += 0
		}
	}

	fmt.Println("Scrabble Score:", score)
}
