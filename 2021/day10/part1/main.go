package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")
	sum := 0
	var left = map[byte]bool{
		'(': true,
		'{': true,
		'[': true,
		'<': true,
	}
	opened := []int{}
	for _, line := range lines {
		for i := 0; i < len(line); i++ {
			_, exist := left[line[i]]
			if exist {
				opened = append(opened, i)
			} else {
				illegal := false
				switch line[i] {
				case '}':
					{
						if line[opened[len(opened)-1]] == '{' {
							//remove last
							opened = opened[:len(opened)-1]
						} else {
							sum += 1197
							illegal = true
						}
					}
				case '>':
					{
						if line[opened[len(opened)-1]] == '<' {
							//remove last
							opened = opened[:len(opened)-1]
						} else {
							sum += 25137
							illegal = true
						}
					}
				case ']':
					{
						if line[opened[len(opened)-1]] == '[' {
							//remove last
							opened = opened[:len(opened)-1]
						} else {
							sum += 57
							illegal = true
						}
					}
				case ')':
					{
						if line[opened[len(opened)-1]] == '(' {
							//remove last
							opened = opened[:len(opened)-1]
						} else {
							sum += 3
							illegal = true
						}
					}
				}
				if illegal {
					break
				}
			}
		}
	}
	fmt.Println(sum)
}
