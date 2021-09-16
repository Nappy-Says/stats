package stats

import (
	"reflect"
	"testing"
	"github.com/nappy-says/bank/v2/pkg/types"
)

func TestCategoriesAvg(t *testing.T) {
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
			Category: "Food",
		},
		{
			ID:       4,
			Amount:   40,
			Category: "Food",
		},
		{
			ID:       5,
			Amount:   500,
			Category: "Attraction",
		},
		{
			ID:       6,
			Amount:   4000,
			Category: "XFit",
		},
		{
			ID:       7,
			Amount:   200,
			Category: "Internet",
		},
	}

	except := map[types.Category]types.Money{
		"Food":       25,
		"Attraction": 500,
		"XFit":       4000,
		"Internet":   200,
	}

	result := CategoriesAvg(payments)

	if !reflect.DeepEqual(except, result) {
		t.Errorf("invalid result, except: %v; actual : %v", except, result)
	}
}

func TestPeriodsDynamic(t *testing.T) {
	fitst := map[types.Category]types.Money{
		"auto": 10,
		"food": 20,
	}

	second := map[types.Category]types.Money{
		"auto": 5,
		"food": 3,
	}

	except := map[types.Category]types.Money{
		"auto": -5,
		"food": -17,
	}

	result := PeriodsDynamic(fitst, second)

	if !reflect.DeepEqual(except, result) {
		t.Errorf("except %v, actual %v", result, except)
	}
}
