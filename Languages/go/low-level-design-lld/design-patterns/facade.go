package main

import "fmt"

// Facade Pattern: Provides a simplified interface to a complex subsystem.
// Hides internal complexity behind a single, easy-to-use API.

// Subsystem 1: Inventory
type InventoryService struct {
	stock map[string]int
}

func NewInventoryService() *InventoryService {
	return &InventoryService{
		stock: map[string]int{"laptop": 10, "phone": 25, "tablet": 5},
	}
}

func (s *InventoryService) CheckStock(product string) bool {
	return s.stock[product] > 0
}

func (s *InventoryService) Reserve(product string) {
	s.stock[product]--
}

// Subsystem 2: Payment
type PaymentService struct{}

func (p *PaymentService) Charge(userID string, amount float64) bool {
	fmt.Printf("[Payment] Charged $%.2f to user %s\n", amount, userID)
	return true
}

func (p *PaymentService) Refund(userID string, amount float64) {
	fmt.Printf("[Payment] Refunded $%.2f to user %s\n", amount, userID)
}

// Subsystem 3: Shipping
type ShippingService struct{}

func (s *ShippingService) CreateShipment(userID, product string) string {
	trackingID := fmt.Sprintf("SHIP-%s-%s", userID, product)
	fmt.Printf("[Shipping] Created shipment %s\n", trackingID)
	return trackingID
}

// Subsystem 4: Notification
type NotificationService struct{}

func (n *NotificationService) SendEmail(userID, message string) {
	fmt.Printf("[Notification] Email to %s: %s\n", userID, message)
}

// Facade - single entry point for placing an order
type OrderFacade struct {
	inventory    *InventoryService
	payment      *PaymentService
	shipping     *ShippingService
	notification *NotificationService
}

func NewOrderFacade() *OrderFacade {
	return &OrderFacade{
		inventory:    NewInventoryService(),
		payment:      &PaymentService{},
		shipping:     &ShippingService{},
		notification: &NotificationService{},
	}
}

type OrderResult struct {
	Success    bool
	TrackingID string
	Error      string
}

func (f *OrderFacade) PlaceOrder(userID, product string, price float64) OrderResult {
	// Step 1: Check inventory
	if !f.inventory.CheckStock(product) {
		return OrderResult{Error: "out of stock"}
	}

	// Step 2: Process payment
	if !f.payment.Charge(userID, price) {
		return OrderResult{Error: "payment failed"}
	}

	// Step 3: Reserve inventory
	f.inventory.Reserve(product)

	// Step 4: Create shipment
	trackingID := f.shipping.CreateShipment(userID, product)

	// Step 5: Notify user
	f.notification.SendEmail(userID, fmt.Sprintf("Order confirmed! Tracking: %s", trackingID))

	return OrderResult{Success: true, TrackingID: trackingID}
}

func main() {
	facade := NewOrderFacade()

	// Client only needs one call - complexity is hidden
	result := facade.PlaceOrder("user123", "laptop", 999.99)

	if result.Success && result.TrackingID != "" {
		fmt.Println("\nPASS: order placed, tracking:", result.TrackingID)
	} else {
		panic("FAIL: order not placed")
	}

	// Out of stock scenario
	for i := 0; i < 10; i++ {
		facade.PlaceOrder("bulk", "tablet", 499.99)
	}
	result2 := facade.PlaceOrder("user456", "tablet", 499.99)
	if result2.Error == "out of stock" {
		fmt.Println("PASS: out of stock detected")
	} else {
		panic("FAIL: should be out of stock")
	}
}
