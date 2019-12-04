package main

import (
	"strings"
	"testing"
)

func TestCodeToInstruction(t *testing.T) {
	tables := []struct {
		code     string
		location Location
	}{
		{"R8", Location{8, 0}},
		{"U5", Location{0, 5}},
		{"L5", Location{-5, 0}},
		{"D3", Location{0, -3}},
		{"U7", Location{0, 7}},
		{"R6", Location{6, 0}},
		{"D4", Location{0, -4}},
		{"L4", Location{-4, 0}},
	}

	for _, table := range tables {
		got := CodeToInstruction(table.code)
		expected := table.location
		if got != expected {
			t.Errorf("ConvertToPoint() = got %v want %v", got, expected)
		}
	}
}

func TestGetLocations(t *testing.T) {
	tables := []struct {
		codeList  []string
		locations []Location
	}{
		{
			strings.Split("R8,U5,L5,D3", ","),
			[]Location{Location{8, 0}, Location{8, 0}, Location{8, 0}, Location{8, 0}},
		},
	}

	for _, table := range tables {
		got := GetLocations(table.codeList)
		expected := table.locations
		if got[0] != expected[0] {
			t.Errorf("GetLocations() = got %v want %v", got, expected)
		}
	}
}
