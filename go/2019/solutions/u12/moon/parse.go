package moon

import (
	"strconv"
	"strings"
)

//ParseStr parsen a string to a moon
func ParseStr(line string) *Moon {
	line = line[1 : len(line)-1] // with out <..> wrap
	vec := [3]float64{}
	for i, c := range strings.Split(line, ", ") {
		c = c[2:]
		value, err := strconv.ParseFloat(c, 64)
		if err != nil {
			panic(err)
		}
		vec[i] = value
	}

	return &Moon{
		Position: vec,
	}
}
