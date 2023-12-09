package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/felixz92/aoc/2023/input"
)

func main() {
	i := input.FromStdin()
	var runes [][]rune = make([][]rune, len(i))
	for i, line := range i {
		runes[i] = []rune(line)
	}

	res := Day4(runes)
	fmt.Println(res)

	res2 := Day4_Part2(i)
	fmt.Println(res2)
}

func Day4(runes [][]rune) int {
	cards := make([]Card, len(runes))
	for i, r := range runes {
		cards[i] = FromString(string(r))
	}

	sum := 0
	for _, c := range cards {
		sum += c.Worth()
	}

	return sum
}

func Day4_Part2(lines []string) int {
	cards := make([]Card, len(lines))
	for i, r := range lines {
		cards[i] = FromString(string(r))
	}

	counts := make([]int, len(cards))
	for i := 0; i < len(cards); i++ {
		counts[i] = 1
	}

	sum := 0

	for i, c := range cards {
		matches := c.Matches()
		for j := 1; j <= matches; j++ {
			counts[i+j] += counts[i]
		}

		sum += counts[i]
	}

	return sum
}

type Card struct {
	Id             int
	WinningNumbers []string
	DrawnNumbers   []string
}

func (c *Card) Worth() int {
	matches := c.Matches()

	if matches == 0 {
		return 0
	}

	return 1 << (matches - 1)
}

func (c *Card) Matches() int {
	matches := 0
	for _, n := range c.DrawnNumbers {
		if slices.Contains(c.WinningNumbers, n) {
			matches++
		}
	}

	return matches
}

func contains(arr []int, n int) bool {
	for _, v := range arr {
		if v == n {
			return true
		}
	}
	return false
}

func (c *Card) String() string {
	return fmt.Sprintf("Card %d: %v | %v", c.Id, c.WinningNumbers, c.DrawnNumbers)
}

func FromString(s string) Card {
	firstSplit := strings.Split(s, ":")
	idSplit := strings.Split(firstSplit[0], " ")
	id, _ := strconv.Atoi(idSplit[1])
	secondSplit := strings.Split(firstSplit[1], "|")
	winningNumbers := strings.Fields(secondSplit[0])
	drawnNumbers := strings.Fields(secondSplit[1])

	return Card{id, winningNumbers, drawnNumbers}
}
