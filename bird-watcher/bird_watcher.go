package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	// panic("Please implement the TotalBirdCount() function")

	total := 0

	for i := 0; i < len(birdsPerDay); i++ {
		total += birdsPerDay[i]
	}

	return total
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	// panic("Please implement the BirdsInWeek() function")

	totalInWeek := 0

	start := (week * 7) - 7

	for i := start; i < start+7; i++ {
		totalInWeek += birdsPerDay[i]
	}

	return totalInWeek
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	// panic("Please implement the FixBirdCountLog() function")

	correctedSlice := birdsPerDay

	for i := 0; i < len(birdsPerDay); i++ {
		if i%2 == 0 {
			correctedSlice[i] = correctedSlice[i] + 1
		}
	}

	return correctedSlice
}
