package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	days := days(string(content))
	fishes := make([]uint64, 9)
	for _, v := range days {
		fishes[v]++
	}
	for i := 0; i < 256; i++ {
		nextFishes := make([]uint64, 9)
		nextFishes[6] = fishes[0]
		nextFishes[8] = fishes[0]
		for j := 1; j < 9; j++ {
			nextFishes[j-1] += fishes[j]
		}
		fishes = nextFishes
	}
	var sum uint64
	for _, v := range fishes {
		sum += v
	}
	fmt.Println(sum)

}

func days(content string) []int {
	days := []int{}
	parts := strings.Split(content, ",")
	for _, v := range parts {
		day, _ := strconv.Atoi(strings.TrimSpace(v))
		days = append(days, day)
	}
	return days
}
