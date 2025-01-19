// Matching Brackets

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("A string with brackets: ")
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')

	pairs := map[rune]rune{
		']': '[',
		'}': '{',
		')': '(',
	}

	stack := []rune{}

	for _, char := range str {
		if char == '[' || char == '{' || char == '(' {
			stack = append(stack, char)
		} else if char == ']' || char == '}' || char == ')' {
			if len(stack) == 0 || pairs[char] != stack[len(stack)-1] {
				fmt.Println("Incorrect!")
				return
			}
			stack = stack[:len(stack)-1]
		}
	}

	if len(stack) == 0 {
		fmt.Println("Correct!")
	} else {
		fmt.Println("Incorrect!")
	}

}
