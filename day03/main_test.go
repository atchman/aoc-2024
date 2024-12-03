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
