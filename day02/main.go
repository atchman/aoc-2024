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
    fmt.Println("Part 1:", part1(parse("input")))
    fmt.Println("Part 2:", part2(parse("input")))
}

func parse(input string) (reports [][]int) {
    // read file
    file, err := os.Open(input)
    if err != nil {
        log.Fatalf("open file error: %v", err)
    }
    defer file.Close()

    // line by line
    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)
    for scanner.Scan() {
        line := scanner.Text()
        splits := strings.Split(line, " ")

        // convert in []int
        report := []int{}
        for _, split := range splits {
            num, err := strconv.Atoi(split)
            if err != nil {
                log.Fatalf("convert str to int: %v", err)
            }
            report = append(report, num)
        }
        reports = append(reports, report)
    }
    err = scanner.Err()
    if err != nil {
        log.Fatalf("scanner error: %s", err)
    }
    return reports
}

func part1(reports [][]int) (safe int) {
    safe = 0
    for _, report := range reports {
        inc := true
        diff := report[0] - report[1]
        if diff > 0 {
            inc = false
        }

        isSafe := true
        for idx, level := range report {
            if idx+1 < len(report) {
                lvl_diff := level - report[idx+1]
                if inc {
                    if lvl_diff < -3 {
                        isSafe = false
                        break
                    } else if lvl_diff > -1 {
                        isSafe = false
                        break
                    }
                }
                if !inc {
                    if lvl_diff < 1 {
                        isSafe = false
                        break
                    } else if lvl_diff > 3 {
                        isSafe = false
                        break
                    }
                }
            }
        }
        if isSafe {
            safe = safe + 1
        }
    }
    return safe
}

func part2(reports [][]int) (safe int) {
    safe = 0
    for _, r := range reports {
        // copy of slice r
        report := make([]int, len(r))
        copy(report, r)

        isSafe := isSafeReport(report)
        if !isSafe {
            for i:= 0; i<len(report); i++ {

                // copy of slice report
                try := make([]int, len(report))
                copy(try, report)
                try = append(try[:i], try[i+1:]...)
                isSafe = isSafeReport(try)
                if isSafe {
                    break
                }
            }
        }
        if isSafe {
            safe = safe + 1
        }
    }
    return safe
}

func isSafeReport(report []int) (isSafe bool) {
    inc := true
    diff := report[0] - report[1]
    if diff > 0 {
        inc = false
    }

    isSafe = true
    for idx, level := range report {
        if idx+1 < len(report) {
            lvl_diff := level - report[idx+1]
            if inc {
                if lvl_diff < -3 || lvl_diff > -1 {
                    isSafe = false
                    break
                }
            }
            if !inc {
                if lvl_diff < 1 || lvl_diff > 3 {
                    isSafe = false
                    break
                }
            }
        }

    }
    return isSafe
}
