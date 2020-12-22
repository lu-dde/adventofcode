package p2

import (
	"fmt"
	"strconv"
	"strings"
)

//Solve is main proxy for solve, takes a string channel
func Solve(p chan string, s chan string) {
	var t = 0

	class := map[string][2][2]int{}

	nameIndex := []string{}
	for line := range p {
		if line == "" {
			break
		}
		nameIndex = append(nameIndex, addClass(class, line))
	}

	<-p // your ticket:
	var yourTicket []int = getTicket(<-p)

	<-p // empty line:
	<-p // nearby tickets:

	var nearbyTickets [][]int = [][]int{}

	for line := range p {
		nearbyTickets = append(nearbyTickets, getTicket(line))
	}

	//	fmt.Println(class)

	var valid [][]int = [][]int{
		yourTicket,
	}

	for _, ticket := range nearbyTickets {
		if scanTicket(class, ticket) {
			valid = append(valid, ticket)
		}
	}

	match := matchColumns(class, valid)
	//fmt.Println(match)
	var mul int = 1
	for _, s := range nameIndex {
		if strings.HasPrefix(s, "departure") {
			mul *= yourTicket[match[s]]
		}
	}

	t = mul

	s <- fmt.Sprintf("Solution: %d", t)
}

func matchColumns(m map[string][2][2]int, tickets [][]int) map[string]int {

	fields := len(m)

	//FIXME CLAIM!
	claimC := map[int]string{}
	claimN := map[string]int{}
	invalid := make(map[int]map[string]bool, fields)

	for k := 0; k < fields; k++ {
		if len(claimC) == fields {
			break
		}
		//fmt.Println("for Claim", len(claimC), fields)

		for i := 0; i < fields; i++ {
			if _, ok := claimC[i]; ok {
				//fmt.Println("  skip claimed by", i, claimC[i])
				continue
			}
			if invalid[i] == nil {
				invalid[i] = map[string]bool{}
			}
			//fmt.Println("  for Field", i)

			for n, rr := range m {
				if _, ok := claimN[n]; ok {
					if _, ok := invalid[i][n]; ok {
						delete(invalid[i], n)
					}
					//fmt.Println("    skip claimed by", n, claimN[n])
					continue
				}
				//fmt.Println("    for Class", n)
				invalid[i][n] = true
				for _, t := range tickets {
					////fmt.Println("      for Ticket", rr, t[i])

					if !inRange(rr, t[i]) {
						//	//fmt.Println("        invalid")

						delete(invalid[i], n)
						break
					}
				}
			}
			//fmt.Println("    field left", invalid[i])
			if len(invalid[i]) == 1 {
				for n := range invalid[i] {
					//fmt.Println("    claim", i, n)
					claimC[i] = n
					claimN[n] = i
				}
			}
		}
		//fmt.Println("left", invalid)

		for n, k1 := range claimN {
			for k2, i := range invalid {
				v, ok := i[n]
				if ok && v == false && k1 != k2 {
					//fmt.Println("delete others", k2, n)
					delete(i, n)
				}
			}
		}
		for i, nm := range invalid {
			if len(nm) == 1 {
				for n := range nm {
					//fmt.Println("extra claim", i, n)
					claimC[i] = n
					claimN[n] = i
				}
			}
		}
		//fmt.Println("claims", claimC)
		//fmt.Println("claims", claimN)
	}
	//fmt.Println(claimC)

	return claimN
}

func inRange(rr [2][2]int, t int) bool {
	//fmt.Println("try", t, "range", rr[0], "or", rr[1], (rr[0][0] <= t && t <= rr[0][1]) || (rr[1][0] <= t && t <= rr[1][1]))
	return (rr[0][0] <= t && t <= rr[0][1]) || (rr[1][0] <= t && t <= rr[1][1])
}

func scanTicket(m map[string][2][2]int, ticket []int) bool {
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
			return false
		}
	}
	return true
}

func addClass(m map[string][2][2]int, line string) string {
	sp := strings.Split(line, ": ")
	var name string = sp[0]
	var r11, r12, r21, r22 int
	fmt.Sscanf(sp[1], "%d-%d or %d-%d", &r11, &r12, &r21, &r22)

	m[name] = [2][2]int{
		{r11, r12},
		{r21, r22},
	}

	return name
}

func getTicket(line string) []int {
	s := strings.Split(line, ",")
	ticket := make([]int, len(s))
	for i, s := range s {
		ticket[i], _ = strconv.Atoi(s)
	}
	return ticket
}
