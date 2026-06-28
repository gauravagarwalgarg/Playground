package main

import "fmt"

// Adapter Pattern: Allows incompatible interfaces to work together.
// Wraps an existing class with a new interface so it can be used where
// a different interface is expected.

// Target interface - what the client expects
type PaymentProcessor interface {
	Pay(amount float64) string
	Refund(transactionID string) bool
}

// Adaptee - legacy system with incompatible interface
type LegacyBankGateway struct {
	connected bool
}

func (g *LegacyBankGateway) Connect() {
	g.connected = true
}

func (g *LegacyBankGateway) ExecuteTransfer(cents int64, ref string) string {
	if !g.connected {
		return ""
	}
	return fmt.Sprintf("LEG-%s-%d", ref, cents)
}

func (g *LegacyBankGateway) ReverseTransfer(legacyRef string) int {
	if !g.connected {
		return -1
	}
	return 0 // 0 = success
}

// Adapter - bridges the gap between Target and Adaptee
type BankGatewayAdapter struct {
	gateway *LegacyBankGateway
	counter int
}

func NewBankGatewayAdapter() *BankGatewayAdapter {
	gw := &LegacyBankGateway{}
	gw.Connect()
	return &BankGatewayAdapter{gateway: gw}
}

func (a *BankGatewayAdapter) Pay(amount float64) string {
	a.counter++
	cents := int64(amount * 100)
	ref := fmt.Sprintf("TXN%d", a.counter)
	return a.gateway.ExecuteTransfer(cents, ref)
}

func (a *BankGatewayAdapter) Refund(transactionID string) bool {
	code := a.gateway.ReverseTransfer(transactionID)
	return code == 0
}

// Modern payment processor that already conforms
type StripeProcessor struct{}

func (s *StripeProcessor) Pay(amount float64) string {
	return fmt.Sprintf("STRIPE-%.2f", amount)
}

func (s *StripeProcessor) Refund(transactionID string) bool {
	return true
}

// Client code works with any PaymentProcessor
func checkout(p PaymentProcessor, amount float64) string {
	return p.Pay(amount)
}

func main() {
	// Use the adapter to plug legacy system into modern interface
	adapter := NewBankGatewayAdapter()
	txn := checkout(adapter, 49.99)
	if txn != "" {
		fmt.Println("PASS: adapter Pay returned:", txn)
	} else {
		panic("FAIL: adapter Pay")
	}

	if adapter.Refund(txn) {
		fmt.Println("PASS: adapter Refund")
	} else {
		panic("FAIL: adapter Refund")
	}

	// Modern processor works the same way
	stripe := &StripeProcessor{}
	txn2 := checkout(stripe, 99.00)
	if txn2 == "STRIPE-99.00" {
		fmt.Println("PASS: stripe Pay")
	} else {
		panic("FAIL: stripe Pay, got: " + txn2)
	}
}
