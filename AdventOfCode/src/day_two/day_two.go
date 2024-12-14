package day_two

import (
	"fmt"
	"math"
)

type Report struct {
	Levels []int
}

func DayTwo(path string) (int, error) {

	reports, err := ReadInputFromFile(path)
	if err != nil {
		return 0, err
	}

	var amountOfSafeReports int

	for i, report := range reports {
		println(fmt.Sprintf("Report#%d: %t", i, reportIsSafe(report.Levels)))

		if reportIsSafe(report.Levels) {
			amountOfSafeReports++
		}
	}

	return amountOfSafeReports, nil
}

func reportIsSafe(levels []int) bool {
	isIncreasing := (levels[0] - levels[1]) > 0

	// Don't have to check the last entry
	for i := 0; i < len(levels)-1; i++ {
		difference := levels[i] - levels[i+1]

		// Check if also increasing/decreasing
		if (difference > 0) != isIncreasing {
			return false
		}

		distance := int(math.Abs(float64(difference)))
		if distance == 0 || distance > 3 {
			return false
		}
	}

	return true
}
