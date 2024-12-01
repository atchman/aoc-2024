package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var left []int
	var right []int

	// result
	left, right = parse("input")
	fmt.Println("Part 1:", part1(left, right))
	fmt.Println("Part 2:", part2(left, right))
}

func parse(input string) (left []int, right []int) {
	// read file
	file, err := os.Open(input)
	if err != nil {
		log.Fatalf("open file error: %v", err)
	}

	// line by line
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "   ")

		num_left, err := strconv.Atoi(split[0])
		if err != nil {
			log.Fatalf("Convert left number as str to int: %v", err)
		}

		num_right, err := strconv.Atoi(split[1])
		if err != nil {
			log.Fatalf("convert right number as str to int: %v", err)
		}

		left = append(left, num_left)
		right = append(right, num_right)
	}
	err = scanner.Err()
	if err != nil {
		log.Fatalf("scanner error: %s", err)
	}
	file.Close()

	return left, right
}

func part1(left []int, right []int) (total int) {
	sort.Ints(left)
	sort.Ints(right)

	total = 0

	for idx, val := range left {
		if val >= right[idx] {
			total = total + (val - right[idx])
		} else {
			total = total + (right[idx] - val)
		}
	}
	return total
}

func part2(left []int, right []int) (score int) {
	score = 0
	for _, val := range left {
		count := 0
		for _, comp := range right {
			if val == comp {
				count = count + 1
			}
		}
		score = score + (val * count)
		count = 0
	}
	return score
}
