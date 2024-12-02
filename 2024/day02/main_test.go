package main

import (
	"fmt"
	"testing"
)

var example = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

func TestPart1(t *testing.T) {
	testCases := []struct {
		input string
		want  int
	}{
		{input: example, want: 2},
	}

	for i, tc := range testCases {
		index := fmt.Sprint(i)
		t.Run(index, func(t *testing.T) {
			got, err := part1(tc.input)
			if err != nil {
				t.Errorf("[case: %s] error running test: %v", index, err)
				return
			}

			if got != tc.want {
				t.Errorf("[case: %s] got %v, want %v", index, got, tc.want)
				return
			}
		})
	}
}

func TestPart2(t *testing.T) {
	testCases := []struct {
		input string
		want  int
	}{
		{input: example, want: 4},
	}

	for i, tc := range testCases {
		index := fmt.Sprint(i)
		t.Run(index, func(t *testing.T) {
			got, err := part2(tc.input)
			if err != nil {
				t.Errorf("[case: %s] error running test: %v", index, err)
				return
			}

			if got != tc.want {
				t.Errorf("[case: %s] got %v, want %v", index, got, tc.want)
				return
			}
		})
	}
}
