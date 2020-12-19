package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Go")
	inputFile, _ := ioutil.ReadFile("input")
	lines := strings.Split(string(inputFile), "\n")
	indexes := make(map[int]bool)
	list := []int{}
	offset := 25
	for i := 0; i < len(lines); i++ {
		n, _ := strconv.Atoi(lines[i])
		list = append(list, n)
		if i < offset {
			indexes[list[i]] = true
		}
	}
	for i := offset; i < len(list); i++ {
		fmt.Printf("Number is %d\n", list[i])
		found := false
		for j := i - offset; j < i; j++ {
			diff := list[i] - list[j]
			fmt.Printf("Diff %d\n", diff)
			if indexes[diff] {
				found = true
				break
			}
		}
		if !found {
			fmt.Printf("Bad number %d\n", list[i])
			return
		}
		indexes[list[i-offset]] = false
		indexes[list[i]] = true
	}
}
