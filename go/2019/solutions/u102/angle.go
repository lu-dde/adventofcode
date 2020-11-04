package u102

import (
	"math"

	"github.com/lu-dde/adventofcode/internal/coord"
)

func angle(distance coord.Pair) float64 {
	//orientation := coord.NewPair(0, -1)

	orientationAngle := math.Atan2(1, 0) // 90deg pi/2 ?
	distanceAngle := math.Atan2(-float64(distance.Y), float64(distance.X))

	//FIXME: räknar fel här
	angle := orientationAngle - distanceAngle
	if angle < 0 {
		angle += math.Pi * 2
	}
	//	fmt.Println(distance, "D", distanceAngle, distanceAngle/(2*math.Pi)*360, "O", orientationAngle, "= A", angle)

	return angle
}
