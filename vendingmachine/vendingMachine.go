package vendingmachine

import (
	"fmt"
)

type state interface {
	SelectItem(string) state
	InsertCoin(int) state
	Dispense() state
}

type PriceAndQuantity struct {
	price    int
	quantity int
}

type VendingMachine struct {
	items        map[string]PriceAndQuantity
	currentState state
}

func (v *VendingMachine) AddItems(item string, count int, price int) {
	if v.IsItemExist(item) {
		pq := v.items[item]
		pq.quantity += count
		pq.price = price
		v.items[item] = pq
	} else {
		v.items[item] = PriceAndQuantity{
			price:    price,
			quantity: count,
		}
	}
}

func (v VendingMachine) GetItems() []string {
	var item []string
	for k := range v.items {
		if v.items[k].quantity > 0 { // only expose non-empty items
			item = append(item, k)
		}
	}
	return item
}

func (v *VendingMachine) IsItemExist(item string) bool {
	pq, ok := v.items[item]
	return ok && pq.quantity > 0
}

func (vm *VendingMachine) purchase(item string , coins ...int)  {
	vm.currentState = vm.currentState.SelectItem(item)
	for _, c := range coins {
		vm.currentState = vm.currentState.InsertCoin(c)
	}
	vm.currentState = vm.currentState.Dispense()
	fmt.Printf("Remaining qty of %s: %d\n", item, vm.items[item].quantity)
}

func RunDemo() {
	fmt.Println("Vending Machine Demo Start")
	vm := &VendingMachine{
		items: make(map[string]PriceAndQuantity),
	}
	vm.currentState = &IdleState{vm: vm}

	// Add some items (name, quantity, price)
	vm.AddItems("chips", 2, 5)
	vm.AddItems("soda", 1, 7)
	vm.AddItems("gum", 3, 2)

	fmt.Println("Initial inventory:")
	for k, v := range vm.items {
		fmt.Printf("  %s => qty=%d price=%d\n", k, v.quantity, v.price)
	}

	// Buy soda (only one in stock) and chips twice
	vm.purchase("soda", 5, 2)  // exact (5+2=7)
	vm.purchase("chips", 2, 3) // pay in parts for price 5
	vm.purchase("chips", 5)    // second pack

	// Exhaust gum completely
	vm.purchase("gum", 2)    // first
	vm.purchase("gum", 2)    // second
	vm.purchase("gum", 1, 1) // third (over-pay split)

	// Try to buy exhausted items again
	vm.purchase("soda", 7)
	vm.purchase("chips", 5)
	vm.purchase("gum", 2)

	fmt.Println("\nFinal inventory (including zero-qty entries):")
	for k, v := range vm.items {
		fmt.Printf("  %s => qty=%d price=%d\n", k, v.quantity, v.price)
	}

	fmt.Println("Available (non-zero) items list:", vm.GetItems())
	fmt.Println("Vending Machine Demo End")
}
