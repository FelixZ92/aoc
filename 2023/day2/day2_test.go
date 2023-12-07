package main

import (
	"testing"

	"github.com/felixz92/aoc/2023/input"
)

func TestFromString(t *testing.T) {
	tc := []struct {
		name  string
		input string
		want  Game
	}{
		{
			name:  "example",
			input: "Game 21: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			want:  Game{Id: 21},
		},
	}
	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			got := FromString(tt.input)
			if got.Id != tt.want.Id {
				t.Errorf("FromString(%s) = %v; want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestRoundFromString(t *testing.T) {
	tc := []struct {
		name  string
		input string
		want  Round
	}{
		{
			name:  "example",
			input: "3 blue, 4 red",
			want:  Round{Blue: 3, Red: 4},
		},
		{
			name:  "example2",
			input: "1 red, 2 green, 6 blue",
			want:  Round{Red: 1, Green: 2, Blue: 6},
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			got := RoundFromString(tt.input)
			if got != tt.want {
				t.Errorf("RoundFromString(%s) = %v; want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestResult(t *testing.T) {
	str := `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`
	i := input.FromString(str)
	games := GamesFromInput(i)
	got := Result(games)
	want := 8
	if got != want {
		t.Errorf("Result(%v) = %v; want %v", games, got, want)
	}
}

func TestPart2(t *testing.T) {
	str := `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`
	i := input.FromString(str)
	games := GamesFromInput(i)
	got := Part2(games)
	want := 2286
	if got != want {
		t.Errorf("Result(%v) = %v; want %v", games, got, want)
	}
}
