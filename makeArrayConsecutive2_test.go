package codesignal

import "testing"

func TestMakeArrayConsecutive2(t *testing.T) {
	type args struct {
		statues []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Valid_FourNumbers_Expects_3", args{[]int{6, 2, 3, 8}}, 3},
		{"Valid_TwoNumbers_Expects_2", args{[]int{0, 3}}, 2},
		{"Valid_ThreeNumbers_Expects_2", args{[]int{5, 4, 6}}, 0},
		{"Valid_TwoNumbers_Expects_0", args{[]int{6, 3}}, 2},
		{"Valid_OneNumber_Expects_0", args{[]int{1}}, 0},
		{"Valid_OneNumber_Expects_0", args{[]int{10}}, 0},
		{"Invalid_Nil_Slice", args{}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeArrayConsecutive2(tt.args.statues); got != tt.want {
				t.Errorf("MakeArrayConsecutive2() = %v, want %v", got, tt.want)
			}
		})
	}
}
