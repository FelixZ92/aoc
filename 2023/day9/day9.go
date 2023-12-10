package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/felixz92/aoc/2023/input"
)

func main() {
	i := input.FromStdin()

	res := Day9(i)
	fmt.Println(res)

	res2 := Day9Part2(i)
	fmt.Println(res2)
}

func Day9(lines []string) int {
	histories := make([][]int, len(lines))
	sum := 0
	sequences := make([][][]int, len(histories))
	for i := range lines {
		histories[i] = HistoryFromLine(lines[i])
		sequences[i] = BuildSequences(histories[i])
		val := Extrapolate(sequences[i])
		sum += val
	}

	return sum
}

func Day9Part2(lines []string) int {
	histories := make([][]int, len(lines))
	sum := 0
	sequences := make([][][]int, len(histories))
	for i := range lines {
		histories[i] = HistoryFromLine(lines[i])
		sequences[i] = BuildSequences(histories[i])
		val := ExtrapolateBackwards(sequences[i])
		sum += val
	}

	return sum
}

func BuildSequences(a []int) [][]int {
	sequences := make([][]int, 0)
	for {
		sequences = append(sequences, a)
		if onlyZeros(a) {
			break
		}
		a = NextSequence(a)
	}

	return sequences
}

func onlyZeros(a []int) bool {
	if len(a) == 0 {
		return false
	}
	for i := range a {
		if a[i] != 0 {
			return false
		}
	}
	return true
}

func NextSequence(a []int) []int {
	next := make([]int, len(a)-1)

	for i := range next {
		next[i] = a[i+1] - a[i]
	}

	return next
}

func Extrapolate(sequences [][]int) int {
	for i := len(sequences) - 1; i > 0; i-- {
		if onlyZeros(sequences[i]) {
			sequences[i] = append(sequences[i], 0)
		}
		nextValue := ExtrapolateNext(sequences[i], sequences[i-1])
		sequences[i-1] = append(sequences[i-1], nextValue)
	}

	return sequences[0][len(sequences[0])-1]
}

func ExtrapolateBackwards(sequences [][]int) int {
	for i := len(sequences) - 1; i > 0; i-- {
		if onlyZeros(sequences[i]) {
			sequences[i] = append(sequences[i], 0)
		}
		nextValue := ExtrapolatePrevious(sequences[i], sequences[i-1])
		sequences[i-1] = append([]int{nextValue}, sequences[i-1]...)
	}

	return sequences[0][0]
}

func ExtrapolatePrevious(a []int, n []int) int {
	return n[0] - a[0]
}

func ExtrapolateNext(a []int, n []int) int {
	return n[len(n)-1] + a[len(a)-1]
}

func HistoryFromLine(s string) []int {
	split := strings.Fields(s)
	nums := make([]int, len(split))
	for i := range split {
		nums[i], _ = strconv.Atoi(split[i])
	}

	return nums
}
