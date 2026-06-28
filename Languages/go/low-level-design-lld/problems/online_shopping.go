package main

import (
	"fmt"
	"sync"
	"time"
)

// Online Shopping System LLD
// Demonstrates: Composition, Strategy (payment), State (order), Repository pattern.
// Models: Product, Cart, Order, Payment strategy, Inventory management.

// --- Product ---
type Product struct {
	ID    string
	Name  string
	Price float64
}

// --- Inventory ---
type Inventory struct {
	mu    sync.Mutex
	stock map[string]int // productID → quantity
}

func NewInventory() *Inventory {
	return &Inventory{stock: make(map[string]int)}
}

func (inv *Inventory) AddStock(productID string, qty int) {
	inv.mu.Lock()
	defer inv.mu.Unlock()
	inv.stock[productID] += qty
}

func (inv *Inventory) Reserve(productID string, qty int) bool {
	inv.mu.Lock()
	defer inv.mu.Unlock()
	if inv.stock[productID] >= qty {
		inv.stock[productID] -= qty
		return true
	}
	return false
}

func (inv *Inventory) Release(productID string, qty int) {
	inv.mu.Lock()
	defer inv.mu.Unlock()
	inv.stock[productID] += qty
}

func (inv *Inventory) Stock(productID string) int {
	inv.mu.Lock()
	defer inv.mu.Unlock()
	return inv.stock[productID]
}

// --- Cart ---
type CartItem struct {
	Product  *Product
	Quantity int
}

type ShoppingCart struct {
	mu    sync.Mutex
	items map[string]*CartItem // productID → item
}

func NewShoppingCart() *ShoppingCart {
	return &ShoppingCart{items: make(map[string]*CartItem)}
}

func (c *ShoppingCart) Add(product *Product, qty int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if item, exists := c.items[product.ID]; exists {
		item.Quantity += qty
	} else {
		c.items[product.ID] = &CartItem{Product: product, Quantity: qty}
	}
}

func (c *ShoppingCart) Remove(productID string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.items, productID)
}

func (c *ShoppingCart) Total() float64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	total := 0.0
	for _, item := range c.items {
		total += item.Product.Price * float64(item.Quantity)
	}
	return total
}

func (c *ShoppingCart) Items() []*CartItem {
	c.mu.Lock()
	defer c.mu.Unlock()
	items := make([]*CartItem, 0, len(c.items))
	for _, item := range c.items {
		items = append(items, item)
	}
	return items
}

func (c *ShoppingCart) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items = make(map[string]*CartItem)
}

// --- Payment Strategy ---
type PaymentMethod interface {
	Pay(amount float64) (string, error)
	Name() string
}

type CreditCardPayment struct {
	CardNumber string
}

func (cc *CreditCardPayment) Name() string { return "CreditCard" }
func (cc *CreditCardPayment) Pay(amount float64) (string, error) {
	return fmt.Sprintf("CC-TXN-%s-%.0f", cc.CardNumber[len(cc.CardNumber)-4:], amount*100), nil
}

type WalletPayment struct {
	WalletID string
	Balance  float64
}

func (w *WalletPayment) Name() string { return "Wallet" }
func (w *WalletPayment) Pay(amount float64) (string, error) {
	if w.Balance < amount {
		return "", fmt.Errorf("insufficient wallet balance")
	}
	w.Balance -= amount
	return fmt.Sprintf("WALLET-TXN-%s", w.WalletID), nil
}

// --- Order ---
type OrderStatus int

const (
	OrderCreated OrderStatus = iota
	OrderPaid
	OrderShipped
	OrderDelivered
	OrderCancelled
)

func (s OrderStatus) String() string {
	return [...]string{"CREATED", "PAID", "SHIPPED", "DELIVERED", "CANCELLED"}[s]
}

type Order struct {
	ID        string
	UserID    string
	Items     []*CartItem
	Total     float64
	Status    OrderStatus
	PaymentID string
	CreatedAt time.Time
}

// --- Shopping Service (Facade) ---
type ShoppingService struct {
	mu        sync.Mutex
	inventory *Inventory
	orders    map[string]*Order
	orderSeq  int
}

func NewShoppingService(inv *Inventory) *ShoppingService {
	return &ShoppingService{
		inventory: inv,
		orders:    make(map[string]*Order),
	}
}

