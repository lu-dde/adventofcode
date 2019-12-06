package main

import (
	"fmt"
	"strconv"
	"strings"
)

//U41 is main proxy for solve, takes a string channel
func U41(p chan string, s chan string) {

	line, _ := <-p
	codes := strings.Split(line, "")

	var low = []int{0, 0, 0, 0, 0, 0, 0}
	var high = []int{0, 0, 0, 0, 0, 0, 0}

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

		if passedPass(&low, &high) {
			break loop
		}
		valid++
	}

	s <- fmt.Sprintf("Solution: %d", valid)
}

func nextValidFormatPass(pass *[]int) bool {
	var carry = false
	var c = 6
	for {
		max := maxI32((*pass)[c], (*pass)[c-1])
		if (*pass)[c] < max {
			(*pass)[c] = max
		} else {
			(*pass)[c]++
			carry = (*pass)[c] > 9
		}
		if carry {
			c--
		} else {
			// reset lower to mine
			for i := c; i < len(*pass); i++ {
				(*pass)[i] = (*pass)[c]
			}
			break
		}
	}
	var prev = -1
	for _, v := range *pass {
		if prev == v {
			return true
		}
		prev = v
	}
	return false
}

func passedPass(pass *[]int, other *[]int) bool {
	for i := range *pass {
		p := (*pass)[i]
		o := (*other)[i]
		if p > o {
			return true
		} else if p < o {
			return false
		}
	}
	return false
}

func maxI32(m, n int) int {
	if m > n {
		return m
	}
	return n
}

func boolName(b bool) string {
	if b {
		return "true"
	}
	return "false"
}
