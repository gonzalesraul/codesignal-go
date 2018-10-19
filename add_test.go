package codesignal_test

import (
	"testing"

	. "github.com/gonzalesraul/codesignal-go"
)

func TestAdd(t *testing.T) {
	type args struct {
		param1 int
		param2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Valid_Expects_3", args{1, 2}, 3},
		{"Valid_Expects_4", args{2, 2}, 4},
		{"Valid_Expects_1001", args{1000, 1}, 1001},
		{"Valid_Expects_-999", args{-1000, 1}, -999},
		{"Valid_Expects_999", args{-1, 1000}, 999},
		{"Invalid_Param1_Gt+1000", args{1001, 1}, 0},
		{"Invalid_Param1_Lt-1000", args{-1001, 1}, 0},
		{"Invalid_Param2_Gt+1000", args{1, 1001}, 0},
		{"Invalid_Param2_Lt-1000", args{1, -1001}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.param1, tt.args.param2); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
