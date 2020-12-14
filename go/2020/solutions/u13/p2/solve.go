package p2

import (
	"fmt"
	"strconv"
	"strings"
)

//Solve is main proxy for solve, takes a string channel
func Solve(p chan string, s chan string) {
	var t uint64 = 0

	for line := range p {
		t = solve(line)
	}

	//	fmt.Println(t)

	s <- fmt.Sprintf("Solution: %d", t)
}

func solve(bussesStr string) uint64 {

	busses := []uint64{}

	for _, bidStr := range strings.Split(bussesStr, ",") {
		bid, err := strconv.ParseUint(bidStr, 10, 64)
		if err != nil {
			bid = 1
		}
		busses = append(busses, bid)
	}

	//busses = busses[0:4]

	var pos, skip, time uint64 = busses[0], 1, 0
	var lo uint64

	for offset, bid := range busses {
		o := uint64(offset)
		diff := o - lo

		if bid == 1 {
			pos++
			lo = o
			//	fmt.Println("< found", pos, skip, bid)
			continue
		}

		//fmt.Println("time", time, "skip", skip, "diff", diff, "next", bid)
		for time = pos; ; time += skip {
			//fmt.Println("  test", time+diff, "%", bid, "=", (time+diff)%bid)
			if (time+diff)%bid == 0 {
				pos = time + diff
				lo = o
				skip *= bid
				//fmt.Println("< found", pos, skip, bid)
				break
			}
		}
	}

	//fmt.Println(pos - lo)

	return pos - lo
}

/*
   17,x,13,19 = 3417.
   67,7,59,61 = 754018.
   67,x,7,59,61 = 779210.
   67,7,x,59,61 = 1261476.
   1789,37,47,1889 = 1202161486.
*/
