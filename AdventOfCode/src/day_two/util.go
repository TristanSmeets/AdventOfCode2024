package day_two

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadInputFromFile(path string) ([]*Report, error) {
	bytes, err := os.ReadFile(path)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	data := string(bytes)

	columns := strings.Split(data, "\r\n")

	reports := make([]*Report, 0)

	for _, row := range columns {
		report := &Report{}

		levels := strings.Split(row, " ")

		for _, level := range levels {
			l, err := strconv.Atoi(level)
			if err != nil {
				return nil, err
			}
			report.Levels = append(report.Levels, l)
		}

		reports = append(reports, report)
	}

	return reports, nil
}
