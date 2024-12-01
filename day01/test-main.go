package main

import (
	"github.com/gkampitakis/go-snaps/snaps"
	"testing"
)

// test part 1
func Test_Part1(t *testing.T) {
	snaps.MatchSnapshot(t, []int{11, part1()})
}

// test part 2
func Test_Part2(t *testing.T) {
	snaps.MatchSnapshot(t, []int{11, part2()})
}
