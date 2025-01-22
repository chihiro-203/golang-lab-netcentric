package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	students int = 100
	maxSeats int = 30
)

func main() {
	startTime := time.Now()
	rand.Seed(time.Now().Unix())
	librarySeats := make(chan int, maxSeats)

	var wg sync.WaitGroup

	for i := 1; i <= students; i++ {
		wg.Add(1)
		go func(studentID int) {
			defer wg.Done()
			readingHour := rand.Intn(4) + 1

			for {
				select {
				case librarySeats <- studentID:
					elapsedTime := int(time.Since(startTime).Seconds())
					fmt.Printf("Time %d: Student %d starts reading at the library\n", elapsedTime, studentID)
					time.Sleep(time.Duration(readingHour) * time.Second)
					elapsedTime = int(time.Since(startTime).Seconds())
					fmt.Printf("Time %d: Student %d is leaving. Spent %d hours reading\n", elapsedTime, studentID, readingHour)
					<-librarySeats
					return
				default:
					elapsedTime := int(time.Since(startTime).Seconds())
					fmt.Printf("Time %d: Student %d is waiting\n", elapsedTime, studentID)
					time.Sleep(1 * time.Second)
				}

			}
		}(i)
	}

	wg.Wait()
	elapsedTime := int(time.Since(startTime).Seconds())
	fmt.Printf("Time %d: No more students. Let's call it a day.", elapsedTime)
}
