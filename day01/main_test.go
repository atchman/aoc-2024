package main

import (
	"testing"
)

// test part 1
func TestPart1(t *testing.T) {
	left, right := parse("input.sample")
	total := part1(left, right)
	if 11 != total {
		t.Errorf("Part 1 total = %d; want 11", total)
	}
}

// test part 2
func TestPart2(t *testing.T) {
	left, right := parse("input.sample")
	score := part2(left, right)
	if 31 != score {
		t.Errorf("Part 2 score = %d, want 31", score)
	}
}
