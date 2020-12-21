package p1

import (
	"fmt"
	"strconv"
	"strings"
)

//Solve is main proxy for solve, takes a string channel
func Solve(p chan string, s chan string) {
	var t = 0

	class := map[string][2][2]int{}

	for line := range p {
		if line == "" {
			break
		}
		addClass(class, line)
	}

	<-p // skip your ticket
	<-p // skip your ticket
	<-p // skip your ticket
	/*
		var yourTicket []int

		<-p // your ticket
		for line := range p {
			if line == "" {
				break
			}
			yourTicket = getTicket(line)
		}
	*/

	var nearbyTickets [][]int = [][]int{}

	<-p // nearby tickets
	for line := range p {
		nearbyTickets = append(nearbyTickets, getTicket(line))
	}

	//	fmt.Println(class)

	for _, ticket := range nearbyTickets {
		errors := scanTicket(class, ticket)
		//fmt.Println(ticket, errors)
		for _, v := range errors {
			t += v
		}
	}

	s <- fmt.Sprintf("Solution: %d", t)
}

func scanTicket(m map[string][2][2]int, ticket []int) []int {
	failed := []int{}
	for _, t := range ticket {
		success := false
	rm:
		for _, rr := range m {
			for _, r := range rr {
				//fmt.Println("try", t, "range", r[0], "-", r[1], "=", r[0] <= t && t <= r[1])
				if r[0] <= t && t <= r[1] {
					success = true
					break rm
				}
			}
		}
		if !success {
			failed = append(failed, t)
		}
	}
	return failed
}

func addClass(m map[string][2][2]int, line string) {
	var name string
	var r11, r12, r21, r22 int
	fmt.Sscanf(line, "%s %d-%d or %d-%d", &name, &r11, &r12, &r21, &r22)

	m[name] = [2][2]int{
		{r11, r12},
		{r21, r22},
	}
}

func getTicket(line string) []int {
	s := strings.Split(line, ",")
	ticket := make([]int, len(s))
	for i, s := range s {
		ticket[i], _ = strconv.Atoi(s)
	}
	return ticket
}
