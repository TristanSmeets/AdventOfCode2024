package day_one

import (
	"slices"
)

func DayOne(path string) (int, error) {
	lists, err := ReadFromFile(path)
	if err != nil {
		return 0, err
	}

	left := lists.Left
	right := lists.Right

	slices.Sort(left)
	slices.Sort(right)

	distances := CalculateDistances(left, right)

	totalDistance := 0
	for _, d := range distances {
		totalDistance += d
	}

	return totalDistance, nil
}
