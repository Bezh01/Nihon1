package payment

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleMax() {
	payments := []types.Payment{
		{
			ID:     12,
			Amount: 20_000_00,
		},
		{
			ID:     34,
			Amount: 10_000_00,
		},
		{
			ID:     5,
			Amount: 90_000_00,
		},
		{
			ID:     45,
			Amount: 23_000_00,
		},
	}
	max := Max(payments)

	fmt.Println(max.ID)

	//Output: 5
}
