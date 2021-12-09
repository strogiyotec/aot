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
	oxygen := oxygen(lines)
	co2 := co2(lines)
	fmt.Println(oxygen * co2)
}

func co2(lines []string) int64 {
	slice := lines[:]
	charIndex := 0
	for len(slice) != 1 && charIndex < len(slice[0]) {
		ones := []string{}
		zeros := []string{}
		for i := 0; i < len(slice); i++ {
			if slice[i] != "" {
				if slice[i][charIndex] == '0' {
					zeros = append(zeros, slice[i])
				} else {
					ones = append(ones, slice[i])
				}
			}
		}
		charIndex++
		if len(zeros) <= len(ones) {
			slice = zeros
		} else {
			slice = ones
		}
	}
	if co2, err := strconv.ParseInt(slice[0], 2, 32); err == nil {
		return co2
	}
	return -1
}
func oxygen(lines []string) int64 {
	slice := lines[:]
	charIndex := 0
	for len(slice) != 1 && charIndex < len(slice[0]) {
		ones := []string{}
		zeros := []string{}
		for i := 0; i < len(slice); i++ {
			if slice[i] != "" {
				if slice[i][charIndex] == '0' {
					zeros = append(zeros, slice[i])
				} else {
					ones = append(ones, slice[i])
				}
			}
		}
		charIndex++
		if len(ones) >= len(zeros) {
			slice = ones
		} else {
			slice = zeros
		}
	}
	if oxygen, err := strconv.ParseInt(slice[0], 2, 32); err == nil {
		return oxygen
	}
	return -1
}
