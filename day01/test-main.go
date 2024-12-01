package main

import (
	"github.com/gkampitakis/go-snaps/snaps"
	"testing"
)

func Test_Part1(t *testing.T) {
	snaps.MatchSnapshot(t,
		[]int{11, part1()})
}
