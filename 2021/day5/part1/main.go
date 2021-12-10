package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type parts struct {
	from int
	to   int
}

func NewPart(line string) parts {
	points := strings.Split(strings.TrimSpace(line), ",")
	from, _ := strconv.Atoi(points[0])
	to, _ := strconv.Atoi(points[1])
	return parts{from: from, to: to}
}

func (part parts) similar(another parts) bool {
	return part.from == another.from || part.to == another.to
}

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")
	cnt := make(map[string]int)
	for _, v := range lines {
		parts := strings.Split(v, "->")
		if len(parts) == 2 {
			firstPart := NewPart(parts[0])
			secondPart := NewPart(parts[1])
			if firstPart.similar(secondPart) {
				from := min(firstPart, secondPart)
				to := max(firstPart, secondPart)
				sameX := sameX(firstPart, secondPart)
				for i := from; i <= to; i++ {
					var element string
					if sameX {
						element = fmt.Sprintf("%d-%d", firstPart.from, i)
					} else {
						element = fmt.Sprintf("%d-%d", i, firstPart.to)
					}
					value, exist := cnt[element]
					if !exist {
						cnt[element] = 1
					} else {
						cnt[element] = value + 1
					}
				}
			}
		}
	}
	totalPoints := 0
	for _, v := range cnt {
		if v >= 2 {
			totalPoints++
		}
	}
	fmt.Println(totalPoints)
}

func sameX(first, second parts) bool {
	if first.from == second.from {
		return true
	}
	return false
}
func max(first, second parts) int {
	if first.from == second.from {
		if first.to > second.to {
			return first.to
		} else {
			return second.to
		}
	} else {
		if first.from > second.from {
			return first.from
		} else {
			return second.from
		}
	}
}
func min(first, second parts) int {
	if first.from == second.from {
		if first.to < second.to {
			return first.to
		} else {
			return second.to
		}
	} else {
		if first.from < second.from {
			return first.from
		} else {
			return second.from
		}
	}
}
