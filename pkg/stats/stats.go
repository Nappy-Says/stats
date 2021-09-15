package stats

import (
	"github.com/nappy-says/bank/pkg/bank/types"
)



func Avg(payments []types.Payment) types.Money {
	avg_sum := types.Money(0)
	
	for _, i := range payments {
		avg_sum += i.Amount
	}

	result := types.Money(int(avg_sum) / len(payments))

	return result 
}


func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sumByCategory := types.Money(0)
	
	for _, i := range payments {
		if i.Category == category {
			sumByCategory += i.Amount
		}
	}

	return sumByCategory
}
