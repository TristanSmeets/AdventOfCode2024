package day_one

import (
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Lists struct {
	Left  []int
	Right []int
}

func ReadFromFile(path string) (*Lists, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	stringData := string(data)

	columns := strings.Split(stringData, "\r\n")

	leftList := make([]int, 0)
	rightList := make([]int, 0)

	for _, column := range columns {
		rows := strings.Split(column, "   ")

		l, err := strconv.Atoi(rows[0])
		if err != nil {
			return nil, err
		}
		leftList = append(leftList, l)

		r, err := strconv.Atoi(rows[1])
		if err != nil {
			return nil, err
		}
		rightList = append(rightList, r)
	}

	return &Lists{Left: leftList, Right: rightList}, nil
}

func CalculateDistances(left []int, right []int) []int {
	distances := make([]int, 0)

	for i := 0; i < len(left); i++ {

		d := math.Abs(float64(left[i]) - float64(right[i]))

		distances = append(distances, int(d))
	}
	return distances
}
