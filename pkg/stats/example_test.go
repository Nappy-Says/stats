package stats

import (
	"fmt"
	"github.com/nappy-says/bank/v2/pkg/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			ID:       	1,
			Amount:   	10,
			Status:  	"FAIL",
			Category:	"Food",
		},
		{
			ID:       	2,
			Amount:   	20,
			Status: 	"INPROGRESS",
			Category: 	"Food",
		},
		{
			ID:       	3,
			Amount:   	30,
			Status: 	"OK",
			Category: 	"Carshering",
		},
		{
			ID:       	4,
			Amount:   	40,
			Status: 	"FAIL",
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
			Status:  	"FAIL",
			Category:	"Food",
		},
		{
			ID:       	2,
			Amount:   	20,
			Status: 	"INPROGRESS",
			Category: 	"Food",
		},
		{
			ID:       	3,
			Amount:   	30,
			Status: 	"OK",
			Category: 	"Carshering",
		},
		{
			ID:       	4,
			Amount:   	40,
			Status: 	"FAIL",
			Category: 	"Credit",
		},
	}

	fmt.Println(TotalInCategory(payments, "Food"))

	//Output:
	// 20
}
