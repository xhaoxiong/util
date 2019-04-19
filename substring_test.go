/**
*@Author: haoxiongxiao
*@Date: 2019/4/19
*@Description: CREATE GO FILE util
 */
package util

import "testing"

func TestSubstr(t *testing.T) {
	type args struct {
		str    string
		start  int
		length int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.

		struct {
			name string
			args args
			want string
		}{name: "1", args: args{str: "123456789", start: 0, length: 2}, want: "12"},
		{name: "2", args: args{str: "123456789", start: -1, length: 2}, want: "89"},
		{name: "3", args: args{str: "哈哈哈哈", start: -1, length: 2}, want: "哈哈"},
		{name: "4", args: args{str: "哈哈哈哈", start: 0, length: 2}, want: "哈哈"},
		{name: "5", args: args{str: "哈哈哈哈，中国人", start: 4, length: 2}, want: "，中"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Substr(tt.args.str, tt.args.start, tt.args.length); got != tt.want {
				t.Errorf("Substr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubstr2(t *testing.T) {
	type args struct {
		str   string
		start int
		end   int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		struct {
			name string
			args args
			want string
		}{name: "1", args: args{str: "123456789", start: 0, end: 1}, want: "1"},
		{name: "2", args: args{str: "123456789", start: 0, end: 9}, want: "123456789"},
		{name: "3", args: args{str: "123456789", start: -1, end: 9}, want: "panic"},
		{name: "4", args: args{str: "123456789", start: 0, end: 10}, want: "panic"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Substr2(tt.args.str, tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("Substr2() = %v, want %v", got, tt.want)
			}
		})
	}
}
