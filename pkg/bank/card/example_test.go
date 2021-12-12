package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExamplePaymentSources() {
	cards := []types.Card{
		{
			Balance: 1_000_00,
			Active:  true,
			PAN:  "5058 1241 5784 8888",
		},
		{
			Balance: 10_000_00,
			Active:  false,
			PAN:  "5058 1242 5784 8888",
		},
		{
			Balance: -10_000_00,
			Active:  true,
			PAN:  "5058 1243 5784 8888",
		},
		{
			Balance: 2_000_00,
			Active:  true,
			PAN:  "5058 1244 5784 8888",
		},
	}
		sources := PaymentSources(cards)
		for _, source := range sources {
			fmt.Println(source.Number)
		}
	
	//Output:
	// 5058 1241 5784 8888
	// 5058 1244 5784 8888
}
