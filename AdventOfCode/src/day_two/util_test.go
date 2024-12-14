package day_two

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadInputFromFile(t *testing.T) {
	path := "test_input.txt"

	result, err := ReadInputFromFile(path)

	assert.NoError(t, err)
	assert.Equal(t, 6, len(result))
	for _, report := range result {
		assert.Equal(t, 5, len(report.Levels))
	}
}
