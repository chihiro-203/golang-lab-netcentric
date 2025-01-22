package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	students    int = 100
	maxStudents int = 30
	// readingHour map[int]int
)

func main() {
	startTime := time.Now().Second()
	rand.Seed(time.Now().Unix())

	var wg sync.WaitGroup
	librarySeats := make(chan int, maxStudents)

	for i := 1; i <= students; i++ {
		wg.Add(1)
		go func(studentID int) {
			defer wg.Done()
			readingHour := rand.Intn(4) + 1

			select {
			case librarySeats <- studentID:
				fmt.Printf("Time %d: Student %d starts reading at the library.\n", time.Now().Second()-startTime, studentID)
				time.Sleep(time.Duration(readingHour) * time.Second)
				fmt.Printf("Time %d: Student %d is leaving. Spent %d hour(s) reading.\n", time.Now().Second()-startTime, studentID, readingHour)
				<-librarySeats
			default:
				fmt.Printf("Time %d: Student %d is waiting.\n", time.Now().Second()-startTime, studentID)
			}
		}(i)
	}

	wg.Wait()
	fmt.Printf("Time %d: No more students. Let's call it a day.", time.Now().Second()-startTime)

}
