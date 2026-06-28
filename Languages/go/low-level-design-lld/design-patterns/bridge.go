package main

import "fmt"

// Bridge Pattern: Decouple an abstraction from its implementation so that
// the two can vary independently.
// Example: Notifications (Email/SMS/Push) × Priority (Urgent/Normal).

// Implementation interface - how to send
type MessageSender interface {
	Send(to, subject, body string) bool
	Name() string
}

// Concrete Implementation: Email
type EmailSender struct{}

func (e *EmailSender) Name() string { return "Email" }
func (e *EmailSender) Send(to, subject, body string) bool {
	fmt.Printf("  [Email] To: %s | Subject: %s | Body: %s\n", to, subject, body)
	return true
}

// Concrete Implementation: SMS
type SMSSender struct{}

func (s *SMSSender) Name() string { return "SMS" }
func (s *SMSSender) Send(to, subject, body string) bool {
	// SMS has character limit
	msg := body
	if len(msg) > 50 {
		msg = msg[:47] + "..."
	}
	fmt.Printf("  [SMS] To: %s | %s\n", to, msg)
	return true
}

// Concrete Implementation: Push Notification
type PushSender struct{}

func (p *PushSender) Name() string { return "Push" }
func (p *PushSender) Send(to, subject, body string) bool {
	fmt.Printf("  [Push] To: %s | Title: %s | %s\n", to, subject, body)
	return true
}

// Abstraction - what kind of notification
type Notification interface {
	Notify(to, message string) bool
}

// Refined Abstraction: Urgent notification
type UrgentNotification struct {
	sender MessageSender
}

func NewUrgentNotification(s MessageSender) *UrgentNotification {
	return &UrgentNotification{sender: s}
}

func (n *UrgentNotification) Notify(to, message string) bool {
	subject := "🚨 URGENT: " + message
	body := "[IMMEDIATE ACTION REQUIRED] " + message
	return n.sender.Send(to, subject, body)
}

// Refined Abstraction: Normal notification
type NormalNotification struct {
	sender MessageSender
}

func NewNormalNotification(s MessageSender) *NormalNotification {
	return &NormalNotification{sender: s}
}

func (n *NormalNotification) Notify(to, message string) bool {
	return n.sender.Send(to, message, message)
}

func main() {
	// Bridge lets us combine any notification type with any sender
	senders := []MessageSender{&EmailSender{}, &SMSSender{}, &PushSender{}}

	fmt.Println("--- Urgent Notifications ---")
	for _, sender := range senders {
		notif := NewUrgentNotification(sender)
		ok := notif.Notify("alice@example.com", "Server is down")
		if !ok {
			panic("FAIL: urgent " + sender.Name())
		}
	}

	fmt.Println("\n--- Normal Notifications ---")
	for _, sender := range senders {
		notif := NewNormalNotification(sender)
		ok := notif.Notify("bob@example.com", "Weekly report ready")
		if !ok {
			panic("FAIL: normal " + sender.Name())
		}
	}

	// Verify we can swap implementations without changing abstraction
	urgentPush := NewUrgentNotification(&PushSender{})
	normalEmail := NewNormalNotification(&EmailSender{})

	fmt.Println("\n--- Mix and match ---")
	urgentPush.Notify("charlie", "Disk full")
	normalEmail.Notify("dave", "Build succeeded")

	fmt.Println("\nPASS: bridge pattern complete")
}
