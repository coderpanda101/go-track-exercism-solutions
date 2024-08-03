package thefarm

import (
	"errors"
	"fmt"
)

// 'DivideFood' function
func DivideFood(fc FodderCalculator, numOfCows int) (float64, error) {
	totalFodderQty, err1 := fc.FodderAmount(numOfCows)
	if err1 != nil {
		return 0.0, err1
	}

	fatteningFactor, err2 := fc.FatteningFactor()
	if err2 != nil {
		return 0.0, err2
	}

	eachFodderQty := (totalFodderQty * fatteningFactor) / float64(numOfCows)

	return eachFodderQty, nil
}

// 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fc FodderCalculator, numOfCows int) (float64, error) {
	if numOfCows > 0 {
		return DivideFood(fc, numOfCows)
	}

	return 0.0, errors.New("invalid number of cows")
}

type InvalidCowsError struct {
	numOfCows    int
	errorMessage string
}

func (ice *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", ice.numOfCows, ice.errorMessage)
}

// 'ValidateNumberOfCows' function
func ValidateNumberOfCows(numOfCows int) error {
	switch {
	case numOfCows < 0:
		return &InvalidCowsError{
			errorMessage: "there are no negative cows",
			numOfCows:    numOfCows,
		}
	case numOfCows == 0:
		return &InvalidCowsError{
			errorMessage: "no cows don't need food",
			numOfCows:    numOfCows,
		}
	default:
		return nil
	}
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
