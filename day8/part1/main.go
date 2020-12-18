package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// TODO check how jump works, the way I implemented is not correct
func main() {
	inputFile, _ := ioutil.ReadFile("input")
	lines := strings.Split(string(inputFile), "\n")
	counterTracker := make(map[int]bool)
	lineCnt := 0
	sum := 0
	for !counterTracker[lineCnt] {
		fmt.Printf("%d %d\n", sum, lineCnt)
		counterTracker[lineCnt] = true
		instruction := lines[lineCnt][0:3]
		sign := lines[lineCnt][4]
		number, _ := strconv.Atoi(lines[lineCnt][5:])
		if instruction == "nop" {
			lineCnt++
		} else if instruction == "acc" {
			if sign == '-' {
				sum -= number
			} else {
				sum += number
			}
			lineCnt++
		} else {
			if sign == '-' {
				lineCnt -= number
			} else {
				lineCnt += number
			}
		}
	}
	fmt.Println(sum)

}
