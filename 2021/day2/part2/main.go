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
	var aim, horizontal, depth int
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
					depth += (aim * steps)
				}
			case "down":
				{
					aim += steps
				}
			case "up":
				{
					aim -= steps
				}
			}
		}
	}
	fmt.Println(horizontal * depth)
}
