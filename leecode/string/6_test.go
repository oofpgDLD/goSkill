package string

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "t1",args: args{str: "42"},want: 42},
		{name: "t2",args: args{str: "-42"},want: -42},
		{name: "t3",args: args{str: "4193 with words"},want: 4193},
		{name: "t4",args: args{str: "words and 987"},want: 0},
		{name: "t5",args: args{str: "-91283472332"},want: -2147483648},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.str); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}