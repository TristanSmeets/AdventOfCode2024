package day_one

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_dayOne(t *testing.T) {

	result, err := DayOne("input/test.txt")

	assert.NoError(t, err)
	assert.Equal(t, 11, result)
}
