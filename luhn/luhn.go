package luhn

import (
	"regexp"
	"strconv"
	"strings"
)

func Valid(id string) bool {
	// Stripping id of any whitepaces
	id = strings.ReplaceAll(id, " ", "")

	// Checking if a string contains only numbers
	re := regexp.MustCompile(`^[0-9]+$`)
	if len(id) <= 1 || !re.MatchString(id) {
		return false
	}

	// Creating an slice of numbers in string format
	idSlc := strings.Split(id, "")

	var ids []int
	// Converting each string number into int
	for i := 0; i < len(idSlc); i++ {
		tmp, err := strconv.Atoi(idSlc[i])
		if err != nil {
			return false
		}
		ids = append(ids, tmp)
	}

	// Doubling every second digit from right
	for i := len(ids) - 2; i >= 0; i -= 2 {
		tmp := ids[i]
		tmp *= 2
		if tmp > 9 {
			tmp -= 9
		}
		ids[i] = tmp
	}

	// Calculating sum of the numbers
	sum := 0
	for i := 0; i < len(ids); i++ {
		sum += ids[i]
	}

	return sum%10 == 0
}
