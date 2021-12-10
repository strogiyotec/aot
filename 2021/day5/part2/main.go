package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type parts struct {
	x int
	y int
}

func NewPart(line string) parts {
	points := strings.Split(strings.TrimSpace(line), ",")
	from, _ := strconv.Atoi(points[0])
	to, _ := strconv.Atoi(points[1])
	return parts{x: from, y: to}
}

func (part parts) similar(another parts) bool {
	return part.x == another.x || part.y == another.y
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
			//horiz or vertical
			if firstPart.similar(secondPart) {
				x := min(firstPart, secondPart)
				y := max(firstPart, secondPart)
				sameX := sameX(firstPart, secondPart)
				for i := x; i <= y; i++ {
					var element string
					if sameX {
						element = fmt.Sprintf("%d-%d", firstPart.x, i)
					} else {
						element = fmt.Sprintf("%d-%d", i, firstPart.y)
					}
					value, exist := cnt[element]
					if !exist {
						cnt[element] = 1
					} else {
						cnt[element] = value + 1
					}
				}
			} else {
				//diagonal
				fromY, toY, fromX, toRight := diagDirection(firstPart, secondPart)
				for fromY <= toY {
					element := fmt.Sprintf("%d-%d", fromX, fromY)
					value, exist := cnt[element]
					if !exist {
						cnt[element] = 1
					} else {
						cnt[element] = value + 1
					}
					fromY++
					if toRight {
						fromX++
					} else {
						fromX--
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
	if first.x == second.x {
		return true
	}
	return false
}

func diagDirection(first, second parts) (int, int, int, bool) {
	if first.y < second.y {
		diagRight := first.x < second.x
		return first.y, second.y, first.x, diagRight
	} else {
		diagRight := second.x < first.x
		return second.y, first.y, second.x, diagRight
	}
}

func max(first, second parts) int {
	if first.x == second.x {
		if first.y > second.y {
			return first.y
		} else {
			return second.y
		}
	} else {
		if first.x > second.x {
			return first.x
		} else {
			return second.x
		}
	}
}
func min(first, second parts) int {
	if first.x == second.x {
		if first.y < second.y {
			return first.y
		} else {
			return second.y
		}
	} else {
		if first.x < second.x {
			return first.x
		} else {
			return second.x
		}
	}
}
