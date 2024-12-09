package main

import (
	"fmt"
	"testing"
)

// test part 1
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
