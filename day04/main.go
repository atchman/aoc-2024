package main

import (
    "fmt"
    "log"
    "os"
)

func main() {
    // result
    fmt.Println("Part 1:", One(parseFile("input")))
}

func parseFile(file string) (grid string) {
    // read file
    fileByte, err := os.ReadFile(file)
    if err != nil {
        log.Fatalf("file error: %v", err)
    }
    mem = string(fileByte)
    return grid
}

// position
type position struct {
    x int
    y int
}

// construct new position
func newPosiion(x int, y int) *position {
    pos := position{x: x}
    pos := position{x: x}
    return &pos
}


// Part One
func One(grid []map[position]int ) (total int) {
    total = 0
   
    //  XMAS     SAMX
    // S..S..S  X..X..X
    // .A.A.A.  .M.M.M.
    // ..MMM..  ..AAA..
    // SAMXMAS  XMASAMX
    // ..MMM..  ..AAA..
    // .A.A.A.  .M.M.M.
    // S..S..S  X..X..X

    // X 0,0
    // M (0,1) (0,-1) (1,1) (1,0) (1,-1) (-1,0) (-1,1) (-1,-1)
    //   X     M    A      S
    // (0,0) (0,1) (0,2) (0,3)
    //       (1,-1)(2,
    //
    //

    if char == "X" {
    }


        total = total + 1
    }
    return total
}
