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
	for index, line := range lines {
		if line != "" {
			num, err := strconv.Atoi(line)
			if err != nil {
				fmt.Println(err)
				return
			}
			if index != 0 && prev < num {
				increased++
			}
			prev = num
		}
	}
	fmt.Println(increased)
}
