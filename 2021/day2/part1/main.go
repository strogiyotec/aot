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
		fmt.Println(err)
		return
	}
	lines := strings.Split(string(content), "\n")
	var horizontal, depth int
	for _, line := range lines {
		if line != "" {
			parts := strings.Split(line, " ")
			action := parts[0]
			steps, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println(err)
				return
			}
			switch action {
			case "forward":
				{
					horizontal += steps
				}
			case "down":
				{
					depth += steps
				}
			case "up":
				{
					depth -= steps
				}
			}
		}
	}
	fmt.Println(horizontal * depth)
}
