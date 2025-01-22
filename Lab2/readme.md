# Lab 2 - Go Concurrency

This repository contains the lab exercises for Net-Centric Programming, focusing on concurrency in the Go programming language. Below are the objectives, sample code snippets, and exercises included in this lab.

## Objective
- Learn and get familiar with concurrency in Go using goroutines and channels.

---

## Sample Codes using Goroutines

### Example 1: Goroutines Writing to a Shared Channel
This program initializes 10 goroutines, each writing to a shared channel.

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	ch := make(chan string)

	for i := 0; i < 10; i++ {
		go func(i int) {
			for j := 0; j < 10; j++ {
				ch <- "Goroutine : " + strconv.Itoa(i)
			}
		}(i)
	}

	for k := 1; k <= 100; k++ {
		fmt.Println(k, <-ch)
	}
}
```

---

### Example 2: Odd and Even Numbers with Goroutines
This program initializes 2 goroutines to check odd and even numbers from a given series of integer.

```go
package main

import (
	"fmt"
)

func main() {
	var intSlice = []int{42, 17, 89, 74, 5, 63, 38, 49, 92, 29}
	chOdd := make(chan int)
	chEven := make(chan int)

	go odd(chOdd)
	go even(chEven)

	for _, value := range intSlice {
		if value%2 != 0 {
			chOdd <- value
		} else {
			chEven <- value
		}
	}
}

func odd(ch <-chan int) {
	for v := range ch {
		fmt.Println("ODD :", v)
	}
}

func even(ch <-chan int) {
	for v := range ch {
		fmt.Println("EVEN:", v)
	}
}
```

---

### Example 3: Synchronizing Goroutines with WaitGroup
This program uses `sync.WaitGroup` to synchronize goroutines.

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	// Add the number of goroutines to wait for
	wg.Add(3)

	go performTask("Task 1", &wg)
	go performTask("Task 2", &wg)
	go performTask("Task 3", &wg)

	// Wait for all goroutines to complete
	wg.Wait()

	fmt.Println("All tasks completed.")
}

func performTask(taskName string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Starting %s\n", taskName)
	time.Sleep(2 * time.Second) // Simulate some work
	fmt.Printf("Completed %s\n", taskName)
}
```

---

## Exercises

### 1. Character Frequency ([Solution](https://github.com/ume-meu/golang-lab-netcentric/blob/main/Lab2/Ex1/main.go))
**Problem:**  
Find the frequency of each character in a string. Improve the code using concurrency.  
**Optional Task:** Perform the task on a text file.

- **Example Input:** `Net-centric Programming`
- **Example Output:**
  ```
  n: 3
  e: 2
  t: 2
  -: 1
  r: 3
  i: 2
  c: 2
  (space): 1
  p: 1
  g: 2
  a: 1
  m: 2
  ```

---

### 2. Library Simulation ([Solution](https://github.com/ume-meu/golang-lab-netcentric/blob/main/Lab2/Ex2/main.go))
**Problem:**  
Simulate the operation of a library using Go concurrency.

**Details:** Given the case study of International University library:
- The library can seat up to 30 students at a given time
- Usually, there is around 100 students vist everyday
- For each visit, students spend from 1 to 4 hours reading at the lib

**Task:** Write a go program to simulate the library operation in one day:
- Randomly generate and assign “reader ID” to student that come to the lib.
- Use sleep to simulate the time (1 second = 1 hour in real life)
- First come first serve. If the lib is full, student need to wait
- After finish studying at the lib, student will leave
- How many hours the lib need to open to serve all the students

**Example Output:**
```
Time 0: Student 3 starts reading at the library
Time 0: Student 15 starts reading at the library
Time 0: Student 2 starts reading at the library
...
Time 1: Student 45 is waiting
...
Time 5: Student 5 is leaving. Spent 3 hours reading
...
Time 10: No more students. Let's call it a day.
```

---

## Notes
These exercises are designed to deepen your understanding of Go concurrency and its applications. Feel free to contribute by submitting your solutions or enhancements!
```