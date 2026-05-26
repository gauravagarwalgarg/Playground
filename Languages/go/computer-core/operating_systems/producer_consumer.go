// Producer-Consumer using Go channels
// Demonstrates goroutines and buffered channels as bounded buffer.
package main

import (
	"fmt"
	"sync"
)

func producer(ch chan<- int, id int, count int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < count; i++ {
		item := id*100 + i
		ch <- item
	}
}

func consumer(ch <-chan int, results *[]int, mu *sync.Mutex, count int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < count; i++ {
		item := <-ch
		mu.Lock()
		*results = append(*results, item)
		mu.Unlock()
	}
}

func main() {
	// Buffered channel acts as bounded buffer
	ch := make(chan int, 5)
	var mu sync.Mutex
	results := make([]int, 0)

	var prodWg, consWg sync.WaitGroup

	// 3 producers, each producing 5 items = 15 total
	for i := 0; i < 3; i++ {
		prodWg.Add(1)
		go producer(ch, i, 5, &prodWg)
	}

	// 3 consumers, each consuming 5 items = 15 total
	for i := 0; i < 3; i++ {
		consWg.Add(1)
		go consumer(ch, &results, &mu, 5, &consWg)
	}

	prodWg.Wait()
	consWg.Wait()

	if len(results) != 15 {
		panic(fmt.Sprintf("FAIL: expected 15 items, got %d", len(results)))
	}
	fmt.Println("PASS: all 15 items consumed")

	// Verify no duplicates
	seen := make(map[int]bool)
	for _, item := range results {
		if seen[item] {
			panic(fmt.Sprintf("FAIL: duplicate item %d", item))
		}
		seen[item] = true
	}
	fmt.Println("PASS: no duplicates")

	fmt.Println("All tests passed!")
}
