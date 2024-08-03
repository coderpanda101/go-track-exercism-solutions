package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	// panic("CalculateWorkingCarsPerHour not implemented")
	successRateFract := successRate / 100
	return (float64(productionRate) * successRateFract)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	// panic("CalculateWorkingCarsPerMinute not implemented")
	successRateFract := successRate / 100
	return int((float64(productionRate) * successRateFract) / 60)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	// panic("CalculateCost not implemented")
	total := 0.0
	tens := (carsCount / 10) * 10
	units := carsCount % 10

	total += float64(tens) * 9500.0
	total += float64(units) * 10000.0

	return uint(total)
}
