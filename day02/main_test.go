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
