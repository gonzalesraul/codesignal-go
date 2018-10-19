package codesignal_test

import (
	"testing"

	. "github.com/gonzalesraul/codesignal-go"
)

func TestCheckPalindrome(t *testing.T) {
	type args struct {
		inputString string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Valid_PairLength", args{"akka"}, true},
		{"Valid_OddLength", args{"aka"}, true},
		{"Valid_EspecialChar", args{".1!a!1."}, true},
		{"Valid_SingleChar", args{"a"}, true},
		{"Valid_HugeString", args{"ksswieohzzonjqosouzbnthoqfktlhokcblmleilmrbxvldvodyhuzknirkdlexrprgqbznnwypsjowrjyteosbhqflkvyfshhgzchhchyeeasbopifwhsusjvxsvlxtchrxzywptdljywcqykstepgzufvcxphtjsnxeveuqybmifdbpnwwruqgyzbltubkjzvxhpsuusphxvzjkbutlbzygqurwwnpbdfimbyquevexnsjthpxcvfuzgpetskyqcwyjldtpwyzxrhctxlvsxvjsushwfipobsaeeyhchhczghhsfyvklfqhbsoetyjrwojspywnnzbqgrprxeldkrinkzuhydovdlvxbrmlielmlbckohltkfqohtnbzuosoqjnozzhoeiwssk"}, true},
		{"Invalid_NotAPalindrome", args{"camigol"}, false},
		{"Invalid_EmptyString", args{""}, false},
		{"Invalid_UpperCase", args{"aKa"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckPalindrome(tt.args.inputString); got != tt.want {
				t.Errorf("CheckPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
