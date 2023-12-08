package main

import (
	"testing"

	"github.com/felixz92/aoc/2023/input"
)

func TestDay3(t *testing.T) {
	str := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

	i := input.FromString(str)
	runes := make([][]rune, len(i))
	for i, line := range i {
		runes[i] = []rune(line)
	}

	got := Day3(runes)
	want := 4361
	if got != want {
		t.Errorf("Day3() = %v, want %v", got, want)
	}
}

func TestDay3_2(t *testing.T) {
	str := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

	i := input.FromString(str)
	runes := make([][]rune, len(i))
	for i, line := range i {
		runes[i] = []rune(line)
	}

	got := Day3_2(runes)
	want := 467835
	if got != want {
		t.Errorf("Day3() = %v, want %v", got, want)
	}
}

func TestNextNumber(t *testing.T) {
	tests := []struct {
		name        string
		in          string
		beforeIndex int
		offset      int
		wantedNum   int
		wantedIndex int
	}{
		{"simple", "123", -1, 0, 123, 3},
		{"first line", "467..114..", -1, 0, 467, 3},
		{"first line 2", "467..114..", 4, 3, 114, 8},
		{"second line", "...*......", -1, 0, -1, -1},
		{"third line", "..35..633.", 1, 0, 35, 4},
		{"third line", "..35..633.", 5, 4, 633, 9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNum, gotBeforeIndex, gotOffset := NextNumber([]rune(tt.in), tt.offset)
			if gotNum != tt.wantedNum {
				t.Errorf("NextNumber() gotNum = %v, wantedNum %v", gotNum, tt.wantedNum)
			}
			if gotOffset != tt.wantedIndex {
				t.Errorf("NextNumber() gotIndex = %v, wantedIndex %v", gotOffset, tt.wantedIndex)
			}

			if gotBeforeIndex != tt.beforeIndex {
				t.Errorf("NextNumber() gotBeforeIndex = %v, wantedBeforeIndex %v", gotBeforeIndex, tt.beforeIndex)
			}
		})
	}
}

func TestHasLeftNeighbor(t *testing.T) {
	tests := []struct {
		name            string
		in              string
		index           int
		hasLeftNeighbor bool
		leftNeighbor    int
	}{
		{"hasNeigbor", ".637*....961.", 4, true, 637},
		{"hasNoNeigbor", ".637.*....961.", 5, false, -1},
		{"nothing at all", ".....*...#..", 5, false, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHasLeftNeighbor, gotLeftNeighbor := HasLeftNeighbor([]rune(tt.in), tt.index)
			if gotHasLeftNeighbor != tt.hasLeftNeighbor {
				t.Errorf("HasLeftNeighbor() gotHasLeftNeighbor = %v, wanted %v", gotHasLeftNeighbor, tt.hasLeftNeighbor)
			}
			if gotLeftNeighbor != tt.leftNeighbor {
				t.Errorf("HasLeftNeighbor() gotLeftNeighbor = %v, wanted %v", gotLeftNeighbor, tt.leftNeighbor)
			}
		})
	}
}

func TestHasRightNeighbor(t *testing.T) {
	tests := []struct {
		name             string
		in               string
		index            int
		hasRightNeighbor bool
		rightNeighbor    int
	}{
		{"hasNeigbor", ".637*961...", 4, true, 961},
		{"hasNoNeigbor", ".637*....961.", 4, false, -1},
		{"nothing at all", ".....*...#..", 5, false, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHasRightNeighbor, gotRightNeighbor := HasRightNeighbor([]rune(tt.in), tt.index)
			if gotHasRightNeighbor != tt.hasRightNeighbor {
				t.Errorf("HasRightNeighbor() gotHasRightNeighbor = %v, wanted %v", gotHasRightNeighbor, tt.hasRightNeighbor)
			}
			if gotRightNeighbor != tt.rightNeighbor {
				t.Errorf("HasRightNeighbor() gotRightNeighbor = %v, wanted %v", gotRightNeighbor, tt.rightNeighbor)
			}
		})
	}
}

func TestHasTopOrAboveNeighbor(t *testing.T) {
	tests := []struct {
		name                  string
		in                    string
		index                 int
		hasTopOrAboveNeighbor bool
		topOrAboveNeighbor    int
	}{
		{"hasLeftNeigbor", ".637......", 3, true, 637},
		{"hasRightNeigbor", ".637..961.", 6, true, 961},
		{"hasNoNeigbor", ".637....961.", 5, false, -1},
		{"nothing at all", ".....*...#..", 5, false, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHasTopOrAboveNeighbor, gotTopOrAboveNeighbor := HasTopOrAboveNeighbor([]rune(tt.in), tt.index)
			if gotHasTopOrAboveNeighbor != tt.hasTopOrAboveNeighbor {
				t.Errorf("HasTopOrAboveNeighbor() gotHasTopOrAboveNeighbor = %v, wanted %v", gotHasTopOrAboveNeighbor, tt.hasTopOrAboveNeighbor)
			}
			if gotTopOrAboveNeighbor != tt.topOrAboveNeighbor {
				t.Errorf("HasTopOrAboveNeighbor() gotTopOrAboveNeighbor = %v, wanted %v", gotTopOrAboveNeighbor, tt.topOrAboveNeighbor)
			}
		})
	}
}
