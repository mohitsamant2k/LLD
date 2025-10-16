package vendingmachine

import (
	"fmt"
)

type DispenseState struct {
	vm     *VendingMachine
	item   string
	refund int
}

func (i *DispenseState) SelectItem(item string) state {
	fmt.Println("DispenseState already seleted the item")
	return i
}

func (i *DispenseState) InsertCoin(coin int) state {
	fmt.Println("DispenseState balce completed")
	return i
}

func (i *DispenseState) Dispense() state {
	pq := i.vm.items[i.item]
	pq.quantity--
	i.vm.items[i.item] = pq
	fmt.Printf("dispensing item: %+v\n", i.item)
	fmt.Printf("returning money %+v\n", i.refund)

	return &IdleState{
		vm: i.vm,
	}
}