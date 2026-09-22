// Barrier.go Template Code
// Copyright (C) 2024 Dr. Joseph Kehoe
//
// Original template provided by Dr. Joseph Kehoe.
// Modified lab work is distributed under the MIT License.
// See the LICENSE file for details.
//
//--------------------------------------------
// Author: Joseph Kehoe (Joseph.Kehoe@setu.ie)
// Created on 30/9/2024
// Modified by: Runge Wang (C00304516)
// Issues: None
//--------------------------------------------

// CLASSMATES I WORKED WITH:
// Dorian Nowacki (C00288613)

// CLASSMATES I HELPED:
// Sean McGrath
package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

// Place a barrier in this function --use Mutex's and Semaphores
func doStuff(goNum int, wg *sync.WaitGroup, counter *int,
	theLock *sync.Mutex, sem *semaphore.Weighted,
	totalRoutines int, ctx context.Context) bool {

	time.Sleep(time.Second)
	fmt.Println("Part A", goNum)

	theLock.Lock()
	(*counter)++

	if *counter == totalRoutines {
		// Last goroutine opens the barrier
		sem.Release(1)
		theLock.Unlock()
	} else {
		theLock.Unlock()

		// Wait until all goroutines finish Part A
		sem.Acquire(ctx, 1)

		// Allow the next goroutine to continue
		sem.Release(1)
	}

	fmt.Println("PartB", goNum)
	wg.Done()
	return true
}
func main() {
	counter := 0
	totalRoutines := 10

	var wg sync.WaitGroup
	wg.Add(totalRoutines)

	ctx := context.TODO()
	var theLock sync.Mutex

	// Create a semaphore with one permit
	sem := semaphore.NewWeighted(1)

	// Close the barrier before starting the goroutines
	sem.Acquire(ctx, 1)

	for i := range totalRoutines {
		go doStuff(i, &wg, &counter, &theLock, sem, totalRoutines, ctx)
	}

	wg.Wait()

	// Take back the final permit
	sem.Acquire(ctx, 1)
}
