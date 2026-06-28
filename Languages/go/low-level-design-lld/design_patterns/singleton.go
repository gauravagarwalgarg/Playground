// Singleton Design Pattern in Go
// Uses sync.Once for thread-safe lazy initialization.
package main

import (
	"fmt"
	"sync"
)

type Database struct {
	Connection string
}

var (
	instance *Database
	once     sync.Once
)

func GetInstance() *Database {
	once.Do(func() {
		instance = &Database{Connection: "PostgreSQL@localhost:5432"}
	})
	return instance
}

func main() {
	db1 := GetInstance()
	db2 := GetInstance()

	if db1 != db2 {
		panic("FAIL: instances should be the same")
	}
	fmt.Println("PASS: same instance")

	if db1.Connection != "PostgreSQL@localhost:5432" {
		panic("FAIL: wrong connection string")
	}
	fmt.Println("PASS: correct connection")

	// Concurrent access test
	var wg sync.WaitGroup
	instances := make([]*Database, 100)
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			instances[idx] = GetInstance()
		}(i)
	}
	wg.Wait()

	for i := 1; i < 100; i++ {
		if instances[i] != instances[0] {
			panic("FAIL: concurrent instances differ")
		}
	}
	fmt.Println("PASS: thread-safe singleton")

	fmt.Println("All tests passed!")
}
