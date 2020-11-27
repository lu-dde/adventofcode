package moon

import (
	"math"

	"github.com/go-gl/mathgl/mgl64"
)

//Moon object
type Moon struct {
	Position mgl64.Vec3
	Velocity mgl64.Vec3
}

//Reset Volocity to zero
func (m *Moon) Reset() {
	m.Velocity = mgl64.Vec3{}
}

//Gravitate pulls m in the direction of other
func (m *Moon) Gravitate(other *Moon) {
	m.Velocity[0] += grav(m.Position[0], other.Position[0])
	m.Velocity[1] += grav(m.Position[1], other.Position[1])
	m.Velocity[2] += grav(m.Position[2], other.Position[2])
}

//Move the moon with it's Velocity
func (m *Moon) Move() {
	m.Position = m.Position.Add(m.Velocity)
}

//Energy get to total potential and kinetic energy of the moon
func (m *Moon) Energy() float64 {
	var sumP, sumV float64
	sumP += math.Abs(m.Position[0])
	sumP += math.Abs(m.Position[1])
	sumP += math.Abs(m.Position[2])
	sumV += math.Abs(m.Velocity[0])
	sumV += math.Abs(m.Velocity[1])
	sumV += math.Abs(m.Velocity[2])
	return sumP * sumV
}

func grav(p, o float64) float64 {
	switch {
	case p < o:
		return 1
	case p > o:
		return -1
	default:
		return 0
	}
}
