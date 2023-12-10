package main

import (
	"testing"

	"github.com/felixz92/aoc/2023/input"
)

func TestDay8(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name: "example1",
			input: `RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)`,
			expected: 2,
		},
		{
			name: "example2",
			input: `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`,
			expected: 6,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			in := input.FromString(test.input)
			actual := Day8(in)
			if actual != test.expected {
				t.Errorf("expected %v, got %v", test.expected, actual)
			}
		})
	}
}

func TestDay8Part2(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name: "example1",
			input: `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`,
			expected: 6,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			in := input.FromString(test.input)
			actual := Day8Part2(in)
			if actual != test.expected {
				t.Errorf("expected %v, got %v", test.expected, actual)
			}
		})
	}
}

func TestStartingNodes(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected map[string]Node
	}{
		{
			name: "example1",
			input: `11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`,
			expected: map[string]Node{
				"11A": {Name: "11A", Left: "11B", Right: "XXX"},
				"22A": {Name: "22A", Left: "22B", Right: "XXX"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			in := input.FromString(test.input)
			nodes := ParseNodes(in)
			actual := StartingNodes(nodes)
			if len(actual) != len(test.expected) {
				t.Errorf("expected %v, got %v", test.expected, actual)
			}
			for k, v := range test.expected {
				if actual[k] != v {
					t.Errorf("expected %v, got %v", test.expected, actual)
				}
			}
		})
	}
}
