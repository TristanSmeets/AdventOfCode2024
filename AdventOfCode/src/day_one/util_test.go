package day_one

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadFromFile(t *testing.T) {
	expected := &Lists{
		Left:  []int{3, 4, 2, 1, 3, 3},
		Right: []int{4, 3, 5, 3, 9, 3},
	}

	result, err := ReadFromFile("input/test.txt")

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestCalculateDistances(t *testing.T) {
	expected := []int{2, 1, 0, 1, 2, 5}

	result := CalculateDistances(
		[]int{1, 2, 3, 3, 3, 4},
		[]int{3, 3, 3, 4, 5, 9},
	)

	assert.Equal(t, expected, result)
}
