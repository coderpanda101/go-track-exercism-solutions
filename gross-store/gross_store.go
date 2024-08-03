package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	measurements := make(map[string]int)
	measurements["quarter_of_a_dozen"] = 3
	measurements["half_of_a_dozen"] = 6
	measurements["dozen"] = 12
	measurements["small_gross"] = 120
	measurements["gross"] = 144
	measurements["great_gross"] = 1728

	return measurements
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {

	valUnit, existsUnit := units[unit]
	if !existsUnit {
		return false
	}

	_, existsItem := bill[item]
	if existsItem {
		currQty := bill[item]
		bill[item] = currQty + valUnit
		return true
	} else {
		bill[item] = valUnit
		return true
	}

}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {

	itemQty, itemExists := bill[item]
	unitQty, unitExists := units[unit]
	if !itemExists || !unitExists || itemQty-unitQty < 0 {
		return false
	}
	if itemQty-unitQty == 0 {
		delete(bill, item)
		return true
	} else {
		bill[item] = itemQty - unitQty
		return true
	}

}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	itemQty, itemExists := bill[item]
	if !itemExists {
		return 0, false
	} else {
		return itemQty, true
	}

}