func (s *ShoppingService) Checkout(userID string, cart *ShoppingCart, payment PaymentMethod) (*Order, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	items := cart.Items()
	if len(items) == 0 {
		return nil, fmt.Errorf("cart is empty")
	}

	// Reserve inventory
	reserved := make([]struct {
		productID string
		qty       int
	}, 0)
	for _, item := range items {
		if !s.inventory.Reserve(item.Product.ID, item.Quantity) {
			// Rollback
			for _, r := range reserved {
				s.inventory.Release(r.productID, r.qty)
			}
			return nil, fmt.Errorf("insufficient stock for %s", item.Product.Name)
		}
		reserved = append(reserved, struct {
			productID string
			qty       int
		}{item.Product.ID, item.Quantity})
	}

	// Process payment
	total := cart.Total()
	txnID, err := payment.Pay(total)
	if err != nil {
		// Rollback inventory
		for _, r := range reserved {
			s.inventory.Release(r.productID, r.qty)
		}
		return nil, fmt.Errorf("payment failed: %w", err)
	}

	// Create order
	s.orderSeq++
	order := &Order{
		ID:        fmt.Sprintf("ORD-%d", s.orderSeq),
		UserID:    userID,
		Items:     items,
		Total:     total,
		Status:    OrderPaid,
		PaymentID: txnID,
		CreatedAt: time.Now(),
	}
	s.orders[order.ID] = order
	cart.Clear()
	return order, nil
}

func (s *ShoppingService) GetOrder(orderID string) (*Order, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	order, exists := s.orders[orderID]
	if !exists {
		return nil, fmt.Errorf("order not found")
	}
	return order, nil
}

func (s *ShoppingService) CancelOrder(orderID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	order, exists := s.orders[orderID]
	if !exists {
		return fmt.Errorf("order not found")
	}
	if order.Status != OrderPaid {
		return fmt.Errorf("can only cancel PAID orders")
	}
	// Release inventory
	for _, item := range order.Items {
		s.inventory.Release(item.Product.ID, item.Quantity)
	}
	order.Status = OrderCancelled
	return nil
}

func main() {
	// Setup
	inventory := NewInventory()
	products := map[string]*Product{
		"p1": {ID: "p1", Name: "Laptop", Price: 999.99},
		"p2": {ID: "p2", Name: "Mouse", Price: 29.99},
		"p3": {ID: "p3", Name: "Keyboard", Price: 79.99},
	}
	inventory.AddStock("p1", 5)
	inventory.AddStock("p2", 50)
	inventory.AddStock("p3", 30)

	service := NewShoppingService(inventory)

	// Add to cart
	cart := NewShoppingCart()
	cart.Add(products["p1"], 1)
	cart.Add(products["p2"], 2)
	cart.Add(products["p3"], 1)

	fmt.Printf("Cart total: $%.2f\n", cart.Total())
	expectedTotal := 999.99 + 29.99*2 + 79.99
	if cart.Total() != expectedTotal {
		panic(fmt.Sprintf("FAIL: expected $%.2f, got $%.2f", expectedTotal, cart.Total()))
	}
	fmt.Println("PASS: cart total correct")

	// Checkout with credit card
	cc := &CreditCardPayment{CardNumber: "4111111111111234"}
	order, err := service.Checkout("user1", cart, cc)
	if err != nil {
		panic("FAIL: " + err.Error())
	}
	fmt.Printf("Order: %s, Status: %s, Payment: %s\n", order.ID, order.Status, order.PaymentID)
	if order.Status != OrderPaid {
		panic("FAIL: order should be PAID")
	}
	fmt.Println("PASS: checkout succeeded")

	// Verify inventory decreased
	if inventory.Stock("p1") != 4 {
		panic("FAIL: inventory should decrease")
	}
	fmt.Println("PASS: inventory reserved")

	// Verify cart is cleared
	if cart.Total() != 0 {
		panic("FAIL: cart should be empty after checkout")
	}
	fmt.Println("PASS: cart cleared")

	// Test wallet payment with insufficient balance
	cart2 := NewShoppingCart()
	cart2.Add(products["p1"], 1)
	wallet := &WalletPayment{WalletID: "W001", Balance: 50.00}
	_, err = service.Checkout("user2", cart2, wallet)
	if err == nil {
		panic("FAIL: should fail with insufficient balance")
	}
	fmt.Println("PASS: insufficient wallet rejected:", err)

	// Verify inventory rolled back
	if inventory.Stock("p1") != 4 {
		panic("FAIL: inventory should be rolled back")
	}
	fmt.Println("PASS: inventory rolled back on payment failure")

	// Cancel order → inventory restored
	err = service.CancelOrder(order.ID)
	if err != nil {
		panic("FAIL: " + err.Error())
	}
	if inventory.Stock("p1") != 5 {
		panic("FAIL: inventory should restore on cancel")
	}
	fmt.Println("PASS: order cancelled, inventory restored")

	// Out of stock
	cart3 := NewShoppingCart()
	cart3.Add(products["p1"], 100) // more than available
	_, err = service.Checkout("user3", cart3, cc)
	if err == nil {
		panic("FAIL: should fail on out of stock")
	}
	fmt.Println("PASS: out of stock handled:", err)

	fmt.Println("\nPASS: online shopping system complete")
}
