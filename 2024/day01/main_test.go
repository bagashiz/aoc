package main

import (
	"fmt"
	"testing"
)

var example = `3   4
4   3
2   5
1   3
3   9
3   3`

func TestPart1(t *testing.T) {
	testCases := []struct {
		input string
		want  int
	}{
		{input: example, want: 11},
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
		{input: example, want: 31},
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
