package p1

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

//Solve is main proxy for solve, takes a string channel
func Solve(p chan string, s chan string) {
	start := time.Now()

	var t = 0

	numsString := strings.Split(<-p, ",")

	last := 0

	spoken := map[int][]int{}
	for turn, ns := range numsString {
		n, _ := strconv.Atoi(ns)
		spoken[n] = append(spoken[n], turn)
		spoken[n] = append(spoken[n], turn)
		last = n
	}

	for turn := len(numsString); turn < 2020; turn++ {
		l := spoken[last]
		firstSpokenTurn := l[0]
		prevSpokenturn := l[len(l)-2]
		lastSpokenturn := l[len(l)-1]
		//fmt.Println("turn", turn, "last", last, firstSpokenTurn, prevSpokenturn, lastSpokenturn)
		if lastSpokenturn == firstSpokenTurn {
			spoken[0] = append(spoken[0], turn)
			last = 0
		} else {
			last = lastSpokenturn - prevSpokenturn
			_, notFirst := spoken[last]
			if !notFirst {
				spoken[last] = append(spoken[last], turn)
			}
			spoken[last] = append(spoken[last], turn)
		}
	}
	//fmt.Println(spoken)

	t = last

	s <- fmt.Sprintf("Solution: %d", t)

	fmt.Println("end", time.Since(start))
}
