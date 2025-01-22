package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	var wg sync.WaitGroup
	cal := make(chan map[rune]int)

	wg.Add(1)
	go func(s string) {
		defer wg.Done()
		freq := map[rune]int{}
		for _, char := range str {
			freq[char]++
		}
		cal <- freq
	}(str)

	go func() {
		wg.Wait()
		close(cal)
	}()

	file, _ := os.Create("output.txt")
	defer file.Close()

	frequency := <-cal
	for char, count := range frequency {
		switch char {
		case ' ':
			fmt.Printf("'space': %d\n", count)
		case '\n':
			fmt.Printf("'newline': %d\n", count)
		case '\t':
			fmt.Printf("'tab': %d\n", count)
		case '\r':
			fmt.Printf("'carriage return': %d\n", count)
		default:
			fmt.Printf("%c: %d\n", char, count)
		}
	}
}
