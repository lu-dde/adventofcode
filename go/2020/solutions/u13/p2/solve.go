package p2

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

//Solve is main proxy for solve, takes a string channel
func Solve(p chan string, s chan string) {
	var t = 0

	timeStr := <-p
	bussesStr := <-p

	time, _ := strconv.Atoi(timeStr)

	busses := []int{}

	for _, bidStr := range strings.Split(bussesStr, ",") {
		bid, err := strconv.Atoi(bidStr)
		if err != nil {
			bid = math.MaxInt64
		}
		busses = append(busses, bid)
	}

	fmt.Println(time)
	fmt.Println(bussesStr)
	//fmt.Println(busses)

	min := math.MaxInt64
	id := 0
	for _, d := range busses {
		diff := (d*(time/d) - time + d)
		if diff < min {
			min = diff
			id = d
		}
	}

	t = min * id

	s <- fmt.Sprintf("Solution: %d", t)
}
