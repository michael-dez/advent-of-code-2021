/*
goal:
find the most common bit (mcb) and least common bit (lcb) for each position of input. e.g.:
12345
-----
00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010
^most common bit is 1
provide answer in decimal.
functions
getBits(numState) return completed numState pointer

structs ?
*/

package main

import (
"os"
"bufio"
"fmt"
"log"
)

type numState struct {
    sNumbers []string
    size int
    oneCount []int
    zeroCount []int
    gRate []string
    eRate []string
}

func New(n []string) *numState {
    e := new(numState)
    e.sNumbers = n
    e.size = len(n)
    e.oneCount = make([]int, len(n))
    e.zeroCount = make([]int, len(n))
    e.gRate = make([]string, len(n))
    e.eRate = make([]string, len(n))

    for i := 0; i < len(n); i++ {
        e.zeroCount[i] = 0
        e.oneCount[i] = 0
    }
    return e
}

func (e *numState) findPower() {
    
}

func (e *numState) findRate() {
// for every index of oneCount, compare count to median of size to calculate
// gamma/epsilon rate
    for i := 0; i < 6; i++ { // I don't like this
        if med := e.size / 2; e.oneCount[i] > med {
            e.gRate[i] = "1"
            e.eRate[i] = "0"
            fmt.Println(i, e.gRate[i], e.eRate[i])
        } else {
            e.gRate[i] = "0"
            e.eRate[i] = "1"
            fmt.Println(i, e.gRate[i], e.eRate[i])
        }
    }
    e.TODO
}

func (e *numState) rowCount(r string) {
// update zeroCount and oneCount for each char in row r
    for i := 0; i < len(r); i++ {
        switch nibbleVal := r[i]; nibbleVal{
            case 48:
                    e.zeroCount[i]++
            case 49:
                    e.oneCount[i]++
        }
    }  

}
func (e *numState) getCounts() {
//call rowCount to update zeroCount and oneCount for each row
    for i := 0; i < e.size; i++ {
       e.rowCount(e.sNumbers[i]) 
    }
    e.findRate()
}
func main() {
    file, err := os.Open("test_input")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    snums := []string{}
    for scanner.Scan() {
        n := scanner.Text()
        snums = append(snums, n)
    }
    tracker := New(snums)
    tracker.getCounts()
    fmt.Println(tracker.eRate, tracker.gRate)
    fmt.Println(tracker.size)
}
