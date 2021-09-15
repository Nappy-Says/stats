package stats

import (
	"github.com/nappy-says/bank/pkg/bank/types"
	"fmt"
)

func ExampleAvg_positive() {
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   10,
			Category: "Food",
		},
		{
			ID:       2,
			Amount:   20,
			Category: "Food",
		},
		{
			ID:       3,
			Amount:   30,
			Category: "Carshering",
		},
		{
			ID:       4,
			Amount:   40,
			Category: "Credit",
		},
	}

	fmt.Println(Avg(payments))

	// Output:
	// 25
}


func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   10,
			Category: "Food",
		},
		{
			ID:       2,
			Amount:   20,
			Category: "Food",
		},
		{
			ID:       3,
			Amount:   30,
			Category: "Carshering",
		},
		{
			ID:       4,
			Amount:   40,
			Category: "Credit",
		},
	}

	fmt.Println(TotalInCategory(payments, "Food"))

	//Output:
	// 30
}
