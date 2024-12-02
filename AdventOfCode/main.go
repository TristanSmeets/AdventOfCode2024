package main

import (
	"AdventOfCode/src/day_one"
	"fmt"
)

func main() {
	totalDistance, err := day_one.DayOne("challenge_inputs/challenge.txt")
	if err != nil {
		println(err)
	}
	println(fmt.Sprintf("Day 1 - Total distance between lists: %d", totalDistance))
}
