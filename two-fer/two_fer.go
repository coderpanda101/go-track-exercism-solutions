package twofer

import "fmt"

// ShareWith function returns a string based upon the person's name and its availability
func ShareWith(name string) string {
	switch {
	case name != "":
		return fmt.Sprintf("One for %s, one for me.", name)
	default:
		return "One for you, one for me."
	}
}
