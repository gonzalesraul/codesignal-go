package codesignal

import "testing"

func TestAlmostIncreasingSequence(t *testing.T) {
	type args struct {
		sequence []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Valid_ScrumbleSlice_Expects_false", args{[]int{1, 3, 2, 1}}, false},
		{"Valid_OneTwoOneTwo_Expects_false", args{[]int{1, 2, 1, 2}}, false},
		{"Valid_ChainSequences_Expects_false", args{[]int{40, 50, 60, 10, 20, 30}}, false},
		{"Valid_ContainsRepeated_Expects_false", args{[]int{1, 4, 10, 4, 2}}, false},
		{"Valid_OneThreeTwo_Expects_true", args{[]int{1, 3, 2}}, true},
		{"Valid_BiggerAtStart_Expects_true", args{[]int{10, 1, 2, 3, 4, 5}}, true},
		{"Valid_NegativeInTheMiddle_Expects_true", args{[]int{0, -2, 5, 6}}, true},
		{"Valid_BiggerInTheMiddle_Expects_true", args{[]int{1, 2, 3, 4, 99, 5, 6}}, true},
		{"Valid_RepeatedNumbersAtEnd_Expects_false", args{[]int{1, 2, 5, 5, 5}}, false},
		{"Valid_RepeatedNumbersAtStart_Expects_false", args{[]int{1, 1, 1, 2, 3}}, false},
		{"Valid_HugeGap_Expects_true", args{[]int{3, 5, 67, 98, 3}}, true},
		{"Invalid_Nil_Slice_Expects_false", args{}, false},
		{"Invalid_Length_Lt+2_Expects_false", args{[]int{1}}, false},
		{"Invalid_Length_Gt+105_Expects_false", args{make([]int, 106)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AlmostIncreasingSequence(tt.args.sequence); got != tt.want {
				t.Errorf("AlmostIncreasingSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
