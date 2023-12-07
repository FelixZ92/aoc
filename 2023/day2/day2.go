package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/felixz92/aoc/2023/input"
)

func main() {
	i := input.FromStdin()
	games := GamesFromInput(i)
	fmt.Println(Result(games))
	fmt.Println(Part2(games))
}

func Result(games []Game) int {
	sum := 0
	for _, g := range games {
		if g.IsPossible() {
			sum += g.Id
		}
	}
	return sum
}

func Part2(games []Game) int {
	power := 0
	for _, g := range games {
		power += g.Power()
	}
	return power
}

func (g *Game) String() string {
	return fmt.Sprintf("Game %d", g.Id)
}

func (g *Game) Power() int {
	return g.maxBlue() * g.maxGreen() * g.maxRed()
}

func (g *Game) maxRed() int {
	max := 0
	for _, r := range g.Rounds {
		if r.Red > max {
			max = r.Red
		}
	}
	return max
}

func (g *Game) maxBlue() int {
	max := 0
	for _, r := range g.Rounds {
		if r.Blue > max {
			max = r.Blue
		}
	}
	return max
}

func (g *Game) maxGreen() int {
	max := 0
	for _, r := range g.Rounds {
		if r.Green > max {
			max = r.Green
		}
	}
	return max
}

type Round struct {
	Blue  int
	Green int
	Red   int
}

func (r *Round) String() string {
	return fmt.Sprintf("Round: %d blue, %d green, %d red", r.Blue, r.Green, r.Red)
}

type Game struct {
	Rounds []Round
	Id     int
}

func (g *Game) IsPossible() bool {
	for _, r := range g.Rounds {
		if !r.IsPossible() {
			return false
		}
	}
	return true
}

func GamesFromInput(s []string) []Game {
	games := make([]Game, len(s))
	for i, line := range s {
		games[i] = FromString(line)
	}
	return games
}

func FromString(s string) Game {
	extracted := strings.Split(s, ":")
	idStr := strings.Split(extracted[0], " ")[1]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		panic(err)
	}

	r := strings.Split(extracted[1], ";")
	rounds := make([]Round, len(r))
	for i, round := range r {
		rounds[i] = RoundFromString(round)
	}

	return Game{Id: id, Rounds: rounds}
}

func (r *Round) IsPossible() bool {
	return r.Blue <= 14 && r.Green <= 13 && r.Red <= 12
}

func RoundFromString(s string) Round {
	cubes := strings.Split(s, ",")
	var red, blue, green int
	for _, c := range cubes {
		c = strings.TrimSpace(c)
		if strings.Contains(c, "red") {
			var err error
			red, err = strconv.Atoi(strings.Split(c, " ")[0])
			if err != nil {
				panic(err)
			}
		}
		if strings.Contains(c, "blue") {
			var err error
			blue, err = strconv.Atoi(strings.Split(c, " ")[0])
			if err != nil {
				panic(err)
			}
		}
		if strings.Contains(c, "green") {
			var err error
			green, err = strconv.Atoi(strings.Split(c, " ")[0])
			if err != nil {
				panic(err)
			}
		}
	}
	return Round{Blue: blue, Green: green, Red: red}
}
