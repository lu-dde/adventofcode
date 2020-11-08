package u102

import (
	"math"
	"testing"

	"github.com/lu-dde/adventofcode/internal/coord"
)

func Test_angle(t *testing.T) {
	type args struct {
		distance coord.Pair
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "vector at 45 degrees right from top",
			args: args{coord.Pair{X: 1, Y: -1}},
			want: math.Pi / 4,
		},
		{
			name: "vector at 90 degrees right from top",
			args: args{coord.Pair{X: 1, Y: 0}},
			want: math.Pi / 2,
		},
		{
			name: "vector at 135 degrees right from top",
			args: args{coord.Pair{X: 1, Y: 1}},
			want: math.Pi / 4 * 3,
		},
		{
			name: "vector at 180 degrees right from top",
			args: args{coord.Pair{X: 0, Y: 1}},
			want: math.Pi / 4 * 4,
		},
		{
			name: "vector at 315 degrees right from top",
			args: args{coord.Pair{X: -1, Y: -1}},
			want: math.Pi / 4 * 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := angle(tt.args.distance); got != tt.want {
				t.Errorf("angle() = %v, want %v", got, tt.want)
			}
		})
	}
}
