package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// TODO check how jump works, the way I implemented is not correct
func main() {
	inputFile, _ := ioutil.ReadFile("small_input")
	lines := strings.Split(string(inputFile), "\n")
	counterTracker := make(map[string]int)
	lineCnt := 0
	sum := 0
	for counterTracker[lines[lineCnt]] != 2 {
		fmt.Printf("%d %d\n", sum, lineCnt)
		counterTracker[lines[lineCnt]]++
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
