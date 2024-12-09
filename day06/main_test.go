package main

import (
	"fmt"
	"testing"
)

// test part one
func TestOneSample(t *testing.T) {
	visits := One(parseFile("input.sample"))
	if 41 != visits {
		t.Errorf("visits = %d; want 41", visits)
	}
	fmt.Println("Part 1 (input.sample): ", visits)
}

func TestOne(t *testing.T) {
	visits := One(parseFile("input"))
	if 4977 != visits {
		t.Errorf("visits = %d; want 4977", visits)
	}
	fmt.Println("Part 1 (input): ", visits)
}

// test part two
func TestTwoSample(t *testing.T) {
	visits := Two(parseFile("input.sample"))
	if 6 != visits {
		t.Errorf("visits = %d; want 6", visits)
	}
	fmt.Println("Part 2 (input.sample): ", visits)
}

func TestTwo(t *testing.T) {
	visits := Two(parseFile("input"))
	if 1 != visits {
		t.Errorf("visits = %d; want 1", visits)
	}
	fmt.Println("Part 2 (input): ", visits)
}
