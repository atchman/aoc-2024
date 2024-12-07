package main

import (
	"fmt"
	"testing"
)

// test part 1
func TestOneSample(t *testing.T) {
	total := One(parseFile("input.sample"))
	if 3749 != total {
		t.Errorf("total = %d; want 3749", total)
	}
	fmt.Println("Part 1 (input.sample): ", total)
}

func TestOne(t *testing.T) {
	total := One(parseFile("input"))
	if 6392012777720 != total {
		t.Errorf("total = %d; want 6392012777720", total)
	}
	fmt.Println("Part 1 (input): ", total)
}

// test part 2
func TestTwoSample(t *testing.T) {
	total := Two(parseFile("input.sample"))
	if 11387 != total {
		t.Errorf("total = %d; want 11387", total)
	}
	fmt.Println("Part 2 (input.sample): ", total)
}

func TestTwo(t *testing.T) {
	total := Two(parseFile("input"))
	if 61561126043536 != total {
		t.Errorf("total = %d; want 61561126043536", total)
	}
	fmt.Println("Part 2 (input): ", total)
}
