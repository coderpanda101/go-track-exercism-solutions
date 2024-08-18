package raindrops

import "strconv"

func Convert(number int) string {
	dropSound := ""

	if number%3 == 0 || number%5 == 0 || number%7 == 0 {
		if number%3 == 0 {
			dropSound = "Pling"
		}
		if number%5 == 0 {
			dropSound = dropSound + "Plang"
		}
		if number%7 == 0 {
			dropSound = dropSound + "Plong"
		}
	} else {
		dropSound = strconv.Itoa(number)
	}

	return dropSound
}
