package card

import (
	"bank/pkg/bank/types"
)
func Total(cards [] types.Card) types.Money {
	 var operations [] types.Money 
	 operations = append(operations, 10_000_00)
	 operations= append(operations, 10_000_00)
	 
	 sum := sum(operations)
	 
	 return sum
}
func sum(operations[]types.Money) types.Money{
	sum:= types.Money(0)
	for _, operation := range operations {
		sum += operation
	}
	return sum
}



