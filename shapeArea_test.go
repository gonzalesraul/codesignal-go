package codesignal

import "testing"

func TestShapeArea(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Valid_Expects_5", args{2}, 5},
		{"Valid_Expects_13", args{3}, 13},
		{"Valid_Expects_41", args{5}, 41},
		{"Invalid_Input_Gt+10000", args{10001}, 0},
		{"Invalid_Input_Lt-1", args{-1}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ShapeArea(tt.args.n); got != tt.want {
				t.Errorf("ShapeArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
