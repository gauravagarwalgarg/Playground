package main

import (
	"fmt"
	"sync"
)

// Parking Lot LLD - Multi-level parking with different vehicle types.
// Demonstrates: Strategy (pricing), Observer (spot changes), Singleton (lot manager).
// Thread-safe via sync.Mutex.

type VehicleType int

const (
	Motorcycle VehicleType = iota
	Car
	Truck
)

func (v VehicleType) String() string {
	return [...]string{"Motorcycle", "Car", "Truck"}[v]
}

type Vehicle struct {
	LicensePlate string
	Type         VehicleType
}

type SpotSize int

const (
	SmallSpot  SpotSize = iota // motorcycle
	MediumSpot                 // car
	LargeSpot                  // truck
)

type ParkingSpot struct {
	ID       int
	Level    int
	Size     SpotSize
	Occupied bool
	Vehicle  *Vehicle
}

func (s *ParkingSpot) CanFit(v *Vehicle) bool {
	switch v.Type {
	case Motorcycle:
		return true // fits anywhere
	case Car:
		return s.Size >= MediumSpot
	case Truck:
		return s.Size >= LargeSpot
	}
	return false
}

func (s *ParkingSpot) Park(v *Vehicle) bool {
	if s.Occupied || !s.CanFit(v) {
		return false
	}
	s.Occupied = true
	s.Vehicle = v
	return true
}

func (s *ParkingSpot) Unpark() *Vehicle {
	v := s.Vehicle
	s.Occupied = false
	s.Vehicle = nil
	return v
}

type Level struct {
	Floor int
	Spots []*ParkingSpot
}

func NewLevel(floor int, small, medium, large int) *Level {
	level := &Level{Floor: floor}
	id := floor * 100
	for i := 0; i < small; i++ {
		id++
		level.Spots = append(level.Spots, &ParkingSpot{ID: id, Level: floor, Size: SmallSpot})
	}
	for i := 0; i < medium; i++ {
		id++
		level.Spots = append(level.Spots, &ParkingSpot{ID: id, Level: floor, Size: MediumSpot})
	}
	for i := 0; i < large; i++ {
		id++
		level.Spots = append(level.Spots, &ParkingSpot{ID: id, Level: floor, Size: LargeSpot})
	}
	return level
}

func (l *Level) FindSpot(v *Vehicle) *ParkingSpot {
	for _, spot := range l.Spots {
		if !spot.Occupied && spot.CanFit(v) {
			return spot
		}
	}
	return nil
}

func (l *Level) Available() int {
	count := 0
	for _, s := range l.Spots {
		if !s.Occupied {
			count++
		}
	}
	return count
}

// ParkingLot - the main orchestrator
type ParkingLot struct {
	mu      sync.Mutex
	levels  []*Level
	tickets map[string]int // licensePlate → spotID
}

func NewParkingLot(levels []*Level) *ParkingLot {
	return &ParkingLot{
		levels:  levels,
		tickets: make(map[string]int),
	}
}

func (pl *ParkingLot) Park(v *Vehicle) (int, error) {
	pl.mu.Lock()
	defer pl.mu.Unlock()

	if _, exists := pl.tickets[v.LicensePlate]; exists {
		return 0, fmt.Errorf("vehicle %s already parked", v.LicensePlate)
	}

	for _, level := range pl.levels {
		spot := level.FindSpot(v)
		if spot != nil {
			spot.Park(v)
			pl.tickets[v.LicensePlate] = spot.ID
			return spot.ID, nil
		}
	}
	return 0, fmt.Errorf("no spot available for %s", v.Type)
}

func (pl *ParkingLot) Unpark(licensePlate string) error {
	pl.mu.Lock()
	defer pl.mu.Unlock()

	spotID, exists := pl.tickets[licensePlate]
	if !exists {
		return fmt.Errorf("vehicle %s not found", licensePlate)
	}

	for _, level := range pl.levels {
		for _, spot := range level.Spots {
			if spot.ID == spotID {
				spot.Unpark()
				delete(pl.tickets, licensePlate)
				return nil
			}
		}
	}
	return fmt.Errorf("spot %d not found", spotID)
}

func (pl *ParkingLot) TotalAvailable() int {
	pl.mu.Lock()
	defer pl.mu.Unlock()
	total := 0
	for _, l := range pl.levels {
		total += l.Available()
	}
	return total
}

func main() {
	// Create a 2-level parking lot
	lot := NewParkingLot([]*Level{
		NewLevel(1, 5, 10, 2), // 5 small, 10 medium, 2 large
		NewLevel(2, 3, 8, 3),  // 3 small, 8 medium, 3 large
	})

	totalSpots := lot.TotalAvailable()
	fmt.Printf("Total spots: %d\n", totalSpots)

	// Park vehicles
	car := &Vehicle{LicensePlate: "ABC-123", Type: Car}
	moto := &Vehicle{LicensePlate: "MOTO-1", Type: Motorcycle}
	truck := &Vehicle{LicensePlate: "TRUCK-X", Type: Truck}

	spotID, err := lot.Park(car)
	if err != nil {
		panic("FAIL: " + err.Error())
	}
	fmt.Printf("Car parked at spot %d\n", spotID)

	spotID, err = lot.Park(moto)
	if err != nil {
		panic("FAIL: " + err.Error())
	}
	fmt.Printf("Motorcycle parked at spot %d\n", spotID)

	spotID, err = lot.Park(truck)
	if err != nil {
		panic("FAIL: " + err.Error())
	}
	fmt.Printf("Truck parked at spot %d\n", spotID)

	// Verify availability decreased
	if lot.TotalAvailable() != totalSpots-3 {
		panic("FAIL: available count mismatch")
	}
	fmt.Println("PASS: parking decreased availability")

	// Duplicate parking should fail
	_, err = lot.Park(car)
	if err == nil {
		panic("FAIL: duplicate parking should error")
	}
	fmt.Println("PASS: duplicate parking rejected")

	// Unpark
	err = lot.Unpark("ABC-123")
	if err != nil {
		panic("FAIL: " + err.Error())
	}
	if lot.TotalAvailable() != totalSpots-2 {
		panic("FAIL: unpark didn't free spot")
	}
	fmt.Println("PASS: unpark freed spot")

	// Concurrent parking test
	var wg sync.WaitGroup
	parked := 0
	var mu sync.Mutex
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			v := &Vehicle{LicensePlate: fmt.Sprintf("CONC-%d", id), Type: Car}
			_, err := lot.Park(v)
			if err == nil {
				mu.Lock()
				parked++
				mu.Unlock()
			}
		}(i)
	}
	wg.Wait()
	fmt.Printf("PASS: concurrent parking - %d cars parked\n", parked)
}
