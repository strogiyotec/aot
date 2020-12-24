package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type dockData struct {
	number   int
	position int
}

func main() {
	inputFile, _ := ioutil.ReadFile("input")
	lines := strings.Split(string(inputFile), "\n")
	index := 0
	table := make(map[int]int)
	r := regexp.MustCompile("mask = ([^ ]+)")
	numR := regexp.MustCompile("mem\\[([^]]+)\\] = ([^ ]+)")
	for index < len(lines) {
		if strings.Contains(lines[index], "mask") {
			res := r.FindAllStringSubmatch(lines[index], -1)
			mask := res[0][1]
			index++
			for !strings.Contains(lines[index], "mask") && len(lines[index]) != 0 {
				fmt.Printf("Line is %s index is %d\n", lines[index], index)
				numPart := numR.FindAllStringSubmatch(lines[index], -1)
				tableIndex, _ := strconv.Atoi(numPart[0][1])
				numP, _ := strconv.Atoi(numPart[0][2])
				numStr := strconv.FormatInt(int64(numP), 2)
				fmt.Printf("Converted %s\n", numStr)
				if len(numStr) < 36 {
					var builder strings.Builder
					trailZeros := 36 - len(numStr)
					for trailZeros > 0 {
						builder.WriteRune('0')
						trailZeros--
					}
					numStr = builder.String() + numStr
				}
				fmt.Printf("Number is %s Length %d\n", numStr, len(numStr))
				fmt.Printf("Mask is %s Length is %d\n", mask, len(mask))
				var newNum strings.Builder
				for i := 0; i < 36; i++ {
					if mask[i] == 'X' {
						newNum.WriteRune(rune(numStr[i]))
					} else {
						if numStr[i] == '0' && mask[i] != '0' {
							newNum.WriteRune('1')
						} else if numStr[i] == '0' && mask[i] == '0' {
							newNum.WriteRune('0')
						} else if numStr[i] == '1' && mask[i] != '1' {
							newNum.WriteRune('0')
						} else if numStr[i] == '1' && mask[i] == '1' {
							newNum.WriteRune('1')
						}
					}
				}
				num, _ := strconv.ParseInt(newNum.String(), 2, 64)
				table[tableIndex] = int(num)
				fmt.Printf("Result is %d %s \n", table[tableIndex], newNum.String())
				index++
			}
		} else {
			index++
		}
	}
	sum := 0
	for _, v := range table {
		sum += v
	}
	fmt.Println(sum)
}
