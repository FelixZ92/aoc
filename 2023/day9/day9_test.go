package main

import "testing"

func TestNextSequence(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "example1",
			input:    []int{0, 3, 6, 9, 12, 15},
			expected: []int{3, 3, 3, 3, 3},
		},
		{
			name:     "example2",
			input:    []int{3, 3, 3, 3, 3},
			expected: []int{0, 0, 0, 0},
		},
		{
			name:     "example3",
			input:    []int{1, 3, 6, 10, 15, 21},
			expected: []int{2, 3, 4, 5, 6},
		},
		{
			name:     "example4",
			input:    []int{2, 3, 4, 5, 6},
			expected: []int{1, 1, 1, 1},
		},
		{
			name:     "example5",
			input:    []int{2, 3, -4, 5, 6},
			expected: []int{1, -7, 9, 1},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := NextSequence(test.input)
			if len(actual) != len(test.expected) {
				t.Errorf("expected %v, got %v", test.expected, actual)
			}
			for i := range actual {
				if actual[i] != test.expected[i] {
					t.Errorf("expected %v, got %v", test.expected, actual)
				}
			}
		})
	}
}
