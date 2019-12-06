package main

import (
	"fmt"
	"strconv"
	"strings"
)

//U42 is main proxy for solve, takes a string channel
func U42(p chan string, s chan string) {

	line, _ := <-p
	codes := strings.Split(line, "")

	var low = []int{-1, 0, 0, 0, 0, 0, 0}
	var high = []int{-1, 0, 0, 0, 0, 0, 0}

	for i, v := range codes[0:6] {
		d, _ := strconv.Atoi(v)
		low[i+1] = d
	}
	for i, v := range codes[7:13] {
		d, _ := strconv.Atoi(v)
		high[i+1] = d
	}
	var valid = 0

loop:
	for {

		if !nextValidFormatPass(&low) {
			continue
		}

		counts := repeatCountInPass(&low)

		if !counts[2] {
			continue
		}

		if passedPass(&low, &high) {
			break loop
		}
		valid++
	}

	s <- fmt.Sprintf("Solution: %d", valid)
}

func repeatCountInPass(pass *[]int) []bool {
	var counts = []bool{false, false, false, false, false, false, false}
	var count = 0
	var prev = -1
	for _, v := range *pass {
		if prev != v {
			counts[count] = true
			count = 0
		}
		count++
		prev = v
	}
	counts[count] = true
	return counts
}
