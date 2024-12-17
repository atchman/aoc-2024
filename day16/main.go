package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func main() {
	// result
	fmt.Println("Part 1:", One(parseFile("input")))
	//fmt.Println("Part 2:", Two(parseFile("input")))
}

func parseFile(file_name string) (start position, end position, wall map[position]uint, node map[position]uint) {
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
	wall = map[position]uint{}
	node = map[position]uint{}
	for scanner.Scan() {
		line := scanner.Text()

		for pos_row, char := range line {
			pos := position{pos_row, pos_col}
			// switch
			switch char {
			case 'S':
				start = pos
				node[pos] = 0
			case 'E':
				end = pos
				node[pos] = 0
			case '#':
				wall[pos] = 1
			case '.':
				node[pos] = 1
			}
		}
		pos_col = pos_col + 1
	}
	return start, end, wall, node
}

type position struct {
	row int
	col int
}

// Part One
func One(start position, end position, wall map[position]uint, node map[position]uint) (score uint) {
	score = 0

	// init dijkstra
	cost := make(map[position]uint, len(node))
	visited := make(map[position]uint, len(node))
	pre := make(map[position][]position, len(node))
	for pos := range node {
		cost[pos] = math.MaxUint32
	}
	cost[start] = 0

	return score
}

// Part Two