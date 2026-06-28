package main

import (
	"container/heap"
	"fmt"
	"sync"
	"time"
)

// Task Scheduler / Job Scheduler LLD
// Demonstrates: Priority Queue (heap), Worker Pool, Producer-Consumer pattern.
// Common in: Distributed systems, cron jobs, background processing.

type Priority int

const (
	Low    Priority = 0
	Medium Priority = 1
	High   Priority = 2
)

func (p Priority) String() string {
	return [...]string{"LOW", "MEDIUM", "HIGH"}[p]
}

type TaskStatus int

const (
	Pending TaskStatus = iota
	Running
	Completed
	Failed
)

func (s TaskStatus) String() string {
	return [...]string{"PENDING", "RUNNING", "COMPLETED", "FAILED"}[s]
}

type Task struct {
	ID        string
	Name      string
	Priority  Priority
	Status    TaskStatus
	CreatedAt time.Time
	Handler   func() error
	index     int // for heap interface
}

// TaskQueue - priority queue backed by a heap
type TaskQueue []*Task

func (pq TaskQueue) Len() int { return len(pq) }
func (pq TaskQueue) Less(i, j int) bool {
	// Higher priority first; if equal, earlier created first
	if pq[i].Priority != pq[j].Priority {
		return pq[i].Priority > pq[j].Priority
	}
	return pq[i].CreatedAt.Before(pq[j].CreatedAt)
}
func (pq TaskQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *TaskQueue) Push(x interface{}) {
	n := len(*pq)
	task := x.(*Task)
	task.index = n
	*pq = append(*pq, task)
}
func (pq *TaskQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	task := old[n-1]
	old[n-1] = nil
	task.index = -1
	*pq = old[:n-1]
	return task
}

// Scheduler - manages task execution with worker pool
type Scheduler struct {
	mu          sync.Mutex
	queue       TaskQueue
	workers     int
	results     map[string]TaskStatus
	taskChan    chan *Task
	wg          sync.WaitGroup
	running     bool
	completedMu sync.Mutex
	completed   int
}

func NewScheduler(workers int) *Scheduler {
	s := &Scheduler{
		workers:  workers,
		results:  make(map[string]TaskStatus),
		taskChan: make(chan *Task, 100),
	}
	heap.Init(&s.queue)
	return s
}

func (s *Scheduler) Submit(task *Task) {
	s.mu.Lock()
	defer s.mu.Unlock()
	task.Status = Pending
	task.CreatedAt = time.Now()
	heap.Push(&s.queue, task)
	s.results[task.ID] = Pending
}

func (s *Scheduler) Start() {
	s.running = true
	// Start workers
	for i := 0; i < s.workers; i++ {
		s.wg.Add(1)
		go s.worker(i)
	}
	// Dispatcher: move tasks from priority queue to worker channel
	go s.dispatch()
}

func (s *Scheduler) dispatch() {
	for s.running {
		s.mu.Lock()
		if s.queue.Len() > 0 {
			task := heap.Pop(&s.queue).(*Task)
			s.mu.Unlock()
			s.taskChan <- task
		} else {
			s.mu.Unlock()
			time.Sleep(1 * time.Millisecond)
		}
	}
}

func (s *Scheduler) worker(id int) {
	defer s.wg.Done()
	for task := range s.taskChan {
		task.Status = Running
		s.mu.Lock()
		s.results[task.ID] = Running
		s.mu.Unlock()

		// Execute the task
		var status TaskStatus
		if task.Handler != nil {
			err := task.Handler()
			if err != nil {
				status = Failed
			} else {
				status = Completed
			}
		} else {
			status = Completed
		}

		task.Status = status
		s.mu.Lock()
		s.results[task.ID] = status
		s.mu.Unlock()

		s.completedMu.Lock()
		s.completed++
		s.completedMu.Unlock()
	}
}

func (s *Scheduler) Stop() {
	s.running = false
	close(s.taskChan)
	s.wg.Wait()
}

func (s *Scheduler) GetStatus(taskID string) TaskStatus {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.results[taskID]
}

func (s *Scheduler) CompletedCount() int {
	s.completedMu.Lock()
	defer s.completedMu.Unlock()
	return s.completed
}

func (s *Scheduler) PendingCount() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.queue.Len()
}

func main() {
	scheduler := NewScheduler(3) // 3 worker goroutines

	// Submit tasks with different priorities
	tasks := []*Task{
		{ID: "t1", Name: "Send email", Priority: Low, Handler: func() error {
			time.Sleep(5 * time.Millisecond)
			return nil
		}},
		{ID: "t2", Name: "Process payment", Priority: High, Handler: func() error {
			time.Sleep(5 * time.Millisecond)
			return nil
		}},
		{ID: "t3", Name: "Generate report", Priority: Medium, Handler: func() error {
			time.Sleep(5 * time.Millisecond)
			return nil
		}},
		{ID: "t4", Name: "Resize image", Priority: Low, Handler: func() error {
			time.Sleep(5 * time.Millisecond)
			return nil
		}},
		{ID: "t5", Name: "Fraud check", Priority: High, Handler: func() error {
			time.Sleep(5 * time.Millisecond)
			return nil
		}},
		{ID: "t6", Name: "Failing task", Priority: Medium, Handler: func() error {
			return fmt.Errorf("intentional failure")
		}},
	}

	for _, t := range tasks {
		scheduler.Submit(t)
	}
	fmt.Printf("Submitted %d tasks\n", len(tasks))
	fmt.Printf("Pending: %d\n", scheduler.PendingCount())

	// Verify priority ordering in the queue
	scheduler.mu.Lock()
	if scheduler.queue.Len() > 0 {
		top := scheduler.queue[0]
		if top.Priority != High {
			panic(fmt.Sprintf("FAIL: expected HIGH priority first, got %s", top.Priority))
		}
		fmt.Printf("PASS: highest priority task '%s' at front of queue\n", top.Name)
	}
	scheduler.mu.Unlock()

	// Start processing
	scheduler.Start()

	// Wait for all tasks to complete
	time.Sleep(200 * time.Millisecond)
	scheduler.Stop()

	// Check results
	fmt.Printf("\n--- Results ---\n")
	for _, t := range tasks {
		status := scheduler.GetStatus(t.ID)
		fmt.Printf("  %s (%s): %s\n", t.Name, t.Priority, status)
	}

	// Verify completion
	if scheduler.CompletedCount() != 6 {
		panic(fmt.Sprintf("FAIL: expected 6 completed, got %d", scheduler.CompletedCount()))
	}
	fmt.Println("\nPASS: all tasks processed")

	// Verify failed task
	if scheduler.GetStatus("t6") != Failed {
		panic("FAIL: t6 should have failed")
	}
	fmt.Println("PASS: failed task detected")

	// Verify successful tasks
	for _, id := range []string{"t1", "t2", "t3", "t4", "t5"} {
		if scheduler.GetStatus(id) != Completed {
			panic(fmt.Sprintf("FAIL: %s should be completed", id))
		}
	}
	fmt.Println("PASS: successful tasks completed")
}
