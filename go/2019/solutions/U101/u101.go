package u101

import (
	"fmt"
)

//Solve U101
func Solve(p chan string, s chan string) {
	var t = 0

	/* starmap # is an astroid
	#.#...#.#.
	.###....#.
	.#....#...
	##.#.#.#.#
	....#.#.#.
	.##..###.#
	..#...##..
	..##....##
	......#...
	.####.###.
	*/
	for line := range p {
		fmt.Println(line)
	}

	s <- fmt.Sprintf("Solution: %d", t)
}
