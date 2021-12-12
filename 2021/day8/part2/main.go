package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	//Read the file
	content, _ := ioutil.ReadFile("input.txt")
	values := strings.Split(string(content), "\n")
	sum := 0
	for _, v := range values {
		parts := strings.Split(v, "|")
		if len(parts) == 2 {
			lengthToStr := lengthToStr(parts)
			numbers := sortedNumbers(parts[1])
			strToLength := make(map[string]int)
			for key, value := range lengthToStr {
				strToLength[value] = key
			}
			var builder strings.Builder
			for _, number := range numbers {
				value, exists := strToLength[number]
				if exists {
					builder.WriteString(fmt.Sprintf("%d", value))
				}
			}
			num, _ := strconv.Atoi(builder.String())
			sum += num
		}
	}
	fmt.Println(sum)
}

func lengthToStr(parts []string) map[int]string {
	numbers := sortedNumbers(parts[0])
	//sort by length
	sort.Strings(numbers)
	lengthToStr := make(map[int]string)
	sixes := []string{}
	fives := []string{}
	for _, number := range numbers {
		if len(number) == 2 {
			lengthToStr[1] = number
		} else if len(number) == 3 {
			lengthToStr[7] = number
		} else if len(number) == 4 {
			lengthToStr[4] = number
		} else if len(number) == 7 {
			lengthToStr[8] = number
		} else if len(number) == 5 {
			fives = append(fives, number)
		} else if len(number) == 6 {
			sixes = append(sixes, number)
		}
	}
	one := lengthToStr[1]
	onesSet := make(map[byte]bool)
	for i := 0; i < len(one); i++ {
		onesSet[one[i]] = true
	}
	four := lengthToStr[4]
	fourSet := make(map[byte]bool)
	for i := 0; i < len(four); i++ {
		fourSet[four[i]] = true
	}
	//assign 0,6,9
	for _, lengthSixNumber := range sixes {
		overlapOne := 0
		overlapFour := 0
		for i := 0; i < len(lengthSixNumber); i++ {
			_, exists := onesSet[lengthSixNumber[i]]
			if exists {
				overlapOne++
			}
			_, exists = fourSet[lengthSixNumber[i]]
			if exists {
				overlapFour++
			}
		}
		if overlapOne == 1 {
			lengthToStr[6] = lengthSixNumber
		} else if overlapFour == 4 {
			lengthToStr[9] = lengthSixNumber
		} else {
			lengthToStr[0] = lengthSixNumber
		}
	}
	six := lengthToStr[6]
	sixSet := make(map[byte]bool)
	for i := 0; i < len(six); i++ {
		sixSet[six[i]] = true
	}
	for _, lengthFiveNumber := range fives {
		overlapOne := 0
		overlaSix := 0
		for i := 0; i < len(lengthFiveNumber); i++ {
			_, exists := onesSet[lengthFiveNumber[i]]
			if exists {
				overlapOne++
			}
			_, exists = sixSet[lengthFiveNumber[i]]
			if exists {
				overlaSix++
			}
		}
		if overlapOne == 2 {
			lengthToStr[3] = lengthFiveNumber
		} else if overlaSix == 5 {
			lengthToStr[5] = lengthFiveNumber
		} else {
			lengthToStr[2] = lengthFiveNumber
		}
	}
	return lengthToStr
}

func sortedNumbers(line string) []string {
	numbers := strings.Split(line, " ")
	sorted := []string{}
	for _, word := range numbers {
		s := []rune(word)
		sort.Slice(s, func(i int, j int) bool { return s[i] < s[j] })
		sorted = append(sorted, string(s))
	}
	return sorted
}
