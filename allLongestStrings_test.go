package codesignal

import (
	"reflect"
	"testing"
)

func TestAllLongestStrings(t *testing.T) {
	type args struct {
		inputArray []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"Valid_LongestLength_Three", args{[]string{"aba", "aa", "ad", "vcd", "aba"}}, []string{"aba", "vcd", "aba"}},
		{"Valid_LongestLength_Four", args{[]string{"abc", "eeee", "abcd", "dcd"}}, []string{"eeee", "abcd"}},
		{"Valid_LongestLength_Six", args{[]string{"a", "abc", "cbd", "zzzzzz", "a", "abcdef", "asasa", "aaaaaa"}}, []string{"zzzzzz", "abcdef", "aaaaaa"}},
		{"Invalid_EmptySlice", args{make([]string, 0)}, nil},
		{"Invalid_NilSlice", args{nil}, nil},
		{"Invalid_SliceLength_Gt+10", args{make([]string, 0)}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AllLongestStrings(tt.args.inputArray); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AllLongestStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
