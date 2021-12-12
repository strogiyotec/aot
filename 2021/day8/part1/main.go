package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	//Read the file
	content, _ := ioutil.ReadFile("input.txt")
	values := strings.Split(string(content), "\n")
	count := 0
	for _, v := range values {
		parts := strings.Split(v, "|")
		if len(parts) == 2 {
			fields := strings.Fields(parts[1])
			for _, field := range fields {
				length := len(field)
				if length == 2 || length == 3 || length == 4 || length == 7 {
					count++
				}
			}
		}
	}
	fmt.Println(count)
}
