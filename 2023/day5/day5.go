package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/felixz92/aoc/2023/input"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		panic("No input file")
	}

	i := input.FromFile(args[0])

	res := Day5(i)
	fmt.Println(res)

	res2 := Day5_Part2(i)
	fmt.Println(res2)
}

func Day5_Part2(lines []string) int {
	maps, seeds := parse(lines)
	fmt.Println(seeds)

	actualSeeds := make([]int, 0)
	for i := 0; i < len(seeds)/2; i++ {
		for j := 0; j < seeds[i*2+1]; j++ {
			actualSeeds = append(actualSeeds, seeds[i*2]+j)
		}
	}

	for _, m := range maps {
		for i, seed := range actualSeeds {
			actualSeeds[i] = MapSeed(m, seed)
		}
	}
	slices.Sort(actualSeeds)
	return actualSeeds[0]
}

func Day5(lines []string) int {
	maps, seeds := parse(lines)
	for _, m := range maps {
		for i, seed := range seeds {
			seeds[i] = MapSeed(m, seed)
		}
	}

	slices.Sort(seeds)
	return seeds[0]
}

type Mapping struct {
	Dest      int
	Source    int
	Range     int
	SourceMax int
	DestMax   int
}

func (m Mapping) String() string {
	return fmt.Sprintf("%d %d %d", m.Dest, m.Source, m.Range)
}

type Map struct {
	Name     string
	Mappings []Mapping
}

func (m Map) String() string {
	str := fmt.Sprintf("%s: ", m.Name)
	for _, m := range m.Mappings {
		str += fmt.Sprintf("%s;", m)
	}

	return str
}

func MapSeed(to Map, seed int) int {
	for _, m := range to.Mappings {
		if seed >= m.Source && seed <= m.SourceMax {
			diff := seed - m.Source
			return m.Dest + diff
		}
	}
	return seed
}

func parse(lines []string) ([]Map, []int) {
	seeds := make([]int, 0)
	maps := make([]Map, 0)
	var current Map
	for _, line := range lines {
		if strings.HasPrefix(line, "seeds: ") {
			trimmed := strings.TrimPrefix(line, "seeds: ")
			fields := strings.Fields(trimmed)
			for _, field := range fields {
				seed, err := strconv.Atoi(field)
				if err == nil {
					seeds = append(seeds, seed)
				}
			}
		}

		var d, s, r int
		n, _ := fmt.Sscanf(line, "%d %d %d", &d, &s, &r)
		if n == 3 {
			destMax := d + r - 1
			sourceMax := s + r - 1
			current.Mappings = append(current.Mappings, Mapping{d, s, r, sourceMax, destMax})
		}

		if strings.Contains(line, "map") {
			current.Name = line
		}

		if line == "" {
			if len(current.Mappings) > 0 {
				maps = append(maps, current)
			}
			current = Map{}
		}
	}

	return maps, seeds
}
