package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")
	var left = map[byte]bool{
		'(': true,
		'{': true,
		'[': true,
		'<': true,
	}
	score := []int{}
	for _, line := range lines {
		sum := 0
		opened := []int{}
		illegal := false
		for i := 0; i < len(line); i++ {
			_, exist := left[line[i]]
			if exist {
				opened = append(opened, i)
			} else {
				switch line[i] {
				case '}':
					{
						if line[opened[len(opened)-1]] == '{' {
							//remove last
							opened = opened[:len(opened)-1]
						} else {
							illegal = true
						}
					}
				case '>':
					{
						if line[opened[len(opened)-1]] == '<' {
							//remove last
							opened = opened[:len(opened)-1]
						} else {
							illegal = true
						}
					}
				case ']':
					{
						if line[opened[len(opened)-1]] == '[' {
							//remove last
							opened = opened[:len(opened)-1]
						} else {
							illegal = true
						}
					}
				case ')':
					{
						if line[opened[len(opened)-1]] == '(' {
							//remove last
							opened = opened[:len(opened)-1]
						} else {
							illegal = true
						}
					}
				}
				if illegal {
					break
				}
			}
		}
		if len(opened) != 0 && !illegal {
			for i := len(opened) - 1; i >= 0; i-- {
				switch line[opened[i]] {
				case '(':
					{
						sum *= 5
						sum += 1
					}

				case '[':
					{
						sum *= 5
						sum += 2
					}
				case '{':
					{
						sum *= 5
						sum += 3
					}
				case '<':
					{
						sum *= 5
						sum += 4
					}
				}
			}
			score = append(score, sum)
		}
	}
	sort.Ints(score)
	fmt.Println(score[len(score)/2])
}
