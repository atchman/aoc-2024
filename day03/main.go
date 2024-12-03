package main

import (
    "fmt"
    "log"
    "os"
    "regexp"
    "strconv"
)

func main() {
    // result
    fmt.Println("Part 1:", One(parseFile("input")))
}

func parseFile(file string) (mem string) {
    // read file
    fileByte, err := os.ReadFile(file)
    if err != nil {
        log.Fatalf("file error: %v", err)
    }
    mem = string(fileByte)
    return mem
}

// Part One
func One(mem string) (sum int) {
    sum = 0
    r, err := regexp.Compile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
    if err != nil {
        log.Fatalf("regexp: %v", err)
    }
    // get list of all finds
    getMul := (r.FindAllString(mem, -1))

    for _, mul := range getMul {
        re, err := regexp.Compile(`[0-9]{1,3}`)
        if err != nil {
            log.Fatalf("regexp for number: %v", err)
        }
        // regexp to get numbers only
        numbers := (re.FindAllString(mul, -1))
        // a
        a, err := strconv.Atoi(numbers[0])
        if err != nil {
            log.Fatalf("convert str to int, a: %v", err)
        }
        // b
        b, err := strconv.Atoi(numbers[1])
        if err != nil {
            log.Fatalf("convert str to int, b: %v", err)
        }
        sum = sum + (a * b)
    }
    return sum
}
