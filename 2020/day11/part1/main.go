package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type tuple struct {
	first  int
	second int
}

func main() {
	inputFile, _ := ioutil.ReadFile("input")
	lines := strings.Split(string(inputFile), "\n")
	matrix := [][]rune{}
	for _, line := range lines {
		matrixLine := []rune{}
		for _, r := range line {
			matrixLine = append(matrixLine, r)
		}
		matrix = append(matrix, matrixLine)
	}
	occupied := make(map[tuple]bool)
	empty := make(map[tuple]bool)
	for i, row := range matrix {
		for j, val := range row {
			if val == '#' {
				occupied[tuple{first: i, second: j}] = true
			} else if matrix[i][j] == 'L' {
				empty[tuple{first: i, second: j}] = true
			}
		}
	}
	positions := []tuple{{first: 0, second: -1}, {first: 0, second: 1}, {first: 1, second: 0}, {first: -1, second: 0}, {first: -1, second: -1}, {first: 1, second: 1}, {first: 1, second: -1}, {first: -1, second: 1}}
	for true {
		newEmpty := make(map[tuple]bool)
		newOccupied := make(map[tuple]bool)
		for k := range occupied {
			adjacentCnt := 0
			for _, position := range positions {
				pos := tuple{first: k.first + position.first, second: k.second + position.second}
				if occupied[pos] {
					adjacentCnt++
				}
				if adjacentCnt >= 4 {
					newEmpty[k] = true
					break
				}
			}
			if adjacentCnt < 4 {
				newOccupied[k] = true
			}
		}
		for k := range empty {
			hasOccupiedAdj := false
			for _, position := range positions {
				if occupied[tuple{first: k.first + position.first, second: k.second + position.second}] {
					hasOccupiedAdj = true
					break
				}
			}
			if !hasOccupiedAdj {
				newOccupied[k] = true
			} else {
				newEmpty[k] = true
			}
		}
		if len(occupied) == len(newOccupied) {
			break
		}
		occupied = newOccupied
		empty = newEmpty
	}
	fmt.Println(len(occupied))
}
