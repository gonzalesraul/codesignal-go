package codesignal

import "testing"

func TestMatrixElementsSum(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Valid_HauntedRoomsEachLevel_Expects_9", args{[][]int{{0, 1, 1, 2}, {0, 5, 0, 0}, {2, 0, 3, 3}}}, 9},
		{"Valid_NoHauntedRoomsAtLastLevel_Expects_9", args{[][]int{{1, 1, 1, 0}, {0, 5, 0, 1}, {2, 1, 3, 10}}}, 9},
		{"Valid_NoHauntedRoomsAtAll_Expects_18", args{[][]int{{1, 1, 1}, {2, 2, 2}, {3, 3, 3}}}, 18},
		{"Valid_AHauntedRoom_Expects_0", args{[][]int{{0}}}, 0},
		{"Valid_HauntedRoof_Expects_0", args{[][]int{{0, 0, 0}, {1, 1, 1}}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MatrixElementsSum(tt.args.matrix); got != tt.want {
				t.Errorf("MatrixElementsSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
