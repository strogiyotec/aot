package main

import "fmt"
import "io/ioutil"
import "strings"

func main() {
	fmt.Println("vim-go")
	inputFile, _ := ioutil.ReadFile("input")
	lines := strings.Split(string(inputFile), "\n\n")
	cnt := 0
	for _, line := range lines {
		chars := make(map[rune]int)
		line = strings.TrimSuffix(line, "\n")
		answers := strings.Split(string(line), "\n")
		answersAmount := len(answers)
		for _, answer := range answers {
			for _, char := range answer {
				chars[char]++
				if chars[char] == answersAmount {
					cnt++
				}
			}
		}
	}
	fmt.Println(cnt)
}
