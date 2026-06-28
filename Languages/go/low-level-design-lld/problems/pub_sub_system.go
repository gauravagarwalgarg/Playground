package main

import (
	"fmt"
	"sync"
	"time"
)

// Pub/Sub System LLD - In-memory topic-based publish/subscribe.
// Demonstrates: Observer pattern, thread-safety, decoupled communication.

type Message struct {
	Topic     string
	Payload   string
	Timestamp time.Time
}

// Subscriber interface
type Subscriber interface {
	ID() string
	Receive(msg Message)
}

// Concrete subscriber
type ConsoleSubscriber struct {
	id       string
	mu       sync.Mutex
	messages []Message
}

func NewConsoleSubscriber(id string) *ConsoleSubscriber {
	return &ConsoleSubscriber{id: id}
}

func (s *ConsoleSubscriber) ID() string { return s.id }

func (s *ConsoleSubscriber) Receive(msg Message) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.messages = append(s.messages, msg)
	fmt.Printf("  [%s] received on '%s': %s\n", s.id, msg.Topic, msg.Payload)
}

func (s *ConsoleSubscriber) MessageCount() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.messages)
}

// Topic manages subscribers for a specific topic
type Topic struct {
	mu          sync.RWMutex
	name        string
	subscribers map[string]Subscriber
}

func NewTopic(name string) *Topic {
	return &Topic{
		name:        name,
		subscribers: make(map[string]Subscriber),
	}
}

func (t *Topic) Subscribe(sub Subscriber) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.subscribers[sub.ID()] = sub
}

func (t *Topic) Unsubscribe(subID string) {
	t.mu.Lock()
	defer t.mu.Unlock()
	delete(t.subscribers, subID)
}

func (t *Topic) Publish(payload string) {
	t.mu.RLock()
	defer t.mu.RUnlock()
	msg := Message{
		Topic:     t.name,
		Payload:   payload,
		Timestamp: time.Now(),
	}
	for _, sub := range t.subscribers {
		sub.Receive(msg)
	}
}

func (t *Topic) SubscriberCount() int {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return len(t.subscribers)
}

// Broker - manages multiple topics
type MessageBroker struct {
	mu     sync.RWMutex
	topics map[string]*Topic
}

func NewMessageBroker() *MessageBroker {
	return &MessageBroker{topics: make(map[string]*Topic)}
}

func (b *MessageBroker) CreateTopic(name string) *Topic {
	b.mu.Lock()
	defer b.mu.Unlock()
	if t, exists := b.topics[name]; exists {
		return t
	}
	t := NewTopic(name)
	b.topics[name] = t
	return t
}

func (b *MessageBroker) Subscribe(topicName string, sub Subscriber) {
	topic := b.CreateTopic(topicName)
	topic.Subscribe(sub)
}

func (b *MessageBroker) Unsubscribe(topicName, subID string) {
	b.mu.RLock()
	defer b.mu.RUnlock()
	if t, exists := b.topics[topicName]; exists {
		t.Unsubscribe(subID)
	}
}

func (b *MessageBroker) Publish(topicName, payload string) {
	b.mu.RLock()
	topic, exists := b.topics[topicName]
	b.mu.RUnlock()
	if exists {
		topic.Publish(payload)
	}
}

func main() {
	broker := NewMessageBroker()

	// Create subscribers
	sub1 := NewConsoleSubscriber("analytics-service")
	sub2 := NewConsoleSubscriber("notification-service")
	sub3 := NewConsoleSubscriber("audit-service")

	// Subscribe to topics
	broker.Subscribe("user.created", sub1)
	broker.Subscribe("user.created", sub2)
	broker.Subscribe("user.created", sub3)
	broker.Subscribe("order.placed", sub1)
	broker.Subscribe("order.placed", sub2)

	// Publish messages
	fmt.Println("--- Publishing user.created ---")
	broker.Publish("user.created", `{"user_id": "u123", "email": "alice@example.com"}`)

	fmt.Println("\n--- Publishing order.placed ---")
	broker.Publish("order.placed", `{"order_id": "o456", "total": 99.99}`)

	// Verify message delivery
	if sub1.MessageCount() != 2 {
		panic(fmt.Sprintf("FAIL: sub1 expected 2 messages, got %d", sub1.MessageCount()))
	}
	if sub2.MessageCount() != 2 {
		panic(fmt.Sprintf("FAIL: sub2 expected 2 messages, got %d", sub2.MessageCount()))
	}
	if sub3.MessageCount() != 1 {
		panic(fmt.Sprintf("FAIL: sub3 expected 1 message, got %d", sub3.MessageCount()))
	}
	fmt.Println("\nPASS: message counts correct")

	// Unsubscribe and verify
	broker.Unsubscribe("user.created", "audit-service")
	fmt.Println("\n--- After unsubscribe, publishing user.created ---")
	broker.Publish("user.created", `{"user_id": "u789"}`)

	if sub3.MessageCount() != 1 {
		panic("FAIL: sub3 should not receive after unsubscribe")
	}
	fmt.Println("PASS: unsubscribe works")

	// Concurrent publish test
	fmt.Println("\n--- Concurrent publish test ---")
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			broker.Publish("order.placed", fmt.Sprintf(`{"id":%d}`, id))
		}(i)
	}
	wg.Wait()
	fmt.Printf("PASS: concurrent publish complete, sub1 total: %d messages\n", sub1.MessageCount())
}
