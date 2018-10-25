package codesignal

import "testing"

func TestIsLucky(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Valid_", args{1230}, true},
		{"Valid_", args{239017}, false},
		{"Valid_", args{134008}, true},
		{"Valid_", args{10}, false},
		{"Valid_", args{11}, true},
		{"Valid_", args{1010}, true},
		{"Valid_", args{261534}, false},
		{"Valid_", args{100000}, false},
		{"Valid_", args{999999}, true},
		{"Valid_", args{123321}, true},
		{"Valid_", args{1230}, true},
		{"Valid_", args{1230}, true},
		{"Invalid_OddNumberOfDigits", args{123}, false},
		{"Invalid_Argument_Lt+10", args{9}, false},
		{"Invalid_Argument_Gt+999999", args{1000000}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsLucky(tt.args.n); got != tt.want {
				t.Errorf("IsLucky() = %v, want %v", got, tt.want)
			}
		})
	}
}
