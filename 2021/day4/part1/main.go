package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Board struct {
	board         [][]int
	nonVisitedSum int
	visitedRows   map[int]bool
	visitedCol    map[int]bool
}

func NewBoard(lines []string) *Board {
	board := make([][]int, 5)
	sum := 0
	for i := 0; i < 5; i++ {
		board[i] = make([]int, 5)
	}
	for i, v := range lines {
		parts := strings.Fields(v)
		for row, part := range parts {
			intVal, _ := strconv.Atoi(part)
			sum += intVal
			board[i][row] = intVal
		}
	}
	return &Board{
		board:         board,
		nonVisitedSum: sum,
		visitedRows:   make(map[int]bool),
		visitedCol:    make(map[int]bool),
	}
}

func (board *Board) mark(val int) {
	for row := 0; row < len(board.board); row++ {
		for col := 0; col < len(board.board[0]); col++ {
			if board.board[row][col] == val {
				board.nonVisitedSum -= val
				board.board[row][col] = -1
			}
		}
	}
}

func (board Board) checkWin() bool {
	for col := 0; col < len(board.board[0]); col++ {
		if board.checkCol(col) {
			_, exist := board.visitedCol[col]
			if !exist {
				for i := 0; i < len(board.board[0]); i++ {
					board.visitedCol[i] = true
				}
				return true
			}
		}
	}
	for row := 0; row < len(board.board); row++ {
		if board.checkRow(row) {
			_, exist := board.visitedRows[row]
			if !exist {
				for i := 0; i < len(board.board); i++ {
					board.visitedRows[i] = true
				}
				return true
			}
		}
	}
	return false
}

func (board Board) checkCol(col int) bool {
	for i := 0; i < len(board.board); i++ {
		if board.board[i][col] != -1 {
			return false
		}
	}
	return true
}
func (board Board) checkRow(row int) bool {
	for _, v := range board.board[row] {
		if v != -1 {
			return false
		}
	}
	return true
}

func main() {
	boards := []Board{}
	content, _ := ioutil.ReadFile("small_input.txt")
	lines := strings.Split(string(content), "\n")
	numbers := numbers(lines[0])
	for i := 2; i+5 < len(lines); i += 6 {
		board := NewBoard(lines[i : i+5])
		boards = append(boards, *board)
	}
	lastNumber := -1
	lastIndex := -1
	for _, number := range numbers {
		for index, board := range boards {
			board.mark(number)
			boards[index].nonVisitedSum = board.nonVisitedSum
			if board.checkWin() {
				print(board.board)
				lastNumber = number
				lastIndex = index
				fmt.Println(lastNumber, lastIndex)
			}
		}
	}
	fmt.Println(lastNumber * boards[lastIndex].nonVisitedSum)
}

func print(board [][]int) {
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++ {
			fmt.Print(board[row][col], ",")
		}
		fmt.Println()
	}
}

func numbers(line string) []int {
	numbers := []int{}
	parts := strings.Split(line, ",")
	for _, v := range parts {
		intVal, _ := strconv.Atoi(v)
		numbers = append(numbers, intVal)
	}
	return numbers
}
