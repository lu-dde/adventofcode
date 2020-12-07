package p1

import (
	"fmt"
	"strconv"
	"strings"
)

//Solve is main proxy for solve, takes a string channel
func Solve(p chan string, s chan string) {
	var t = 0

	containedBy := map[string][]string{}

	for line := range p {
		contain := strings.Split(line, " bags contain ")
		color := contain[0]
		containsString := contain[1]
		if containsString != "no other bags." {
			//fmt.Println(color)
			containsArr := strings.Fields(containsString)

			contains := map[string]int{}
			for i := 0; i < len(containsArr); i += 4 {
				nr, _ := strconv.Atoi(containsArr[i])
				col := containsArr[i+1] + " " + containsArr[i+2]
				contains[col] = nr
				containedBy[col] = append(containedBy[col], color)
			}
		}

	}

	index := 0
	parents := []string{"shiny gold"}
	count := map[string]bool{}
	for {
		if index == len(parents) {
			break
		}

		parent := parents[index]
		count[parent] = true
		parents = append(parents, containedBy[parent]...)

		index++
	}

	t = len(count) - 1

	s <- fmt.Sprintf("Solution: %d", t)
}
