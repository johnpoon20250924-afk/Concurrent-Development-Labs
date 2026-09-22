//--------------------------------------------
// Author: Runge Wang (C00304516)
// Modified by: Runge Wang (C00304516)
// Issues: None
//--------------------------------------------

// CLASSMATES I WORKED WITH:
// Dorian Nowacki (C00288613)

// CLASSMATES I HELPED:
// Sean

// This program implements a rendezvous using Go channels.
// All goroutines must finish Part A before they can start Part B.
// --------------------------------------------
package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func WorkWithRendezvous(
	wg *sync.WaitGroup,
	Num int,
	arrived chan bool,
	release chan bool,
) bool {

	var X time.Duration
	X = time.Duration(rand.IntN(5))
	time.Sleep(X * time.Second)

	fmt.Println("Part A", Num)

	// Tell main: I have finished Part A
	arrived <- true

	// Wait here until main allows us to continue
	<-release

	fmt.Println("PartB", Num)

	wg.Done()
	return true
}

func main() {
	var wg sync.WaitGroup
	threadCount := 5

	// Channel used to tell main that a goroutine has arrived
	arrived := make(chan bool)

	// Channel used to release the waiting goroutines
	release := make(chan bool)

	wg.Add(threadCount)

	for N := range threadCount {
		go WorkWithRendezvous(&wg, N, arrived, release)
	}

	// Wait until all 5 goroutines have finished Part A
	// Wait until all goroutines finish Part A
	for i := 0; i < threadCount; i++ {
		<-arrived
	}

	// Release all goroutines to start Part B
	for i := 0; i < threadCount; i++ {
		release <- true
	}

	wg.Wait()
}
