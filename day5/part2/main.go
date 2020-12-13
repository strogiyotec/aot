package main

import "fmt"
import "io/ioutil"
import "strings"

func main() {
	fmt.Println("vim-go")
	inputFile, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(inputFile), "\n")
	ids := make(map[int]bool)
	maxId := 0
	minId := 100000
	for _, line := range lines {
		if len(line) != 0 {
			row := line[0:7]
			rowNum := computeMax(0, 127, row)
			column := line[7:10]
			columnNum := computeMax(0, 7, column)
			id := rowNum*8 + columnNum
			if maxId < id {
				maxId = id
			}
			if minId > id {
				minId = id
			}
			fmt.Println(id)
			ids[id] = true
		}
	}
	minId++
	maxId--
	missing := 0
	fmt.Println("here")
	for i := minId; i < maxId; i++ {
		if ids[i-1] && ids[i+1] && !ids[i] {
			missing = i
		}
	}
	fmt.Println(missing)
}

func computeMax(lower int, upper int, line string) int {
	rowNum := 0
	for _, char := range line {
		if char == 'F' || char == 'L' {
			upper = (upper + lower) / 2
			rowNum = upper
		} else {
			lower = (upper+lower)/2 + 1
			rowNum = lower
		}
	}
	return rowNum
}
