// Goroutines
/*
* -> It is a way to run multiple functions and execute them concurrently
* -> Concurrency != Parallel execution
 */

package main

import (
	"fmt"
	"sync"
	"time"
)

var m = sync.RWMutex{}
var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func main() {
	t0 := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Printf("Total execution time: %v\n", time.Since(t0))
	// fmt.Printf("The results are %v\n", results)

}

func dbCall(i int) {
	// Simulating DB call delay
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	save(dbData[i])
	log()
	wg.Done()
}

func save(result string) {
	m.Lock()
	results = append(results, result)
	m.Unlock()
}

func log() {
	m.RLock()
	fmt.Printf("The current results are: %v\n", results)
	m.RUnlock()
}
