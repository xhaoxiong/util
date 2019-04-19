/**
*@Author: haoxiongxiao
*@Date: 2019/4/17
*@Description: CREATE GO FILE util
 */
package util

import (
	"testing"
)

func TestVerifiyEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		struct {
			name string
			args args
			want bool
		}{name: "1", args: args{email: "asdasdas"}, want: false},
		{name: "2", args: args{email: "269812428@qq@.com"}, want: false},
		{name: "2", args: args{email: "269812428@qq.com"}, want: true},
		{name: "2", args: args{email: "269812428@@qq.com"}, want: false},
		{name: "2", args: args{email: "2@69812428@@qq.com"}, want: false},
		{name: "2", args: args{email: "2@69812428qq.com"}, want: true},
		{name: "2", args: args{email: "2@sdsss@.qq.com"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VerifiyEmail(tt.args.email); got != tt.want {
				t.Errorf("VerifiyEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVerification(t *testing.T) {
	type args struct {
		idCardNo string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		struct {
			name string
			args args
			want bool
		}{name: "1", args: args{idCardNo: "43252219971030489X"}, want: true},
		{name: "2", args: args{idCardNo: "43252219971020444"}, want: false},
		{name: "2", args: args{idCardNo: "43252219971020444x"}, want: false},
		{name: "2", args: args{idCardNo: "43252219971040489s"}, want: false},
		{name: "2", args: args{idCardNo: "4325221997103048922"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Verification(tt.args.idCardNo); got != tt.want {
				t.Errorf("Verification() = %v, want %v", got, tt.want)
			}
		})
	}
}
