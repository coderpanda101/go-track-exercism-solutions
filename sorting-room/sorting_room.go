package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %.1f", float64(nb.Number()))
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	fancyNum, _ := strconv.Atoi((fnb.Value()))
	switch fnb.(type) {
	case FancyNumber:
		return fancyNum
	default:
		return 0
	}
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	fancyNum, _ := strconv.Atoi(fnb.Value())
	switch fnb.(type) {
	case FancyNumber:
		return fmt.Sprintf("This is a fancy box containing the number %.1f", float64(fancyNum))
	default:
		return "This is a fancy box containing the number 0.0"
	}
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	switch b := i.(type) {
	case int:
		return DescribeNumber(float64(b))
	case float64:
		return DescribeNumber(b)
	case NumberBox:
		return DescribeNumberBox(b)
	case FancyNumberBox:
		return DescribeFancyNumberBox(b)
	default:
		return "Return to sender"
	}
}
