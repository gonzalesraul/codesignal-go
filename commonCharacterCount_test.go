package codesignal

import "testing"

func TestCommonCharacterCount(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Valid_aabcc_adcaa_Expects_3", args{"aabcc", "adcaa"}, 3},
		{"Valid_zzzz_zzzzzzz_Expects_4", args{"zzzz", "zzzzzzz"}, 4},
		{"Valid_abca_xyzbac_Expects_3", args{"abca", "xyzbac"}, 3},
		{"Valid_a_b_Expects_0", args{"a", "b"}, 0},
		{"Valid_a_aaa_Expects_1", args{"a", "aaa"}, 1},
		{"Invalid_ArgOneLength_Gt+15", args{"abcdefghijklmnop", "aaa"}, 0},
		{"Invalid_ArgTwoLength_Gt+15", args{"a", "abcdefghijklmnop"}, 0},
		{"Invalid_EmptyArgs", args{"", ""}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CommonCharacterCount(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("CommonCharacterCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
