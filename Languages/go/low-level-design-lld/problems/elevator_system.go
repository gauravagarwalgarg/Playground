package main

import (
	"fmt"
	"math"
	"sync"
)

// Elevator System LLD - Multi-elevator scheduling.
// Demonstrates: State pattern (elevator states), Strategy (scheduling algorithm).

type Direction int

const (
	Up   Direction = 1
	Down Direction = -1
	Idle Direction = 0
)

func (d Direction) String() string {
	switch d {
	case Up:
		return "UP"
	case Down:
		return "DOWN"
	default:
		return "IDLE"
	}
}

type ElevatorState int

const (
	StateIdle ElevatorState = iota
	StateMoving
	StateDoorOpen
)

type Elevator struct {
	mu           sync.Mutex
	ID           int
	CurrentFloor int
	Direction    Direction
	State        ElevatorState
	Requests     map[int]bool // floors to visit
	MinFloor     int
	MaxFloor     int
}

func NewElevator(id, minFloor, maxFloor int) *Elevator {
	return &Elevator{
		ID:           id,
		CurrentFloor: 1,
		Direction:    Idle,
		State:        StateIdle,
		Requests:     make(map[int]bool),
		MinFloor:     minFloor,
		MaxFloor:     maxFloor,
	}
}

func (e *Elevator) AddRequest(floor int) {
	e.mu.Lock()
	defer e.mu.Unlock()
	if floor >= e.MinFloor && floor <= e.MaxFloor {
		e.Requests[floor] = true
	}
}

func (e *Elevator) Step() {
	e.mu.Lock()
	defer e.mu.Unlock()

	if len(e.Requests) == 0 {
		e.Direction = Idle
		e.State = StateIdle
		return
	}

	// Check if current floor is a destination
	if e.Requests[e.CurrentFloor] {
		delete(e.Requests, e.CurrentFloor)
		e.State = StateDoorOpen
		return
	}

	// Determine direction
	e.State = StateMoving
	if e.Direction == Idle {
		// Pick the nearest request
		nearest := e.nearestRequest()
		if nearest > e.CurrentFloor {
			e.Direction = Up
		} else {
			e.Direction = Down
		}
	}

	// Move one floor
	e.CurrentFloor += int(e.Direction)

	// Reverse if no more requests in current direction
	if !e.hasRequestsInDirection() {
		e.Direction = Direction(-e.Direction)
	}
}

func (e *Elevator) nearestRequest() int {
	nearest := -1
	minDist := math.MaxInt
	for floor := range e.Requests {
		dist := int(math.Abs(float64(floor - e.CurrentFloor)))
		if dist < minDist {
			minDist = dist
			nearest = floor
		}
	}
	return nearest
}

func (e *Elevator) hasRequestsInDirection() bool {
	for floor := range e.Requests {
		if e.Direction == Up && floor > e.CurrentFloor {
			return true
		}
		if e.Direction == Down && floor < e.CurrentFloor {
			return true
		}
	}
	return false
}

func (e *Elevator) PendingRequests() int {
	e.mu.Lock()
	defer e.mu.Unlock()
	return len(e.Requests)
}

// ElevatorController - dispatches requests using nearest-elevator strategy
type ElevatorController struct {
	elevators []*Elevator
}

func NewElevatorController(numElevators, minFloor, maxFloor int) *ElevatorController {
	ctrl := &ElevatorController{}
	for i := 0; i < numElevators; i++ {
		ctrl.elevators = append(ctrl.elevators, NewElevator(i+1, minFloor, maxFloor))
	}
	return ctrl
}

func (c *ElevatorController) RequestElevator(fromFloor int, dir Direction) *Elevator {
	var best *Elevator
	bestScore := math.MaxInt

	for _, e := range c.elevators {
		e.mu.Lock()
		score := int(math.Abs(float64(e.CurrentFloor - fromFloor)))
		// Prefer idle elevators or those moving toward the request
		if e.State == StateIdle {
			score -= 5
		} else if (e.Direction == Up && fromFloor >= e.CurrentFloor && dir == Up) ||
			(e.Direction == Down && fromFloor <= e.CurrentFloor && dir == Down) {
			score -= 3
		}
		e.mu.Unlock()

		if score < bestScore {
			bestScore = score
			best = e
		}
	}

	if best != nil {
		best.AddRequest(fromFloor)
	}
	return best
}

func (c *ElevatorController) StepAll() {
	for _, e := range c.elevators {
		e.Step()
	}
}

func (c *ElevatorController) Status() {
	for _, e := range c.elevators {
		e.mu.Lock()
		fmt.Printf("  Elevator %d: Floor %d, Dir %s, Pending %d\n",
			e.ID, e.CurrentFloor, e.Direction, len(e.Requests))
		e.mu.Unlock()
	}
}

func main() {
	ctrl := NewElevatorController(3, 1, 10)

	// Person at floor 5 wants to go up
	e := ctrl.RequestElevator(5, Up)
	fmt.Printf("Assigned Elevator %d for floor 5 UP\n", e.ID)
	e.AddRequest(8) // destination: floor 8

	// Person at floor 2 wants to go up
	e2 := ctrl.RequestElevator(2, Up)
	fmt.Printf("Assigned Elevator %d for floor 2 UP\n", e2.ID)
	e2.AddRequest(6) // destination: floor 6

	// Simulate steps
	fmt.Println("\n--- Simulation ---")
	for step := 0; step < 15; step++ {
		ctrl.StepAll()
	}
	ctrl.Status()

	// Verify all requests served
	allServed := true
	for _, elev := range ctrl.elevators {
		if elev.PendingRequests() > 0 {
			allServed = false
		}
	}
	if allServed {
		fmt.Println("PASS: all elevator requests served")
	} else {
		panic("FAIL: pending requests remain")
	}

	// Test: multiple concurrent requests
	fmt.Println("\n--- Burst requests ---")
	for floor := 1; floor <= 10; floor++ {
		ctrl.RequestElevator(floor, Up)
	}
	for step := 0; step < 30; step++ {
		ctrl.StepAll()
	}
	ctrl.Status()
	fmt.Println("PASS: elevator system complete")
}
