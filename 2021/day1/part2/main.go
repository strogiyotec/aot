package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return
	}
	var prev = 0
	var increased = 0
	lines := strings.Split(string(content), "\n")
	for index := range lines {
		currentSum := 0
		for currentIndex := index; currentIndex < len(lines) && currentIndex < index+3; currentIndex++ {
			if lines[currentIndex] != "" {
				num, err := strconv.Atoi(lines[currentIndex])
				if err != nil {
					fmt.Println(err)
					return
				}
				currentSum += num
			}
		}
		if index != 0 && prev < currentSum {
			increased++
		}
		prev = currentSum
	}
	fmt.Println(increased)
}
