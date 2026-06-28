package main

import (
	"fmt"
	"sync"
	"time"
)

// Concurrency Patterns in Go
// Covers: Fan-Out/Fan-In, Worker Pool, Pipeline, Rate-Limited Workers,
//         Bounded Parallelism, Timeout/Cancellation.

// --- Fan-Out / Fan-In ---
// Fan-out: start multiple goroutines to read from the same channel.
// Fan-in: multiplex multiple input channels into a single output channel.
func fanIn(channels ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	for _, ch := range channels {
		wg.Add(1)
		go func(c <-chan int) {
			defer wg.Done()
			for val := range c {
				out <- val
			}
		}(ch)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func producer(id, count int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 0; i < count; i++ {
			out <- id*100 + i
		}
	}()
	return out
}

// --- Worker Pool ---
type Job struct {
	ID      int
	Payload int
}

type Result struct {
	JobID  int
	Output int
}

func workerPool(numWorkers int, jobs <-chan Job) <-chan Result {
	results := make(chan Result, len(jobs))
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for job := range jobs {
				// Simulate work: square the payload
				output := job.Payload * job.Payload
				results <- Result{JobID: job.ID, Output: output}
			}
		}(i)
	}

	go func() {
		wg.Wait()
		close(results)
	}()
	return results
}

// --- Pipeline ---
// Each stage is a goroutine that processes and forwards data.
func generate(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * n
		}
	}()
	return out
}

func filter(in <-chan int, predicate func(int) bool) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			if predicate(n) {
				out <- n
			}
		}
	}()
	return out
}

// --- Bounded Parallelism ---
func boundedParallel(tasks []func() int, maxConcurrency int) []int {
	sem := make(chan struct{}, maxConcurrency)
	var mu sync.Mutex
	results := make([]int, len(tasks))
	var wg sync.WaitGroup

	for i, task := range tasks {
		wg.Add(1)
		sem <- struct{}{} // Acquire semaphore
		go func(idx int, fn func() int) {
			defer wg.Done()
			defer func() { <-sem }() // Release semaphore
			result := fn()
			mu.Lock()
			results[idx] = result
			mu.Unlock()
		}(i, task)
	}
	wg.Wait()
	return results
}

// --- Timeout Pattern ---
func withTimeout(fn func() int, timeout time.Duration) (int, error) {
	ch := make(chan int, 1)
	go func() {
		ch <- fn()
	}()
	select {
	case result := <-ch:
		return result, nil
	case <-time.After(timeout):
		return 0, fmt.Errorf("timeout after %v", timeout)
	}
}

func main() {
	// Test Fan-Out / Fan-In
	fmt.Println("--- Fan-Out / Fan-In ---")
	p1 := producer(1, 5)
	p2 := producer(2, 5)
	p3 := producer(3, 5)
	merged := fanIn(p1, p2, p3)

	var allVals []int
	for val := range merged {
		allVals = append(allVals, val)
	}
	if len(allVals) != 15 {
		panic(fmt.Sprintf("FAIL: expected 15 values, got %d", len(allVals)))
	}
	fmt.Printf("PASS: Fan-In collected %d values from 3 producers\n", len(allVals))

	// Test Worker Pool
	fmt.Println("\n--- Worker Pool ---")
	jobs := make(chan Job, 20)
	for i := 0; i < 20; i++ {
		jobs <- Job{ID: i, Payload: i + 1}
	}
	close(jobs)

	results := workerPool(4, jobs)
	var poolResults []Result
	for r := range results {
		poolResults = append(poolResults, r)
	}
	if len(poolResults) != 20 {
		panic(fmt.Sprintf("FAIL: expected 20 results, got %d", len(poolResults)))
	}
	// Verify a result
	found := false
	for _, r := range poolResults {
		if r.JobID == 4 && r.Output == 25 { // (4+1)^2 = 25
			found = true
		}
	}
	if !found {
		panic("FAIL: expected job 4 result = 25")
	}
	fmt.Printf("PASS: Worker Pool processed %d jobs\n", len(poolResults))

	// Test Pipeline
	fmt.Println("\n--- Pipeline ---")
	// generate → square → filter (only even)
	nums := generate(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	squared := square(nums)
	evens := filter(squared, func(n int) bool { return n%2 == 0 })

	var pipelineResults []int
	for n := range evens {
		pipelineResults = append(pipelineResults, n)
	}
	// Even squares: 4, 16, 36, 64, 100
	if len(pipelineResults) != 5 {
		panic(fmt.Sprintf("FAIL: expected 5 even squares, got %d: %v", len(pipelineResults), pipelineResults))
	}
	fmt.Printf("PASS: Pipeline produced %d even squares: %v\n", len(pipelineResults), pipelineResults)

	// Test Bounded Parallelism
	fmt.Println("\n--- Bounded Parallelism ---")
	tasks := make([]func() int, 100)
	for i := range tasks {
		val := i
		tasks[i] = func() int {
			time.Sleep(1 * time.Millisecond)
			return val * 2
		}
	}
	start := time.Now()
	bpResults := boundedParallel(tasks, 10) // max 10 concurrent
	elapsed := time.Since(start)
	if len(bpResults) != 100 {
		panic("FAIL: bounded parallel")
	}
	if bpResults[5] != 10 { // 5*2 = 10
		panic(fmt.Sprintf("FAIL: expected 10, got %d", bpResults[5]))
	}
	fmt.Printf("PASS: 100 tasks with max 10 concurrent in %v\n", elapsed)

	// Test Timeout
	fmt.Println("\n--- Timeout Pattern ---")
	// Fast function succeeds
	result, err := withTimeout(func() int {
		return 42
	}, 100*time.Millisecond)
	if err != nil || result != 42 {
		panic("FAIL: fast function should succeed")
	}
	fmt.Println("PASS: fast function returned before timeout")

	// Slow function times out
	_, err = withTimeout(func() int {
		time.Sleep(200 * time.Millisecond)
		return 0
	}, 50*time.Millisecond)
	if err == nil {
		panic("FAIL: slow function should timeout")
	}
	fmt.Println("PASS: slow function timed out:", err)

	fmt.Println("\nPASS: all concurrency patterns complete")
}
