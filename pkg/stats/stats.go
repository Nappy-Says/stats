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
	dict := make(map[types.Category]int)

	for _, i := range payments {
		last := resultMap[i.Category]
		dict[i.Category] = dict[i.Category] + 1
		resultMap[i.Category] = last + i.Amount
	}

	for key, value := range dict {
		resultMap[key] /= types.Money(value)
	}

	return resultMap
}

func PeriodsDynamic(first map[types.Category]types.Money,
	second map[types.Category]types.Money) map[types.Category]types.Money {
	resultMap := map[types.Category]types.Money{}

	for key, value := range first {
			resultMap[key] = second[key] - value
	}


	if len(first) == len(second) { return resultMap }


	for key1, value1 := range second {
		if value1 > first[key1] && resultMap[key1] == 0{
			resultMap[key1] = value1
		}
	}

	return resultMap
}
