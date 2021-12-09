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
	var builder strings.Builder
	for i := 0; i < len(lines[0]); i++ {
		var ones, zeroes int
		for index := 0; index < len(lines); index++ {
			if lines[index] != "" {
				if lines[index][i] == '0' {
					zeroes++
				} else {
					ones++
				}
			}
		}
		if ones > zeroes {
			builder.WriteRune('1')
		} else {
			builder.WriteRune('0')
		}
	}
	gamma := builder.String()
	var epsilonBuilder strings.Builder
	for i := range gamma {
		if gamma[i] == '0' {
			epsilonBuilder.WriteRune('1')
		} else {
			epsilonBuilder.WriteRune('0')
		}
	}
	if gammaInt, err := strconv.ParseInt(gamma, 2, 32); err == nil {
		if epsilonInt, err := strconv.ParseInt(epsilonBuilder.String(), 2, 32); err == nil {
			fmt.Println(gammaInt * epsilonInt)
		}
	}
}
