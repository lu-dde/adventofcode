package p2

import (
	"fmt"
	"math"
)

//Solve is main proxy for solve, takes a string channel
func Solve(p chan string, s chan string) {

	var total float64 = 0
	var min float64 = math.Inf(0)
	var max float64 = 0
	var nr float64 = 1

	for line := range p {
		seatID := 0
		for i := 0; i < 10; i++ {
			seatID <<= 1
			c := line[i]
			if c == 'B' || c == 'R' {
				seatID++
			}
		}

		//fmt.Println(seatID, line)
		fid := float64(seatID)
		total += fid
		min = math.Min(fid, min)
		max = math.Max(fid, max)
		nr++
		//fmt.Println(seatID)
	}

	//fmt.Println("(", min, "+", max, ") *", nr, "/", 2, "-", total, "=", (min+max)*nr/2-total)

	s <- fmt.Sprintf("Solution: %.0f", (min+max)*nr/2-total)
}
