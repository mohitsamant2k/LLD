package vendingmachine

import (
	"fmt"
)

type IdleState struct {
	vm *VendingMachine
}

func (i *IdleState) SelectItem(item string) state {
	if i.vm.IsItemExist(item) {
		fmt.Printf("Item exist please Pay : %+v\n", i.vm.items[item].price)
		return &TakeMoneyState{
			vm:      i.vm,
			balance: i.vm.items[item].price,
			item:    item,
		}
	}
	fmt.Println("item do not exists")
	return i
}

func (i *IdleState) InsertCoin(coin int) state {
	fmt.Println("Idle state do not support InsertCoin")
	return i
}

func (i *IdleState) Dispense() state {
	fmt.Println("Idle state do not support dispense")
	return i
}