package day_two

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDayTwo(t *testing.T) {
	result, err := DayTwo("test_input.txt")

	assert.NoError(t, err)
	assert.Equal(t, 2, result)
}

func TestIsReportSafe(t *testing.T) {
	type scenario struct {
		name   string
		input  *Report
		result bool
	}

	scenarios := []scenario{
		{
			name:   "All levels decreasing by 1 or 2",
			input:  &Report{Levels: []int{7, 6, 4, 2, 1, 0}},
			result: true,
		},
		{
			name:   "Has increase of 5",
			input:  &Report{Levels: []int{1, 2, 7, 8, 9}},
			result: false,
		},
		{
			name:   "Has decrease of 4",
			input:  &Report{Levels: []int{9, 7, 6, 2, 1}},
			result: false,
		},
		{
			name:   "Has increase and decrease",
			input:  &Report{Levels: []int{1, 3, 2, 4, 5}},
			result: false,
		},
		{
			name:   "Has no increase or decrease",
			input:  &Report{Levels: []int{8, 6, 4, 4, 1}},
			result: false,
		},
		{
			name:   "All levels are increasing by 1,2 or 3",
			input:  &Report{Levels: []int{1, 3, 6, 7, 9}},
			result: true,
		},
		{
			name:   "challenge report",
			input:  &Report{Levels: []int{41, 44, 45, 48, 49, 50, 50}},
			result: false,
		},
	}

	for _, test := range scenarios {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.result, reportIsSafe(test.input.Levels))
		})
	}
}
