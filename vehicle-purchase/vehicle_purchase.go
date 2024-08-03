package purchase

import "fmt"

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	// panic("NeedsLicense not implemented")
	if kind == "car" || kind == "truck" {
		return true
	} else {
		return false
	}
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	for i := 0; i < len(option1) && i < len(option2); i++ {
		if option1[i] < option2[i] {
			return fmt.Sprintf("%v is clearly the better choice.", option1)
		} else if option1[i] > option2[i] {
			return fmt.Sprintf("%v is clearly the better choice.", option2)
		}
	}

	if len(option1) < len(option2) {
		return fmt.Sprintf("%v is clearly the better choice.", option1)
	} else {
		return fmt.Sprintf("%v is clearly the better choice.", option2)
	}
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	// panic("CalculateResellPrice not implemented")
	if age < 3.0 {
		return originalPrice * 0.8
	} else if age >= 10.0 {
		return originalPrice * 0.5
	} else {
		return originalPrice * 0.7
	}
}
