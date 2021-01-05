package p1

import (
	"reflect"
	"testing"
)

func Test_cw3d_neighbours(t *testing.T) {
	type args struct {
		c Dim
	}
	tests := []struct {
		name string
		args args
		want []Dim
	}{
		{
			name: "origo",
			args: args{
				c: Dim{0, 0, 0},
			},
			want: []Dim{
				{-1, -1, -1},
				{-1, -1, 0},
				{-1, -1, 1},
				{-1, 0, -1},
				{-1, 0, 0},
				{-1, 0, 1},
				{-1, 1, -1},
				{-1, 1, 0},
				{-1, 1, 1},
				{0, -1, -1},
				{0, -1, 0},
				{0, -1, 1},
				{0, 0, -1},
				{0, 0, 1},
				{0, 1, -1},
				{0, 1, 0},
				{0, 1, 1},
				{1, -1, -1},
				{1, -1, 0},
				{1, -1, 1},
				{1, 0, -1},
				{1, 0, 0},
				{1, 0, 1},
				{1, 1, -1},
				{1, 1, 0},
				{1, 1, 1},
			},
		},
		{
			name: "1,1,1",
			args: args{
				c: Dim{1, 1, 1},
			},
			want: []Dim{
				{0, 0, 0},
				{0, 0, 1},
				{0, 0, 2},
				{0, 1, 0},
				{0, 1, 1},
				{0, 1, 2},
				{0, 2, 0},
				{0, 2, 1},
				{0, 2, 2},
				{1, 0, 0},
				{1, 0, 1},
				{1, 0, 2},
				{1, 1, 0},
				{1, 1, 2},
				{1, 2, 0},
				{1, 2, 1},
				{1, 2, 2},
				{2, 0, 0},
				{2, 0, 1},
				{2, 0, 2},
				{2, 1, 0},
				{2, 1, 1},
				{2, 1, 2},
				{2, 2, 0},
				{2, 2, 1},
				{2, 2, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cw := &cw3d{}
			if got := cw.neighbours(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cw3d.neighbours() = %v, want %v", got, tt.want)
			}
		})
	}
}
