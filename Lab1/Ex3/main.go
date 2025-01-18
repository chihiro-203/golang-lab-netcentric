// Luhn Algorithm

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader((os.Stdin))
	fmt.Print("Input number: ")
	str, _ := reader.ReadString('\n')

	str = strings.ReplaceAll(str, " ", "")
	str = strings.TrimSpace(str)

	var sum int

	if len(str) <= 1 {
		fmt.Println("Invalid String")
	} else {
		for i, v := range str {
			val, _ := strconv.Atoi(string(v))
			cal := val
			if i%2 == 0 {
				cal = val * 2
				if cal > 9 {
					cal -= 9
				}
			}
			sum += cal
		}
	}
	if sum%10 == 0 {
		fmt.Println("Sum:", sum, "Divisible by 10\nValid")
	} else {
		fmt.Println("Sum:", sum, "Not divisible by 10\nInvalid")
	}
}
