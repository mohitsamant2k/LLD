package vendingmachine

import (
	"fmt"
)


type TakeMoneyState struct {
	vm      *VendingMachine
	balance int
	item    string
}

func (i *TakeMoneyState) SelectItem(item string) state {
	fmt.Println("TakeMoneyState already seleted the item")
	return i
}

func (i *TakeMoneyState) InsertCoin(coin int) state {
	i.balance -= coin
	if i.balance <= 0 {
		return &DispenseState{
			vm:     i.vm,
			item:   i.item,
			refund: -i.balance,
		}
	}
	return i
}

func (i *TakeMoneyState) Dispense() state {
	fmt.Println("TakeMoneyState state do not support dispense")
	return i
}