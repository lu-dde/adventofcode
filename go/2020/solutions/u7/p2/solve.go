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

	for line := range p {
		contain := strings.Split(line, " bags contain ")
		color := contain[0]
		containsString := contain[1]
		if containsString != "no other bags." {
			containsArr := strings.Fields(containsString)

			contains := map[string]int{}
			for i := 0; i < len(containsArr); i += 4 {
				nr, _ := strconv.Atoi(containsArr[i])
				col := containsArr[i+1] + " " + containsArr[i+2]
				contains[col] = nr
			}

			bags[color] = contains
		}

	}
	//fmt.Println(bags)
	//fmt.Println(containedBy)
	//fmt.Println()

	index := 0

	parents := []string{"shiny gold"}

	for {
		if index == len(parents) {
			break
		}

		parent := parents[index]

		for bag, amount := range bags[parent] {
			for i := 0; i < amount; i++ {
				parents = append(parents, bag)
			}
		}

		index++
	}
	//fmt.Println(len(parents), parents)

	t = len(parents) - 1

	s <- fmt.Sprintf("Solution: %d", t)
}
