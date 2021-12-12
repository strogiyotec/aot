package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)
	sc.Scan()
	positions := parse(sc.Text())
	sort.Ints(positions)

	alignPoint := positions[len(positions)/2]

	fuelCost := 0
	for _, position := range positions {
		fuelCost += int(math.Abs(float64(position - alignPoint)))
	}
	fmt.Println(fuelCost)
}

func median(numbers []int) int {
	middle := len(numbers) / 2
	return numbers[middle]
}

func parse(line string) []int {
	numbers := []int{}
	parts := strings.Split(line, ",")
	for _, v := range parts {
		num, _ := strconv.Atoi(v)
		numbers = append(numbers, num)
	}
	return numbers
}
