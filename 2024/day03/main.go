package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/atotto/clipboard"
)

// https://adventofcode.com/2024/day/3

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
	lines := strings.Split(strings.TrimSpace(input), "\n")

	// captures "mul([3 digits], [3 digits])"
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	total := 0

	for _, line := range lines {
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			res, err := mul(match[1], match[2])
			if err != nil {
				return 0, fmt.Errorf("error multiplying: %v", err)
			}
			total += res
		}

	}

	return total, nil
}

func part2(input string) (int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	// captures "mul([3 digits], [3 digits])", "do()", and "don't()"
	re := regexp.MustCompile(`(?:mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\))`)

	total := 0
	do := true

	for _, line := range lines {
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			switch match[0] {
			case "do()":
				do = true
				continue

			case "don't()":
				do = false
				continue

			default:
				if do {
					res, err := mul(match[1], match[2])
					if err != nil {
						return 0, fmt.Errorf("error multiplying: %v", err)
					}
					total += res
				}
			}
		}
	}

	return total, nil
}

// mul convert 2 strings to int then multiply them.
func mul(a, b string) (int, error) {
	x, err := strconv.Atoi(a)
	if err != nil {
		return 0, fmt.Errorf("failed to convert string to int: %v", err)
	}
	y, err := strconv.Atoi(b)
	if err != nil {
		return 0, fmt.Errorf("failed to convert string to int: %v", err)
	}

	return x * y, nil
}
