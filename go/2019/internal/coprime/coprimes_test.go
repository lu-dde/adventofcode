package coprime

import (
	"testing"

	"github.com/lu-dde/adventofcode/internal/coord"
)

func Test_GetCoprimes_limitM10N10(t *testing.T) {
	cop := New(40, 40)

	coprimesList := cop.GetCoprimes()

	expectedCoprimes := coord.Slice{{X: 2, Y: 1}, {X: 3, Y: 1}, {X: 3, Y: 2}, {X: 4, Y: 3}, {X: 5, Y: 3}, {X: 5, Y: 4}, {X: 6, Y: 5}, {X: 7, Y: 5}, {X: 7, Y: 6}, {X: 8, Y: 7}, {X: 9, Y: 7}, {X: 9, Y: 8}, {X: 10, Y: 9}, {X: 11, Y: 9}, {X: 11, Y: 10}, {X: 12, Y: 11}, {X: 13, Y: 11}, {X: 13, Y: 12}, {X: 14, Y: 13}, {X: 15, Y: 13}, {X: 15, Y: 14}, {X: 16, Y: 15}, {X: 17, Y: 15}, {X: 17, Y: 16}, {X: 18, Y: 17}, {X: 19, Y: 17}, {X: 19, Y: 18}, {X: 20, Y: 19}, {X: 21, Y: 19}, {X: 21, Y: 20}, {X: 23, Y: 21}, {X: 25, Y: 23}, {X: 27, Y: 25}, {X: 29, Y: 27}, {X: 31, Y: 29}, {X: 33, Y: 31}, {X: 35, Y: 33}, {X: 37, Y: 35}, {X: 39, Y: 37}}

	if len(coprimesList) != len(expectedCoprimes) {
		t.Error("lists must have equal length")
	}

	for i, expected := range expectedCoprimes {
		found := coprimesList[i]
		if expected != found {
			t.Errorf("list must be equal %d %v != %v", i, expected, found)
		}
	}
}
