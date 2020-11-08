package grav

import (
	"strconv"
	"strings"
)

//ParseStr parsen a string to a moon
func (gs *Set) ParseStr(line string) {
	line = line[1 : len(line)-1] // with out <..> wrap
	for axis, c := range strings.Split(line, ", ") {
		c = c[2:]
		if gs.axis == axis {
			value, err := strconv.ParseInt(c, 10, 16)
			if err != nil {
				panic(err)
			}
			gs.addPos(int(value))
		}
	}
}

func (gs *Set) addPos(v int) {
	gs.positions = append(gs.positions, v)
	gs.velocity = append(gs.velocity, 0)
}
