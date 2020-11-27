package arcade

import "github.com/lu-dde/adventofcode/internal/coord"

type tracker struct {
	ball       int
	paddle     int
	paddleSeen bool
	ballSeen   bool
}

func (t *tracker) track(c coord.Pair, oid int) {
	if oid == objBall {
		t.trackBall(c)
	}
	if oid == objPaddle {
		t.trackPaddle(c)
	}
}

func (t *tracker) trackBall(c coord.Pair) {
	t.ballSeen = true
	t.ball = c.X
}
func (t *tracker) trackPaddle(c coord.Pair) {
	t.paddleSeen = true
	t.paddle = c.X
}

func (t *tracker) getDirection() int64 {
	switch {
	case t.paddle < t.ball:
		return joystickRight
	case t.paddle > t.ball:
		return joystickLeft
	default:
		return joystickNet
	}
}
