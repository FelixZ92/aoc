package main

import "testing"

func TestMapSeed(t *testing.T) {
	tests := []struct {
		name     string
		to       Map
		seed     int
		expected int
	}{
		{
			name: "example-seed-to-soil-79",
			to: Map{Mappings: []Mapping{
				{Source: 98, Dest: 50, Range: 2, SourceMax: 99, DestMax: 51},
				{Source: 50, Dest: 52, Range: 48, SourceMax: 97, DestMax: 99},
			}},
			seed:     79,
			expected: 81,
		},
		{
			name: "example-seed-to-soil-14",
			to: Map{Mappings: []Mapping{
				{Source: 98, Dest: 50, Range: 2, SourceMax: 99, DestMax: 51},
				{Source: 50, Dest: 52, Range: 48, SourceMax: 97, DestMax: 99},
			}},
			seed:     14,
			expected: 14,
		},
		{
			name: "example-seed-to-soil-55",
			to: Map{Mappings: []Mapping{
				{Source: 98, Dest: 50, Range: 2, SourceMax: 99, DestMax: 51},
				{Source: 50, Dest: 52, Range: 48, SourceMax: 97, DestMax: 99},
			}},
			seed:     55,
			expected: 57,
		},
		{
			name: "example-seed-to-soil-13",
			to: Map{Mappings: []Mapping{
				{Source: 98, Dest: 50, Range: 2, SourceMax: 99, DestMax: 51},
				{Source: 50, Dest: 52, Range: 48, SourceMax: 97, DestMax: 99},
			}},
			seed:     13,
			expected: 13,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := MapSeed(test.to, test.seed)
			if res != test.expected {
				t.Errorf("Expected %d, got %d", test.expected, res)
			}
		})
	}
}
