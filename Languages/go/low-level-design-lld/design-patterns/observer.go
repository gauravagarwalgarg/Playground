package main

import (
	"fmt"
	"sync"
)

// Observer Pattern: EventBus with Subscribe/Publish using callbacks.
// Subscribers register handlers for specific event types.
// When an event is published, all registered handlers are invoked.

// Event represents a generic event.
type Event struct {
	Type    string
	Payload interface{}
}

// Handler is a callback function for events.
type Handler func(Event)

// EventBus manages subscriptions and publishing.
type EventBus struct {
	mu          sync.RWMutex
	subscribers map[string][]Handler
}

func NewEventBus() *EventBus {
	return &EventBus{
		subscribers: make(map[string][]Handler),
	}
}

// Subscribe registers a handler for an event type.
func (eb *EventBus) Subscribe(eventType string, handler Handler) {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	eb.subscribers[eventType] = append(eb.subscribers[eventType], handler)
}

// Publish sends an event to all subscribers of that event type.
func (eb *EventBus) Publish(event Event) {
	eb.mu.RLock()
	defer eb.mu.RUnlock()

	handlers, ok := eb.subscribers[event.Type]
	if !ok {
		return
	}
	for _, handler := range handlers {
		handler(event)
	}
}

// PublishAsync sends events to handlers concurrently.
func (eb *EventBus) PublishAsync(event Event) {
	eb.mu.RLock()
	defer eb.mu.RUnlock()

	handlers, ok := eb.subscribers[event.Type]
	if !ok {
		return
	}

	var wg sync.WaitGroup
	for _, handler := range handlers {
		wg.Add(1)
		go func(h Handler) {
			defer wg.Done()
			h(event)
		}(handler)
	}
	wg.Wait()
}

func main() {
	bus := NewEventBus()

	// Subscribe to "user.created" events
	bus.Subscribe("user.created", func(e Event) {
		fmt.Printf("[EmailService] Sending welcome email to: %v\n", e.Payload)
	})

	bus.Subscribe("user.created", func(e Event) {
		fmt.Printf("[Analytics] Tracking new user: %v\n", e.Payload)
	})

	// Subscribe to "order.placed" events
	bus.Subscribe("order.placed", func(e Event) {
		fmt.Printf("[Inventory] Processing order: %v\n", e.Payload)
	})

	// Publish events
	fmt.Println("--- Publishing user.created ---")
	bus.Publish(Event{Type: "user.created", Payload: "alice@example.com"})

	fmt.Println("\n--- Publishing order.placed ---")
	bus.Publish(Event{Type: "order.placed", Payload: map[string]interface{}{
		"id":    "ORD-123",
		"total": 99.99,
	}})

	fmt.Println("\n--- Publishing async user.created ---")
	bus.PublishAsync(Event{Type: "user.created", Payload: "bob@example.com"})
}
