package main

import (
	"AdventOfCode/src/day_one"
	"AdventOfCode/src/day_two"
	"fmt"
)

func main() {
	totalDistance, err := day_one.DayOne("challenge_inputs/challenge.txt")
	if err != nil {
		println(err)
	}
	println(fmt.Sprintf("Day 1 - Total distance between lists: %d", totalDistance))

	safeReports, err := day_two.DayTwo("challenge_inputs/day_two.txt")
	if err != nil {
		println(err)
	}
	println(fmt.Sprintf("Day 2 - Safe reports: %d", safeReports))
}
