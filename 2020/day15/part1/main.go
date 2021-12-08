package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type history struct {
	first int
	last  int
}

func NewHistory(first int) history {
	return history{first: first, last: -1}
}

func main() {
	inputFile, _ := ioutil.ReadFile("input")
	lines := strings.Split(strings.TrimSpace(string(inputFile)), ",")
	numbers := []int{}
	tracker := make(map[int]history)
	for i, n := range lines {
		number, _ := strconv.Atoi(n)
		numbers = append(numbers, number)
		tracker[number] = NewHistory(i + 1)
	}
	for i := len(numbers); i <= 30000000; i++ {
		lastNum := numbers[len(numbers)-1]
		if val, ok := tracker[lastNum]; ok {
			if val.last == -1 {
				numbers = append(numbers, 0)
				if _, k := tracker[0]; k {
					if tracker[0].last == -1 {
						tracker[0] = history{first: tracker[0].first, last: len(numbers)}
					} else {
						tracker[0] = history{first: tracker[0].last, last: len(numbers)}
					}
				} else {
					tracker[0] = NewHistory(len(numbers))
				}
			} else {
				diff := tracker[lastNum].last - tracker[lastNum].first
				numbers = append(numbers, diff)
				if _, ok := tracker[diff]; ok {
					if tracker[diff].last == -1 {
						tracker[diff] = history{first: tracker[diff].first, last: len(numbers)}
					} else {
						tracker[diff] = history{first: tracker[diff].last, last: len(numbers)}
					}
				} else {
					tracker[diff] = NewHistory(len(numbers))
				}
			}
		}
	}
	fmt.Println(numbers[30000000-1])

}
