package main

type Race struct {
	Time   int
	Record int
}

func Day6(races []Race) int {
	res := 1
	for _, r := range races {
		res *= r.Winnings()
	}

	return res
}

func (r *Race) Winnings() int {
	wins := 0
	for speed := 1; speed < r.Time+1; speed++ {
		travelTime := r.Time - speed
		distance := speed * travelTime
		if distance > r.Record {
			wins++
		}
	}

	return wins
}
