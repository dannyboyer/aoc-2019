package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// Location reprensting an X, Y coordinate
type Location struct {
	X, Y int
}

// CodeToInstruction converts a string to x, y delta
func CodeToInstruction(code string) Location {
	// man strings are low level in Go...

	direction := rune(code[0])
	value, err := strconv.Atoi(string(code[1]))
	if err != nil {
		log.Fatal(err)
	}

	if direction == 'R' {
		return Location{value, 0}
	} else if direction == 'L' {
		return Location{-value, 0}
	} else if direction == 'U' {
		return Location{0, value}
	} else if direction == 'D' {
		return Location{0, -value}
	} else {
		return Location{0, 0}
	}
}

// GetLocations parse a list of codes and return a list of Locations
func GetLocations(codeList []string) []Location {
	var locations []Location

	// For each comma separated value(string)
	for i, code := range codeList {
		// Convert the string to a location(delta) struct
		delta := CodeToInstruction(code)
		var next Location

		// Create a new location that is the sum of previous location and the delta
		if i == 0 {
			next = delta
		} else {
			previous := locations[i-1]
			next = Location{previous.X + delta.X, previous.Y + delta.Y}
		}
		// Add the new location to a list
		locations = append(locations, next)
	}

	return locations
}

func main() {
	// Parse the file and extract the two instructions
	// todo

	// For each wire
	wire1Codes := strings.Split("R8,U5,L5,D3", ",")
	wire2Codes := strings.Split("U7,R6,D4,L4", ",")

	// Parse both list and keep the location present in both
	wire1Locations := GetLocations(wire1Codes)
	wire2Locations := GetLocations(wire2Codes)

	fmt.Println(wire1Locations)
	fmt.Println(wire2Locations)

	// Extrapolate the full path from the Location(in between)
	// For each maching location
	// Calculate the manhatan distance from (0, 0)
	// Print the shortest distance
}
