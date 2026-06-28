package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// Circuit Breaker Pattern - Prevents cascading failures in distributed systems.
// States: CLOSED (normal) → OPEN (failing) → HALF_OPEN (testing recovery)
// Used in: microservices, API gateways, HTTP clients.

type CircuitState int

const (
	Closed   CircuitState = iota // Normal operation
	Open                         // Failing, reject requests
	HalfOpen                     // Testing if service recovered
)

func (s CircuitState) String() string {
	return [...]string{"CLOSED", "OPEN", "HALF_OPEN"}[s]
}

type CircuitBreaker struct {
	mu              sync.Mutex
	state           CircuitState
	failureCount    int
	successCount    int
	failureThreshold int
	successThreshold int // successes needed in half-open to close
	timeout          time.Duration
	lastFailureTime time.Time
	onStateChange   func(from, to CircuitState)
}

type CircuitBreakerConfig struct {
	FailureThreshold int
	SuccessThreshold int
	Timeout          time.Duration
	OnStateChange    func(from, to CircuitState)
}

func NewCircuitBreaker(cfg CircuitBreakerConfig) *CircuitBreaker {
	return &CircuitBreaker{
		state:            Closed,
		failureThreshold: cfg.FailureThreshold,
		successThreshold: cfg.SuccessThreshold,
		timeout:          cfg.Timeout,
		onStateChange:    cfg.OnStateChange,
	}
}

func (cb *CircuitBreaker) Execute(fn func() error) error {
	cb.mu.Lock()
	if !cb.canExecute() {
		cb.mu.Unlock()
		return errors.New("circuit breaker is OPEN")
	}
	cb.mu.Unlock()

	// Execute the function
	err := fn()

	cb.mu.Lock()
	defer cb.mu.Unlock()
	if err != nil {
		cb.onFailure()
		return err
	}
	cb.onSuccess()
	return nil
}

func (cb *CircuitBreaker) canExecute() bool {
	switch cb.state {
	case Closed:
		return true
	case Open:
		// Check if timeout has elapsed → transition to half-open
		if time.Since(cb.lastFailureTime) >= cb.timeout {
			cb.transition(HalfOpen)
			return true
		}
		return false
	case HalfOpen:
		return true
	}
	return false
}

func (cb *CircuitBreaker) onSuccess() {
	switch cb.state {
	case Closed:
		cb.failureCount = 0
	case HalfOpen:
		cb.successCount++
		if cb.successCount >= cb.successThreshold {
			cb.transition(Closed)
		}
	}
}

func (cb *CircuitBreaker) onFailure() {
	switch cb.state {
	case Closed:
		cb.failureCount++
		if cb.failureCount >= cb.failureThreshold {
			cb.transition(Open)
		}
	case HalfOpen:
		// Any failure in half-open → back to open
		cb.transition(Open)
	}
}

func (cb *CircuitBreaker) transition(to CircuitState) {
	from := cb.state
	cb.state = to
	cb.failureCount = 0
	cb.successCount = 0
	if to == Open {
		cb.lastFailureTime = time.Now()
	}
	if cb.onStateChange != nil {
		cb.onStateChange(from, to)
	}
}

func (cb *CircuitBreaker) State() CircuitState {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	return cb.state
}

// --- Simulated External Service ---
type ExternalService struct {
	mu         sync.Mutex
	healthy    bool
	callCount  int
}

func NewExternalService(healthy bool) *ExternalService {
	return &ExternalService{healthy: healthy}
}

func (s *ExternalService) Call() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.callCount++
	if !s.healthy {
		return errors.New("service unavailable")
	}
	return nil
}

func (s *ExternalService) SetHealthy(h bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.healthy = h
}

func (s *ExternalService) CallCount() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.callCount
}

func main() {
	transitions := []string{}

	cb := NewCircuitBreaker(CircuitBreakerConfig{
		FailureThreshold: 3,
		SuccessThreshold: 2,
		Timeout:          100 * time.Millisecond,
		OnStateChange: func(from, to CircuitState) {
			msg := fmt.Sprintf("%s → %s", from, to)
			transitions = append(transitions, msg)
			fmt.Printf("  [CB] State change: %s\n", msg)
		},
	})

	service := NewExternalService(true)

	// Phase 1: Normal operation (CLOSED)
	fmt.Println("--- Phase 1: Normal operation ---")
	for i := 0; i < 5; i++ {
		err := cb.Execute(service.Call)
		if err != nil {
			panic("FAIL: should succeed when service is healthy")
		}
	}
	if cb.State() != Closed {
		panic("FAIL: should remain CLOSED")
	}
	fmt.Println("PASS: healthy calls succeed in CLOSED state")

	// Phase 2: Service starts failing → trips to OPEN
	fmt.Println("\n--- Phase 2: Service fails → OPEN ---")
	service.SetHealthy(false)
	for i := 0; i < 3; i++ {
		cb.Execute(service.Call)
	}
	if cb.State() != Open {
		panic("FAIL: should be OPEN after 3 failures")
	}
	fmt.Println("PASS: circuit opened after threshold failures")

	// Phase 3: Requests rejected in OPEN state
	fmt.Println("\n--- Phase 3: Requests rejected in OPEN ---")
	callsBefore := service.CallCount()
	err := cb.Execute(service.Call)
	if err == nil {
		panic("FAIL: should reject in OPEN state")
	}
	if service.CallCount() != callsBefore {
		panic("FAIL: should not call service when OPEN")
	}
	fmt.Println("PASS: requests rejected, service not called")

	// Phase 4: After timeout → HALF_OPEN, test with success
	fmt.Println("\n--- Phase 4: Timeout → HALF_OPEN → CLOSED ---")
	service.SetHealthy(true)
	time.Sleep(150 * time.Millisecond) // wait for timeout

	// First call in half-open
	err = cb.Execute(service.Call)
	if err != nil {
		panic("FAIL: should succeed in HALF_OPEN")
	}
	if cb.State() != HalfOpen {
		panic("FAIL: should be in HALF_OPEN")
	}

	// Second success → close circuit
	err = cb.Execute(service.Call)
	if err != nil {
		panic("FAIL: second call should succeed")
	}
	if cb.State() != Closed {
		panic("FAIL: should be CLOSED after success threshold")
	}
	fmt.Println("PASS: circuit recovered (HALF_OPEN → CLOSED)")

	// Phase 5: Fail again in HALF_OPEN → back to OPEN
	fmt.Println("\n--- Phase 5: Fail in HALF_OPEN → back to OPEN ---")
	service.SetHealthy(false)
	// Trip to OPEN
	for i := 0; i < 3; i++ {
		cb.Execute(service.Call)
	}
	// Wait for timeout → HALF_OPEN
	time.Sleep(150 * time.Millisecond)
	// One failure in HALF_OPEN → back to OPEN
	cb.Execute(service.Call)
	if cb.State() != Open {
		panic("FAIL: should be OPEN after failure in HALF_OPEN")
	}
	fmt.Println("PASS: failure in HALF_OPEN sends back to OPEN")

	// Summary
	fmt.Printf("\n--- State transitions: %v ---\n", transitions)
	fmt.Println("PASS: circuit breaker complete")
}
