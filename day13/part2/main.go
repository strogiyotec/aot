package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	inputFile, _ := ioutil.ReadFile("input")
	lines := strings.Split(string(inputFile), "\n")
	busses := []int{}

	ids := strings.Split(string(lines[1]), ",")
	for _, id := range ids {
		numId, e := strconv.Atoi(id)
		if e == nil {
			busses = append(busses, numId)
		} else {
			busses = append(busses, 1)
		}
	}
	times := 0
	stepSize := busses[0]
	for i := 1; i < len(busses); i++ {
		for (times+i)%busses[i] != 0 {
			times += stepSize
		}
		stepSize *= busses[i]
	}
	fmt.Println(times)
}
