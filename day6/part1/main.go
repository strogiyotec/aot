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
		chars := make(map[rune]bool)
		for _, char := range line {
			if char != '\n' {
				chars[char] = true
			}
		}
		cnt += len(chars)
	}
	fmt.Println(cnt)

}
