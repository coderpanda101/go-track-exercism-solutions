package partyrobot

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	// panic("Please implement the Welcome function")
	return fmt.Sprintf("Welcome to my party, %s!", name)
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	// panic("Please implement the HappyBirthday function")
	return fmt.Sprintf("Happy birthday %s! You are now %v years old!", name, age)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	// panic("Please implement the AssignTable function")
	return fmt.Sprintf("Welcome to my party, %v!\nYou have been assigned to table %.3d. Your table is %v, exactly %.1f meters from here.\nYou will be sitting next to %v.", name, table, direction, distance, neighbor)
}
