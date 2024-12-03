package main

import (
	"fmt"
	"testing"
)

func TestPart1(t *testing.T) {
	testCases := []struct {
		input string
		want  int
	}{
		{input: "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))", want: 161},
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
		{input: "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))", want: 48},
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
