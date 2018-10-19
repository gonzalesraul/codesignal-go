package codesignal_test

import (
	"testing"

	. "github.com/gonzalesraul/codesignal-go"
)

func TestCenturyFromYear(t *testing.T) {
	type args struct {
		year int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Valid_Expects_21", args{2005}, 21},
		{"Valid_Expects_20", args{1901}, 20},
		{"Valid_Expects_19", args{1900}, 19},
		{"Valid_Expects_2", args{190}, 2},
		{"Valid_Expects_1", args{100}, 1},
		{"Invalid_Year_Gt+2005", args{2006}, 0},
		{"Invalid_Year_Lt-1", args{-1}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CenturyFromYear(tt.args.year); got != tt.want {
				t.Errorf("CenturyFromYear() = %v, want %v", got, tt.want)
			}
		})
	}
}
