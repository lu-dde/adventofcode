package p1

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const plus string = "+"
const mult string = "*"

//Solve is main proxy for solve, takes a string channel
func Solve(p chan string, s chan string) {
	var t = 0

	penParents := regexp.MustCompile("\\([^()]+\\)")

	for line := range p {

		for {

			l := penParents.FindAllString(line, 10)
			if len(l) == 0 {
				break
			}

			for _, sub := range l {
				s := fmt.Sprintf("%d", eval(sub[1:len(sub)-1]))
				line = strings.Replace(line, sub, s, -1)
			}
		}

		t += eval(line)

	}

	s <- fmt.Sprintf("Solution: %d", t)
}

func eval(s string) (total int) {
	//s = 1 + 2 * 3 => 9 odd presedence

	op := plus
	for _, v := range strings.Fields(s) {
		switch v {
		case plus:
			op = plus
		case mult:
			op = mult
		default:
			value, _ := strconv.Atoi(v)
			switch op {
			case plus:
				total += value
			case mult:
				total *= value
			}
		}
	}

	return
}
