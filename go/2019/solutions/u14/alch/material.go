package alch

import (
	"strconv"
	"strings"
)

//Material contains an element key and amount
type Material struct {
	Element string
	Amount  int
}

func parseMaterialStr(line string) Material {
	l := strings.Split(line, " ")
	amount := l[0]
	element := l[1]

	value, _ := strconv.Atoi(amount)

	return Material{
		Element: element,
		Amount:  value,
	}
}
