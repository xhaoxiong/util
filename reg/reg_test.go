/**
*@Author: haoxiongxiao
*@Date: 2019/4/19
*@Description: CREATE GO FILE reg
 */
package reg

import "testing"

func TestCheckUserName(t *testing.T) {
	type args struct {
		str string
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
		}{name: "1", args: args{"["}, want: false},
		{name: "2", args: args{"[]"}, want: false},
		{name: "2", args: args{"!"}, want: false},
		{name: "2", args: args{"*"}, want: false},
		{name: "2", args: args{"/"}, want: false},
		{name: "2", args: args{"+"}, want: false},
		{name: "2", args: args{"-"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckUserName(tt.args.str); got != tt.want {
				t.Errorf("CheckUserName() = %v, want %v", got, tt.want)
			}
		})
	}
}
