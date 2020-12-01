package alch

import "strings"

//Reaction contains a result and which materials creates it
type Reaction struct {
	Reagents []Material
	Product  Material
}

func parseReactionStr(line string) Reaction {
	l := strings.Split(line, " => ")
	product := parseMaterialStr(l[1])
	rStr := strings.Split(l[0], ", ")
	reagents := []Material{}
	for _, r := range rStr {
		reagents = append(reagents, parseMaterialStr(r))
	}

	return Reaction{
		Reagents: reagents,
		Product:  product,
	}
}
