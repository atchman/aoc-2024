package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// result
	fmt.Println("Part 1:", One(parseFile("input")))
	fmt.Println("Part 2:", Two(parseFile("input")))
}

func parseFile(file_name string) (equations [][]int) {
	// read file
	file, err := os.Open(file_name)
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	// line by line
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()

		// splits by two seperators
		// anonymous function
		string_list := strings.FieldsFunc(line, func(r rune) bool {
			return r == ':' || r == ' '
		})

		var numbers []int
		for _, str := range string_list {
			num, err := strconv.Atoi(str)
			if err != nil {
				log.Fatalf("convert str to int: %v", err)
			}
			numbers = append(numbers, num)
		}
		equations = append(equations, numbers)
	}
	return equations
}

// Part One
func One(equations [][]int) (total int) {
	total = 0

	for _, equation := range equations {
		isEqual := false
		result := equation[0]
		rest := make([]int, len(equation)-1)
		copy(rest, equation[1:])

		result_list := make([]int, 1)
		result_list[0] = rest[0]
		rest = rest[1:]

		for _, num := range rest {
			var new_list []int
			for _, x := range result_list {
				new_list = append(new_list, (x + num))
				new_list = append(new_list, (x * num))
			}
			result_list = new_list
		}

		for _, y := range result_list {
			if y == result {
				isEqual = true
				break
			}
		}

		if isEqual {
			total = total + result
		}
	}
	return total
}

// Part Two
func Two(equations [][]int) (total int) {
	total = 0

	for _, equation := range equations {
		isEqual := false
		result := equation[0]
		rest := make([]int, len(equation)-1)
		copy(rest, equation[1:])

		result_list := make([]int, 1)
		result_list[0] = rest[0]
		rest = rest[1:]

		for _, num := range rest {
			var new_list []int
			for _, x := range result_list {
				new_list = append(new_list, (x + num))
				new_list = append(new_list, (x * num))

				// concatenation
				new_str := strconv.Itoa(x) + strconv.Itoa(num)
				new_int, err := strconv.Atoi(new_str)
				if err != nil {
					log.Fatalf("convert str to int, new_int: %v ", err)
				}
				new_list = append(new_list, new_int)
			}
			result_list = new_list
		}

		// looking if result is found
		for _, y := range result_list {
			if y == result {
				isEqual = true
				break
			}
		}

		if isEqual {
			total = total + result
		}
	}
	return total
}
