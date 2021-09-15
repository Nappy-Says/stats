package stats

import (
	"fmt"
	"github.com/nappy-says/bank/v2/pkg/types"
)

func ExampleAvg_positive() {
	payments := []types.Payment{
		{
			ID:       	1,
			Amount:   	10,
			Status:  	types.StatusFail,
			Category:	"Food",
		},
		{
			ID:       	2,
			Amount:   	20,
			Status: 	types.StatusInProgress,
			Category: 	"Food",
		},
		{
			ID:       	3,
			Amount:   	30,
			Status: 	types.StatusOK,
			Category: 	"Carshering",
		},
		{
			ID:       	4,
			Amount:   	40,
			Status: 	types.StatusFail,
			Category: 	"Credit",
		},
	}

	fmt.Println(Avg(payments))

	// Output:
	// 12
}


func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			ID:       	1,
			Amount:   	10,
			Status:  	types.StatusFail,
			Category:	"Food",
		},
		{
			ID:       	2,
			Amount:   	20,
			Status: 	types.StatusInProgress,
			Category: 	"Food",
		},
		{
			ID:       	3,
			Amount:   	30,
			Status: 	types.StatusOK,
			Category: 	"Carshering",
		},
		{
			ID:       	4,
			Amount:   	40,
			Status: 	types.StatusFail,
			Category: 	"Credit",
		},
	}

	fmt.Println(TotalInCategory(payments, "Food"))

	//Output:
	// 20
}
