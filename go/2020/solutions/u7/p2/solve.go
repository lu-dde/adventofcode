package p2

import (
	"fmt"
	"strconv"
	"strings"
)

//Solve is main proxy for solve, takes a string channel
func Solve(p chan string, s chan string) {
	var t = 0

	bags := map[string]map[string]int{}

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

			bags[color] = contains

			//fmt.Println("", contains)
		}

	}
	//fmt.Println(bags)
	//fmt.Println(containedBy)
	//fmt.Println()

	index := 0

	parents := []string{"shiny gold"}

	count := map[string]bool{}
	for {
		if index == len(parents) {
			break
		}

		parent := parents[index]

		count[parent] = true

		moreParents := containedBy[parent]

		index++

		parents = append(parents, moreParents...)

		//fmt.Println(len(count), parent, "is contained by:", strings.Join(moreParents, ","))

	}
	//fmt.Println(len(count), count)

	t = len(count) - 1

	s <- fmt.Sprintf("Solution: %d", t)
}
