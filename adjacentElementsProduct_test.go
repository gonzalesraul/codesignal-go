package codesignal

import "testing"

func TestAdjacentElementsProduct(t *testing.T) {
	type args struct {
		inputArray []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Valid_Expects_21", args{[]int{3, 6, -2, -5, 7, 3}}, 21},
		{"Valid_Expects_2", args{[]int{-1, -2}}, 2},
		{"Valid_Expects_6", args{[]int{5, 1, 2, 3, 1, 4}}, 6},
		{"Valid_Expects_50", args{[]int{9, 5, 10, 2, 24, -1, -48}}, 50},
		{"Valid_Expects_-12", args{[]int{-23, 4, -3, 8, -12}}, -12},
		{"Valid_Expects_0", args{[]int{1, 0, 1, 0, 1000}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AdjacentElementsProduct(tt.args.inputArray); got != tt.want {
				t.Errorf("AdjacentElementsProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
