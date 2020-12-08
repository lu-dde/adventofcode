package handheld

import "github.com/lu-dde/adventofcode/internal/ops"

//GetConsole converts a string channel to a Console
func GetConsole(input <-chan string) Console {
	instructions := []ops.Cmd{}

	for line := range input {
		instructions = append(instructions, ops.New(line))
	}

	return Console{
		Ops: instructions,
		Fin: len(instructions),
	}
}
