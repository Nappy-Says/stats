package stats

import "github.com/nappy-says/bank/v2/pkg/types"



func Avg(payments []types.Payment) types.Money {
	avg_sum := types.Money(0)
	passCount := 0

	for _, i := range payments {
		if i.Status != "FAIL" {
			passCount +=1
			avg_sum += i.Amount
		}
	}

	result := types.Money(int(avg_sum) / passCount)

	return result 
}


func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sumByCategory := types.Money(0)
	
	for _, i := range payments {
		if i.Category == category && i.Status != "FAIL" {
			sumByCategory += i.Amount
		}
	}

	return sumByCategory
}

func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	resultMap := map[types.Category]types.Money{}
	dict:= make(map[types.Category]int)
	
	for _, i:= range payments {
		last := resultMap[i.Category]
		dict[i.Category] = dict[i.Category]+1
		resultMap[i.Category] = last + i.Amount	
	}
	
	for key, value := range dict {
		resultMap[key] /= types.Money(value)
	}
	
	// return dict
	return resultMap
}
