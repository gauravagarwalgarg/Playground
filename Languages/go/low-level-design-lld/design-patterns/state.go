package main

import "fmt"

// State Pattern: Allow an object to change its behavior when its internal
// state changes. The object will appear to change its class.
// Classic example: Vending Machine states.

// State interface
type VendingState interface {
	InsertCoin(v *VendingMachine)
	SelectProduct(v *VendingMachine)
	Dispense(v *VendingMachine)
	String() string
}

// Context
type VendingMachine struct {
	state   VendingState
	product string
}

func NewVendingMachine() *VendingMachine {
	vm := &VendingMachine{}
	vm.state = &IdleState{}
	return vm
}

func (vm *VendingMachine) SetState(s VendingState) {
	fmt.Printf("  [State Transition] %s → %s\n", vm.state, s)
	vm.state = s
}

func (vm *VendingMachine) InsertCoin()    { vm.state.InsertCoin(vm) }
func (vm *VendingMachine) SelectProduct() { vm.state.SelectProduct(vm) }
func (vm *VendingMachine) Dispense()      { vm.state.Dispense(vm) }

// Concrete State: Idle (waiting for coin)
type IdleState struct{}

func (s *IdleState) String() string { return "Idle" }

func (s *IdleState) InsertCoin(v *VendingMachine) {
	fmt.Println("  Coin inserted")
	v.SetState(&HasCoinState{})
}

func (s *IdleState) SelectProduct(v *VendingMachine) {
	fmt.Println("  Please insert a coin first")
}

func (s *IdleState) Dispense(v *VendingMachine) {
	fmt.Println("  Please insert a coin and select a product")
}

// Concrete State: HasCoin (coin inserted, waiting for selection)
type HasCoinState struct{}

func (s *HasCoinState) String() string { return "HasCoin" }

func (s *HasCoinState) InsertCoin(v *VendingMachine) {
	fmt.Println("  Coin already inserted")
}

func (s *HasCoinState) SelectProduct(v *VendingMachine) {
	fmt.Println("  Product selected")
	v.product = "Soda"
	v.SetState(&DispensingState{})
}

func (s *HasCoinState) Dispense(v *VendingMachine) {
	fmt.Println("  Please select a product first")
}

// Concrete State: Dispensing
type DispensingState struct{}

func (s *DispensingState) String() string { return "Dispensing" }

func (s *DispensingState) InsertCoin(v *VendingMachine) {
	fmt.Println("  Please wait, dispensing in progress")
}

func (s *DispensingState) SelectProduct(v *VendingMachine) {
	fmt.Println("  Already dispensing")
}

func (s *DispensingState) Dispense(v *VendingMachine) {
	fmt.Printf("  Dispensing: %s\n", v.product)
	v.product = ""
	v.SetState(&IdleState{})
}

func main() {
	vm := NewVendingMachine()

	// Invalid actions in idle state
	fmt.Println("--- Test invalid actions ---")
	vm.SelectProduct()
	vm.Dispense()

	// Normal flow
	fmt.Println("\n--- Normal flow ---")
	vm.InsertCoin()
	vm.SelectProduct()
	vm.Dispense()

	// Verify back to idle
	fmt.Println("\n--- Verify reset to Idle ---")
	vm.InsertCoin()
	vm.InsertCoin() // duplicate coin
	vm.SelectProduct()
	vm.Dispense()

	fmt.Println("\nPASS: state pattern complete")
}
