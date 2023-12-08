package main

import (
	"fmt"
	"strconv"
	"unicode"

	"github.com/felixz92/aoc/2023/input"
)

func main() {
	i := input.FromStdin()
	fmt.Println(i)
	var runes [][]rune = make([][]rune, len(i))
	for i, line := range i {
		runes[i] = []rune(line)
	}

	res := Day3(runes)
	fmt.Println(res)
	res2 := Day3_2(runes)
	fmt.Println(res2)
}

func Day3(runes [][]rune) int {
	sum := 0
	for i, r := range runes {
		offset := 0
		n := 0
		var beforeIndex int

		for n != -1 {
			if offset >= len(r) {
				break
			}
			n, beforeIndex, offset = NextNumber(r, offset)
			if n != -1 {
				fmt.Printf("number = %d ,beforeIndex = %d , index = %d\n", n, beforeIndex, offset)
				// directly adjacent
				if (beforeIndex > -1 && r[beforeIndex] != '.') || offset < len(r) && r[offset] != '.' {
					sum += n
				} else if (i > 0 && checkLine(runes[i-1], beforeIndex, offset)) || (i < len(runes)-1 && checkLine(runes[i+1], beforeIndex, offset)) {
					sum += n
				}
			}
		}
	}

	return sum
}

func Day3_2(runes [][]rune) int {
	sum := 0
	for y, r := range runes {
		for x, c := range r {
			if c == '*' {
				isGear, neighbor1, neighbor2 := IsGear(runes, x, y)
				if isGear {
					sum += neighbor1 * neighbor2
				}
			}
		}
	}

	return sum
}

func IsGear(runes [][]rune, x, y int) (bool, int, int) {
	// has direct neighbors
	// left
	neighbors := make([]int, 0)

	hasLeft, leftNeighbor := HasLeftNeighbor(runes[y], x)
	if hasLeft {
		neighbors = append(neighbors, leftNeighbor)
	}

	hasRight, rightNeighbor := HasRightNeighbor(runes[y], x)
	if hasRight {
		neighbors = append(neighbors, rightNeighbor)
	}

	// top
	if y > 0 {
		neighbors = append(neighbors, TopOrAboveNeighbors(runes[y-1], x)...)
	}

	// above
	if y < len(runes)-1 {
		neighbors = append(neighbors, TopOrAboveNeighbors(runes[y+1], x)...)
	}

	if len(neighbors) == 2 {
		return true, neighbors[0], neighbors[1]
	}

	return false, 0, 0
}

func TopOrAboveNeighbors(runes []rune, gearIndex int) []int {
	neighbors := make([]int, 0)
	hasLeft, leftNeighbor := HasLeftNeighbor(runes, gearIndex)
	if hasLeft {
		neighbors = append(neighbors, leftNeighbor)
	}

	hasRight, rightNeighbor := HasRightNeighbor(runes, gearIndex)
	if hasRight {
		neighbors = append(neighbors, rightNeighbor)
	}

	if hasDirect, directNeighbor := HasTopOrAboveNeighbor(runes, gearIndex); hasDirect {
		neighbors = append(neighbors, directNeighbor)
	}

	return neighbors
}

func HasTopOrAboveNeighbor(r []rune, gearIndex int) (bool, int) {
	offset := 0
	n := -1
	beforeIndex := 0
	for offset != -1 {
		n, beforeIndex, offset = NextNumber(r, offset)
		if beforeIndex < gearIndex && offset > gearIndex {
			return true, n
		}
	}

	return false, -1
}

func HasLeftNeighbor(r []rune, gearIndex int) (bool, int) {
	if gearIndex == 0 {
		return false, -1
	}

	offset := 0
	n := -1
	for offset != -1 && offset < gearIndex {
		n, _, offset = NextNumber(r, offset)
		if n != -1 && offset == gearIndex {
			return true, n
		}
	}

	return false, -1
}

func HasRightNeighbor(r []rune, gearIndex int) (bool, int) {
	if gearIndex == len(r)-1 {
		return false, -1
	}

	n, beforeIndex, _ := NextNumber(r, gearIndex)
	if n != -1 && beforeIndex == gearIndex {
		return true, n
	}

	return false, -1
}

func checkLine(r []rune, beforeIndex, offset int) bool {
	if beforeIndex == -1 {
		beforeIndex = 0
	}
	if offset == len(r) {
		offset = len(r) - 1
	}

	for i := beforeIndex; i <= offset; i++ {
		if r[i] != '.' {
			return true
		}
	}

	return false
}

func NextNumber(runes []rune, offset int) (int, int, int) {
	if offset >= len(runes) {
		return -1, -1, -1
	}

	var firstDigitIndex int
	var n []rune
	for i := offset; i < len(runes); i++ {
		if unicode.IsDigit(runes[i]) {
			firstDigitIndex = i
			n = append(n, runes[i])
			if i < len(runes)-1 {
				j := i + 1
				for unicode.IsDigit(runes[j]) {
					n = append(n, runes[j])
					j++
					if j >= len(runes) {
						break
					}

				}

			}
			num, err := strconv.Atoi(string(n))
			if err != nil {
				panic(err)
			}
			return num, firstDigitIndex - 1, i + len(n)
		}
	}

	return -1, -1, -1
}
