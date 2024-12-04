package main

import (
	"testing"
)

// test part 1
func TestOne(t *testing.T) {
	sum := One(parseFile("input.sample"))
	if 161 != sum {
		t.Errorf("Sum = %d; want 161", sum)
	}
}

// test part 2
func TestTwo(t *testing.T) {
	sum := Two(parseFile("input.sample"))
	if 48 != sum {
		t.Errorf("Sum = %d; want 48", sum)
	}
}
