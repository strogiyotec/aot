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
	earliest, _ := strconv.Atoi(lines[0])
	busses := []int{}

	ids := strings.Split(string(lines[1]), ",")
	for _, id := range ids {
		numId, e := strconv.Atoi(id)
		if e == nil {
			busses = append(busses, numId)
		}
	}
	fmt.Println(busses, earliest)
	smallestNearest := 2147483647
	busId := 0
	for _, bus := range busses {
		wait := ((earliest / bus) * bus) + bus
		if smallestNearest > wait {
			smallestNearest = wait
			busId = bus
		}
	}
	needToWait := (smallestNearest - earliest) * busId
	fmt.Println(needToWait)
}
