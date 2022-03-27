package main

import (
"os"
"bufio"
"log"
"strings"
"strconv"
"fmt"
)

/*functions?
get input(slice)

struct?
submarine
    props
    - depth
    - horizontal position (hpos)
    functions
    - depthXHpos(self) int
    - move
*/

type submarine struct {
    Depth int
    HPos int
    Aim int
}

func New() *submarine {
    e := new(submarine)
    e.Depth = 0
    e.HPos = 0
    e.Aim = 0
    return e
}

func (s submarine) DepthXHPos() int {
    return s.Depth * s.HPos
}

func (s *submarine) Move(m string) {
    command := strings.Fields(m)
    count, err := strconv.Atoi(command[1])
    if err != nil {
        log.Fatal(err)
    }

    switch direction := command[0]; direction {
    case "forward":
        s.HPos += count
        s.Depth += s.Aim * count
    case "up":
        s.Aim -= count
    case "down":
        s.Aim += count
    }

    fmt.Printf("mv=%[1]s %[2]d, hpos=%[3]d, depth=%[4]d, aim=%[5]d\n", command[0], count, s.HPos, s.Depth, s.Aim) //debug
}

func main() {
    file, err := os.Open("input") // open puzzle input file
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    moves := []string{}
    for scanner.Scan() {
        s := scanner.Text()
        moves = append(moves, s)
    }

    sub := New()

    for _, v := range moves {
        sub.Move(v)
    }

    fmt.Println(sub.DepthXHPos()) //output puzzle 1 solution
}
