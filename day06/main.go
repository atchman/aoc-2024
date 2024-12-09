package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// result
	fmt.Println("Part 1:", One(parseFile("input")))
	fmt.Println("Part 2:", Two(parseFile("input")))
}

func parseFile(file_name string) (width int, height int, obst map[position]int, guard position) {
	// read file
	file, err := os.Open(file_name)
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	// line by line
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	pos_col := 0
	width = 0
	height = 0
	obst = map[position]int{}

	for scanner.Scan() {
		line := scanner.Text()
		height = pos_col
		width = len(line)
		for pos_row, char := range line {
			new_pos := position{pos_row, pos_col}
			if char == '^' {
				guard = position{pos_row, pos_col}
			}
			if char == '#' {
				obst[new_pos] = 1
			}
		}
		pos_col = pos_col + 1
	}
	return width, height, obst, guard
}

// postion
type position struct {
	row int
	col int
}

// rotate (row, col) by 90 degrees to the right
func rotate(old position) position {
	return position{-old.col, old.row}
}

// Part One
func One(width int, height int, obst map[position]int, guard position) int {
	visits := map[position]int{}
	isInside := true
	step := position{0, -1}
	for isInside {
		next_step := position{guard.row + step.row, guard.col + step.col}
		// going outside
		if next_step.row > width || next_step.col > height {
			break
		}

		// obstruction found
		if obst[next_step] > 0 {
			step = rotate(step)
			continue
		}

		// already visited
		if visits[next_step] > 0 {
			guard = next_step
		} else {
			visits[next_step] = 1
			guard = next_step
		}
	}
	return len(visits)
}

// part two
func Two(width int, height int, obst map[position]int, guard position) (positions int) {
	positions = 0
	visits := map[position]int{}
	isInside := true
	step := position{0, -1}
	for isInside {
		next_step := position{guard.row + step.row, guard.col + step.col}
		// going outside
		if next_step.row > width || next_step.col > height {
			break
		}

		// obstruction found
		if obst[next_step] > 0 {
			step = rotate(step)
			continue
		}

		// already visited
		if visits[next_step] > 0 {
			guard = next_step
		} else {
			visits[next_step] = 1
			guard = next_step
		}
	}
	return positions
}
