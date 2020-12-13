package main

import "fmt"
import "io/ioutil"
import "strings"

func main() {
	fmt.Println("vim-go")
	inputFile, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(inputFile), "\n")
	maxId := 0
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
		}
	}
	fmt.Println(maxId)

}

func computeMax(lower int, upper int, line string) int {
	rowNum := 0
	for _, char := range line {
		if char == 'F' || char == 'L' {
			upper = (upper + lower) / 2
			rowNum = upper
		} else {
			lower = (upper+lower)/2 + 1
			rowNum = upper
		}
	}
	return rowNum
}
