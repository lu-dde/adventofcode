package p2

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

		//	fmt.Println(line)
		for {

			l := penParents.FindAllString(line, 10)
			if len(l) == 0 {
				break
			}

			for _, sub := range l {
				s := fmt.Sprintf("%d", eval(sub[1:len(sub)-1]))
				line = strings.Replace(line, sub, s, -1)
				//fmt.Println(" ", line)
			}
		}

		t += eval(line)

	}

	s <- fmt.Sprintf("Solution: %d", t)
}

func eval(s string) (total int) {
	//s = 1 + 2 * 3 => 9 odd presedence
	l := strings.Fields(s)
	i := 0
	for {
		if len(l) <= i {
			break
		}
		if l[i] == plus {
			p, _ := strconv.Atoi(l[i-1])
			q, _ := strconv.Atoi(l[i+1])
			l[i-1] = fmt.Sprint(p + q)
			//	fmt.Println("  ", l[i-1])
			new := l[0:i]
			if i+1 < len(l) {
				rest := l[i+2:]
				new = append(new, rest...)
			}
			l = new
		} else {
			i++
		}
		//fmt.Println("  ", l)
	}

	total = 1

	for _, v := range l {
		if v != mult {
			t, _ := strconv.Atoi(v)
			total *= t
		}
	}

	//	fmt.Println("eval", s, "=", total)

	return
}
