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
}

func New() *submarine {
    e := new(submarine)
    e.Depth = 0
    e.HPos = 0
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
    case "up":
        s.Depth -= count
    case "down":
        s.Depth += count
    }

    fmt.Printf("move:%[1]s %[2] = %[3]d\n", command[0], count, s.HPos) //debug
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
