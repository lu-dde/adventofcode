package hullpainter

import (
	"reflect"
	"testing"

	"github.com/lu-dde/adventofcode/internal/coord"
)

func Test_rotate(t *testing.T) {
	type args struct {
		rotate int64
		p      coord.Pair
	}
	tests := []struct {
		name string
		args args
		want coord.Pair
	}{

		// leftRotate
		{
			name: "rotate 0 north to west",
			args: args{rotate: 0, p: north},
			want: west,
		},
		{
			name: "rotate 0 west to south",
			args: args{rotate: 0, p: west},
			want: south,
		},
		{
			name: "rotate 0 south to east",
			args: args{rotate: 0, p: south},
			want: east,
		},
		{
			name: "rotate 0 east to north",
			args: args{rotate: 0, p: east},
			want: north,
		},

		// rightRotate
		{
			name: "rotate 1 north to east",
			args: args{rotate: 1, p: north},
			want: east,
		},
		{
			name: "rotate 1 east to south",
			args: args{rotate: 1, p: north},
			want: east,
		},
		{
			name: "rotate 1 south to west",
			args: args{rotate: 1, p: north},
			want: east,
		},
		{
			name: "rotate 1 west to north",
			args: args{rotate: 1, p: north},
			want: east,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotate(tt.args.rotate, tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotate() = %v, want %v", got, tt.want)
			}
		})
	}
}
