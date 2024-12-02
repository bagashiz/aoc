package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/atotto/clipboard"
)

// https://adventofcode.com/2024/day/2

//go:embed input.txt
var input string

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()

	if part != 1 && part != 2 {
		log.Fatalf("invalid part: %d, must be 1 or 2", part)
	}

	log.Println("running part", part)

	if len(input) == 0 {
		log.Fatal("input file is empty")
	}

	var ans int
	var err error

	switch part {
	case 1:
		ans, err = part1(input)
	case 2:
		ans, err = part2(input)
	}

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("answer: %v", ans)

	log.Println("writing answer to clipboard")

	if err := clipboard.WriteAll(strconv.Itoa(ans)); err != nil {
		log.Fatalf("error copy to clipboard: %v", err)
	}
}

func part1(input string) (int, error) {
	reports, err := parseInput(input)
	if err != nil {
		return 0, fmt.Errorf("error processing input file: %v", err)
	}

	var safe int

	for _, levels := range reports {
		if isSafe(levels) {
			safe++
		}
	}

	return safe, nil
}

func isSafe(levels []int) bool {
	increasing := true
	decreasing := true

	for i := 0; i < len(levels)-1; i++ {
		diff := levels[i] - levels[i+1]

		validDiff := diff != 0 && diff >= -3 && diff <= 3
		if !validDiff {
			return false
		}

		if diff < 0 {
			increasing = false
		}
		if diff > 0 {
			decreasing = false
		}
	}

	return increasing || decreasing
}

func part2(input string) (int, error) {
	reports, err := parseInput(input)
	if err != nil {
		return 0, fmt.Errorf("error processing input file: %v", err)
	}

	var safe int

	for _, levels := range reports {
		if isSafeDampened(levels) {
			safe++
		}
	}

	return safe, nil
}

func isSafeDampened(levels []int) bool {
	if isSafe(levels) {
		return true
	}

	for i := range levels {
		trimmed := make([]int, len(levels)-1)
		copy(trimmed, levels[:i])
		copy(trimmed[i:], levels[i+1:])

		if isSafe(trimmed) {
			return true
		}
	}

	return false
}

// parseInput parses the input string into slice of slices of integers.
func parseInput(input string) ([][]int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	reports := make([][]int, len(lines))

	for i, line := range lines {
		parts := strings.Fields(line)
		levels := make([]int, len(parts))

		for i, part := range parts {
			level, err := strconv.Atoi(part)
			if err != nil {
				return nil, fmt.Errorf("error convert string to number")
			}
			levels[i] = level
		}

		reports[i] = levels
	}

	return reports, nil
}
