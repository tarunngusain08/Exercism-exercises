package partyrobot

import (
	"fmt"
	"math"
	"strconv"
)

// Welcome greets a person by name.
func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %v!", name)
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %v! You are now %v years old!", name, age)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	t := "00" + strconv.Itoa(table)
	return fmt.Sprintf("Welcome to my party, %v!\nYou have been assigned to table %v. Your table is %v, exactly %.1f meters from here.\nYou will be sitting next to %v.", name, t[len(t)-3:], direction, math.Round(distance*10)/10, neighbor)
}
