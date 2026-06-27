package main

import (
	"fmt"
	"sync"
)

// Singleton demonstrates thread-safe singleton using sync.Once.
// sync.Once guarantees the initialization function runs exactly once,
// even when called concurrently from multiple goroutines.

type database struct {
	connection string
}

var (
	dbInstance *database
	once       sync.Once
)

// GetDatabase returns the singleton database instance.
// The first call initializes the connection; subsequent calls return the same instance.
func GetDatabase() *database {
	once.Do(func() {
		fmt.Println("Creating database connection (this runs only once)")
		dbInstance = &database{connection: "postgres://localhost:5432/mydb"}
	})
	return dbInstance
}

func main() {
	// Demonstrate thread-safety: spawn multiple goroutines requesting the singleton
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			db := GetDatabase()
			fmt.Printf("Goroutine %d got connection: %s\n", id, db.connection)
		}(i)
	}

	wg.Wait()

	// Verify same instance
	db1 := GetDatabase()
	db2 := GetDatabase()
	fmt.Printf("Same instance? %v\n", db1 == db2)
}
