// Hamming Distance

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	dis := 0
	set := "CAGT"
	for i := 0; i < 1000; i++ {
		dna1 := set[rand.Intn(len(set))]
		dna2 := set[rand.Intn(len(set))]

		if dna1 == dna2 {
			dis++
		}
	}
	fmt.Println("Hamming Distance:", dis)
}
