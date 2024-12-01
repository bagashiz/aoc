package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/atotto/clipboard"
)

// https://adventofcode.com/2024/day/1

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
	lefts, rights, err := parseInput(input)
	if err != nil {
		return 0, fmt.Errorf("error processing input file: %v", err)
	}

	slices.Sort(lefts)
	slices.Sort(rights)

	var sumDiffs int

	for i := 0; i < len(lefts); i++ {
		diff := lefts[i] - rights[i]
		if diff < 0 {
			diff = -diff
		}
		sumDiffs += diff
	}

	return sumDiffs, nil
}

func part2(input string) (int, error) {
	lefts, rights, err := parseInput(input)
	if err != nil {
		return 0, fmt.Errorf("error processing input file: %v", err)
	}

	m := make(map[int]int)
	for _, left := range lefts {
		for _, right := range rights {
			if left == right {
				m[left] += right
			}
		}
	}

	var sumDiffs int
	for _, v := range m {
		sumDiffs += v
	}

	return sumDiffs, nil
}

// parseInput parses the input string into two sorted slices of integers.
func parseInput(input string) ([]int, []int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	lefts := make([]int, len(lines))
	rights := make([]int, len(lines))

	for i, line := range lines {
		parts := strings.Fields(line)
		if len(parts) != 2 {
			return nil, nil, fmt.Errorf("invalid line format")
		}

		left, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, nil, fmt.Errorf("error parsing left number")
		}

		right, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, nil, fmt.Errorf("error parsing left number")
		}

		lefts[i] = left
		rights[i] = right
	}

	return lefts, rights, nil
}
