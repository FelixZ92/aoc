package main

import (
	"fmt"
	"testing"
)

func TestDay6(t *testing.T) {
	tests := []struct {
		name     string
		races    []Race
		expected int
	}{
		{
			name: "example",
			races: []Race{
				{Time: 7, Record: 9},
				{Time: 15, Record: 40},
				{Time: 30, Record: 200},
			},
			expected: 288,
		},
		{
			name: "input",
			races: []Race{
				{Time: 51, Record: 222},
				{Time: 92, Record: 2031},
				{Time: 68, Record: 1126},
				{Time: 90, Record: 1225},
			},
			expected: 500346,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := Day6(test.races)
			if res != test.expected {
				t.Errorf("Expected %d, got %d", test.expected, res)
			}
		})
	}
}

func TestWinnings(t *testing.T) {
	tests := []struct {
		race     Race
		expected int
	}{
		{race: Race{Time: 7, Record: 9}, expected: 4},
		{race: Race{Time: 15, Record: 40}, expected: 8},
		{race: Race{Time: 30, Record: 200}, expected: 9},
		{race: Race{Time: 71530, Record: 940200}, expected: 71503},
		// Part 2
		{race: Race{Time: 51926890, Record: 222203111261225}, expected: 42515755},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%d-%d", test.race.Time, test.race.Record), func(t *testing.T) {
			res := test.race.Winnings()
			if res != test.expected {
				t.Errorf("Expected %d, got %d", test.expected, res)
			}
		})
	}
}
