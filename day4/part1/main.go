package main

import "fmt"
import "io/ioutil"
import "strings"

func main() {
	fmt.Println("vim-go")
	inputFile, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(inputFile), "\n\n")
	keys := map[string]bool{
		"byr": true,
		"iyr": true,
		"eyr": true,
		"hgt": true,
		"hcl": true,
		"ecl": true,
		"pid": true,
		"cid": true,
	}
	totalCnt := 0
	for _, line := range lines {
		line = strings.Replace(line, "\n", " ", -1)
		parts := strings.Split(string(line), " ")
		cnt := 0
		hasCid := false
		for _, part := range parts {
			passport := strings.Split(string(part), ":")[0]
			if keys[passport] {
				cnt++
				if passport == "cid" {
					hasCid = true
				}
			}
		}
		if cnt == 8 {
			totalCnt++
		} else if cnt == 7 && !hasCid {
			totalCnt++
		}
	}
	fmt.Println(totalCnt)
}
