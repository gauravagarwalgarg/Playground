package main

import "fmt"

// Mediator Pattern: Reduces chaotic dependencies between objects by restricting
// direct communications, forcing them to collaborate via a mediator object.
// Example: Chat room (users communicate through the room, not directly).

// Mediator interface
type ChatMediator interface {
	Register(user *ChatUser)
	SendMessage(from *ChatUser, message string)
	SendDirect(from *ChatUser, toName string, message string)
}

// Colleague
type ChatUser struct {
	name     string
	mediator ChatMediator
	inbox    []string
}

func NewChatUser(name string, mediator ChatMediator) *ChatUser {
	u := &ChatUser{name: name, mediator: mediator}
	mediator.Register(u)
	return u
}

func (u *ChatUser) Send(message string) {
	fmt.Printf("  [%s] sends: %s\n", u.name, message)
	u.mediator.SendMessage(u, message)
}

func (u *ChatUser) SendTo(toName, message string) {
	fmt.Printf("  [%s] DM to %s: %s\n", u.name, toName, message)
	u.mediator.SendDirect(u, toName, message)
}

func (u *ChatUser) Receive(from, message string) {
	msg := fmt.Sprintf("%s: %s", from, message)
	u.inbox = append(u.inbox, msg)
	fmt.Printf("  [%s] received: %s\n", u.name, msg)
}

func (u *ChatUser) InboxCount() int { return len(u.inbox) }

// Concrete Mediator: ChatRoom
type ChatRoom struct {
	name  string
	users map[string]*ChatUser
}

func NewChatRoom(name string) *ChatRoom {
	return &ChatRoom{name: name, users: make(map[string]*ChatUser)}
}

func (r *ChatRoom) Register(user *ChatUser) {
	r.users[user.name] = user
	fmt.Printf("  [Room:%s] %s joined\n", r.name, user.name)
}

func (r *ChatRoom) SendMessage(from *ChatUser, message string) {
	for name, user := range r.users {
		if name != from.name {
			user.Receive(from.name, message)
		}
	}
}

func (r *ChatRoom) SendDirect(from *ChatUser, toName string, message string) {
	if user, exists := r.users[toName]; exists {
		user.Receive(from.name, message)
	}
}

func main() {
	room := NewChatRoom("engineering")

	alice := NewChatUser("Alice", room)
	bob := NewChatUser("Bob", room)
	charlie := NewChatUser("Charlie", room)

	// Broadcast: Alice sends to everyone
	fmt.Println("\n--- Broadcast ---")
	alice.Send("Hello everyone!")

	if bob.InboxCount() != 1 || charlie.InboxCount() != 1 || alice.InboxCount() != 0 {
		panic(fmt.Sprintf("FAIL: broadcast delivery. Alice=%d Bob=%d Charlie=%d",
			alice.InboxCount(), bob.InboxCount(), charlie.InboxCount()))
	}
	fmt.Println("PASS: broadcast delivered to all except sender")

	// Direct message: Bob sends only to Alice
	fmt.Println("\n--- Direct Message ---")
	bob.SendTo("Alice", "Hey Alice, can you review my PR?")

	// Alice: 0(she was sender in broadcast) + 1(DM from Bob) = 1
	if alice.InboxCount() != 1 {
		panic(fmt.Sprintf("FAIL: DM not received, alice inbox=%d", alice.InboxCount()))
	}
	if charlie.InboxCount() != 1 {
		panic("FAIL: Charlie should NOT receive DM")
	}
	fmt.Println("PASS: DM delivered only to target")

	// Another broadcast from Charlie
	fmt.Println("\n--- Another broadcast ---")
	charlie.Send("Meeting in 5 minutes")
	// Alice: 1(DM) + 1(charlie broadcast) = 2
	// Bob: 1(alice broadcast) + 1(charlie broadcast) = 2
	if alice.InboxCount() != 2 {
		panic(fmt.Sprintf("FAIL: alice expected 2, got %d", alice.InboxCount()))
	}
	if bob.InboxCount() != 2 {
		panic(fmt.Sprintf("FAIL: bob expected 2, got %d", bob.InboxCount()))
	}
	fmt.Println("PASS: mediator pattern complete")
}
