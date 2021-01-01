package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type validRange struct {
	from int
	to   int
}

func parseRanges(first, second string) validRange {
	from, _ := strconv.Atoi(first)
	to, _ := strconv.Atoi(second)
	return validRange{from: from, to: to}
}

func (r validRange) isValid(number int) bool {
	return number >= r.from && number <= r.to
}

func isValid(number int, ranges []validRange) bool {
	for _, r := range ranges {
		if r.isValid(number) {
			return true
		}
	}
	return false
}

func main() {
	r := regexp.MustCompile("([0-9]+)-([0-9]+)")
	content, _ := ioutil.ReadFile("input")
	lines := strings.Split(string(content), "\n\n")
	ranges := []validRange{}
	rangesLines := strings.Split(lines[0], "\n")
	for _, line := range rangesLines {
		res := r.FindAllStringSubmatch(line, -1)
		ranges = append(ranges, parseRanges(res[0][1], res[0][2]))
		ranges = append(ranges, parseRanges(res[1][1], res[1][2]))
	}

	ticketLines := strings.Split(string(lines[2]), "\n")
	sum := 0
	for i, tickets := range ticketLines {
		if i != 0 {
			lineWithTickets := strings.Split(string(tickets), ",")
			for _, ticket := range lineWithTickets {
				numTicket, _ := strconv.Atoi(ticket)
				if !isValid(numTicket, ranges) {
					sum += numTicket
				}
			}
		}
	}
	fmt.Println(sum)

}
