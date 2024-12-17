package main

import (
	"fmt"
	"testing"
)

// test part 1
func TestOneSampleOne(t *testing.T) {
	score := One(parseFile("input.sample"))
	expect := 7036
	if score != expect {
		t.Errorf("score = %d; want %v", score, expect)
	}
	fmt.Println("Part 1 (input.sample.one): ", score)
}

func TestOneSampleTwo(t *testing.T) {
	score := One(parseFile("input.sample.two"))
	expect := 11048
	if score != expect {
		t.Errorf("score = %d; want %v", score, expect)
	}
	fmt.Println("Part 1 (input.sample.two): ", score)
}

func TestOne(t *testing.T) {
	score := One(parseFile("input.sample"))
	expect := 1
	if score != expect {
		t.Errorf("score = %d; want %v", score, expect)
	}
	fmt.Println("Part 1 (input): ", score)
}
