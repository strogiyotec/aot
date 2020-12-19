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
	var badNumberIndex int
	for i := offset; i < len(list); i++ {
		found := false
		for j := i - offset; j < i; j++ {
			diff := list[i] - list[j]
			if indexes[diff] {
				found = true
				break
			}
		}
		if !found {
			badNumberIndex = i
			fmt.Printf("Bad number %d\n", list[i])
			break
		}
		indexes[list[i-offset]] = false
		indexes[list[i]] = true
	}
	fmt.Printf("Index is %d\n", badNumberIndex)
	sum := 0
	smallest := 2147483647
	biggest := 0
	left := 0
	for i := 0; i < badNumberIndex; i++ {
		sum += list[i]
		if smallest > list[i] {
			smallest = list[i]
		}
		if biggest < list[i] {
			biggest = list[i]
		}
		for sum > list[badNumberIndex] {
			sum -= list[left]
			left++
			smallest = 2147483647
			biggest = 0
			// We can user priority queue to store min and max
			for j := left; j < i; j++ {
				if smallest > list[j] {
					smallest = list[j]
				}
				if biggest < list[j] {
					biggest = list[j]
				}
			}
		}
		if sum == list[badNumberIndex] {
			diff := smallest + biggest
			fmt.Printf("Smallest %d Biggest %d Sum %d\n", smallest, biggest, diff)
		}
	}
}
