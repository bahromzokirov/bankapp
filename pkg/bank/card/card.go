package card

import (
	"bank/pkg/bank/types"
)

func PaymentSources(cards []types.Card) []types.PaymentSource {
	c := []types.PaymentSource{}
	for _, card := range cards {
		if card.Balance > 0 && card.Active {
			c = append(c, types.PaymentSource{
				Number:  card.PAN,
				Balance: card.Balance,
				Active:  card.Active,
			})
		}
	}
	return c
}
