package main

import "fmt"

// SOLID Principles in Go
// Each principle demonstrated with a practical example.

// =============================================================
// S - Single Responsibility Principle
// A struct should have only one reason to change.
// =============================================================

// BAD: One struct doing too much
// type UserService struct{} handles user CRUD AND sends emails AND generates reports

// GOOD: Separate responsibilities
type UserRepository struct{}

func (r *UserRepository) Save(name string) { fmt.Printf("  [UserRepo] Saved %s\n", name) }

type EmailService struct{}

func (e *EmailService) SendWelcome(email string) {
	fmt.Printf("  [Email] Welcome sent to %s\n", email)
}

type UserRegistration struct {
	repo  *UserRepository
	email *EmailService
}

func (ur *UserRegistration) Register(name, email string) {
	ur.repo.Save(name)
	ur.email.SendWelcome(email)
}

// =============================================================
// O - Open/Closed Principle
// Open for extension, closed for modification.
// Use interfaces to add new behavior without changing existing code.
// =============================================================

type DiscountCalculator interface {
	Calculate(price float64) float64
	Name() string
}

type RegularDiscount struct{}

func (d *RegularDiscount) Name() string               { return "Regular" }
func (d *RegularDiscount) Calculate(price float64) float64 { return price * 0.05 }

type PremiumDiscount struct{}

func (d *PremiumDiscount) Name() string               { return "Premium" }
func (d *PremiumDiscount) Calculate(price float64) float64 { return price * 0.20 }

// Adding a new discount type doesn't modify existing code - just add a new struct.
type SeasonalDiscount struct{ Rate float64 }

func (d *SeasonalDiscount) Name() string               { return "Seasonal" }
func (d *SeasonalDiscount) Calculate(price float64) float64 { return price * d.Rate }

func applyDiscount(price float64, d DiscountCalculator) float64 {
	discount := d.Calculate(price)
	fmt.Printf("  [%s] $%.2f discount on $%.2f\n", d.Name(), discount, price)
	return price - discount
}

// =============================================================
// L - Liskov Substitution Principle
// Subtypes must be substitutable for their base types.
// If it implements an interface, it should honor the contract fully.
// =============================================================

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width, Height float64
}

func (r *Rectangle) Area() float64 { return r.Width * r.Height }

type Square struct {
	Side float64
}

func (s *Square) Area() float64 { return s.Side * s.Side }

// GOOD: Both satisfy Shape interface and can be used interchangeably.
func printArea(s Shape) {
	fmt.Printf("  Area = %.2f\n", s.Area())
}

// =============================================================
// I - Interface Segregation Principle
// Clients should not depend on interfaces they don't use.
// Prefer many small interfaces over one large one.
// =============================================================

// BAD: type Worker interface { Work(); Eat(); Sleep(); }
// A Robot can Work() but not Eat() or Sleep().

// GOOD: Split into focused interfaces
type Workable interface {
	Work()
}

type Eatable interface {
	Eat()
}

type Human struct{ Name string }

func (h *Human) Work() { fmt.Printf("  [%s] working\n", h.Name) }
func (h *Human) Eat()  { fmt.Printf("  [%s] eating\n", h.Name) }

type Robot struct{ ID string }

func (r *Robot) Work() { fmt.Printf("  [Robot-%s] working\n", r.ID) }
// Robot doesn't implement Eatable - it doesn't need to.

func assignWork(w Workable) {
	w.Work()
}

// =============================================================
// D - Dependency Inversion Principle
// High-level modules should not depend on low-level modules.
// Both should depend on abstractions.
// =============================================================

// Abstraction
type Logger interface {
	Log(msg string)
}

// Low-level: Console logger
type ConsoleLogger struct{}

func (l *ConsoleLogger) Log(msg string) { fmt.Printf("  [Console] %s\n", msg) }

// Low-level: File logger (simulated)
type FileLogger struct{ Path string }

func (l *FileLogger) Log(msg string) { fmt.Printf("  [File:%s] %s\n", l.Path, msg) }

// High-level: Payment service depends on Logger abstraction, not concrete type
type PaymentProcessor struct {
	logger Logger
}

func NewPaymentProcessor(logger Logger) *PaymentProcessor {
	return &PaymentProcessor{logger: logger}
}

func (p *PaymentProcessor) Process(amount float64) {
	p.logger.Log(fmt.Sprintf("Processing payment of $%.2f", amount))
}

func main() {
	fmt.Println("=== S: Single Responsibility ===")
	reg := &UserRegistration{repo: &UserRepository{}, email: &EmailService{}}
	reg.Register("Alice", "alice@example.com")
	fmt.Println("PASS: SRP")

	fmt.Println("\n=== O: Open/Closed ===")
	price := 100.0
	applyDiscount(price, &RegularDiscount{})
	applyDiscount(price, &PremiumDiscount{})
	applyDiscount(price, &SeasonalDiscount{Rate: 0.30})
	fmt.Println("PASS: OCP - extended without modifying existing code")

	fmt.Println("\n=== L: Liskov Substitution ===")
	shapes := []Shape{
		&Rectangle{Width: 5, Height: 3},
		&Square{Side: 4},
	}
	for _, s := range shapes {
		printArea(s)
	}
	fmt.Println("PASS: LSP - both are valid Shapes")

	fmt.Println("\n=== I: Interface Segregation ===")
	human := &Human{Name: "Bob"}
	robot := &Robot{ID: "R2D2"}
	assignWork(human)
	assignWork(robot)
	human.Eat() // Only human eats
	fmt.Println("PASS: ISP - Robot only implements what it needs")

	fmt.Println("\n=== D: Dependency Inversion ===")
	consoleProcessor := NewPaymentProcessor(&ConsoleLogger{})
	consoleProcessor.Process(49.99)
	fileProcessor := NewPaymentProcessor(&FileLogger{Path: "/var/log/payments.log"})
	fileProcessor.Process(99.99)
	fmt.Println("PASS: DIP - high-level depends on abstraction")

	fmt.Println("\n=== All SOLID principles demonstrated ===")
}
