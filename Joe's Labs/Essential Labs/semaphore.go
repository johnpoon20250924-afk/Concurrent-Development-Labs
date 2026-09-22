package main

import (
    "fmt"
    "sync"
    "time"
)
//make struct containing channel
//add init, acquire and release
type semaphore struct {
	theCounter chan struct{}
}
func Init(sem *semaphore, maxGoroutines int){
    sem.theCounter = make(chan struct{}, maxGoroutines)
}
func Acquite(sem *semaphore){
    sem.theCounter <- struct{}{}
}
func Release(sem *semaphore) {
    <-sem.theCounter
}
func main() {
    maxGoroutines := 5
    sembox := semaphore{}
    Init( &sembox , maxGoroutines)
    var wg sync.WaitGroup
    for i := 0; i < 20; i++ {
        wg.Add(1)
        go func(i int) {
            defer wg.Done()
            Acquite(&sembox)
            defer Release(&sembox)
            
            // Simulate a task
            fmt.Printf("Running task %d\n", i)
            time.Sleep(2 * time.Second)
        }(i)
    }
    wg.Wait()
}
