package main

import (
	"testing"
)

// test part 1
func TestPart1(t *testing.T) {
	safe := part1(parse("input.sample"))
	if 2 != safe {
		t.Errorf("Part 1 safe = %d; want 2", safe)
	}
}

// test part 2
func TestPart2(t *testing.T) {
	safe := part2(parse("input.sample"))
	if 4 != safe {
		t.Errorf("Part 2 safe = %d; want 4", safe)
	}
}
