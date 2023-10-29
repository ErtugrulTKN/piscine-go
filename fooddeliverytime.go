package piscine

type food struct { // menu ve time tanımladığımz için struct kullandık.
	preptime int
}

func FoodDeliveryTime(order string) int {
	menu := map[string]food{
		"burger":  {preptime: 15},
		"chips":   {preptime: 10},
		"nuggets": {preptime: 12},
	}

	item, exists := menu[order]
	if !exists {
		return 404
	}

	return item.preptime
}
