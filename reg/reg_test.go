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

func TestSerializeJSON(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		struct {
			name string
			args args
			want string
		}{name: "1", args: args{`{total: 100,rows: [{goodId: "4893ed8518ff41268baf5a2f7f3e5398",reportedStaffid: "",reportedName: "",reportedPst: "",reportedOrg: "",reportedTime: "",goodTitle: "",goodContent: "",reportStaffid: "",reportName: "",reportPst: "",reportOrg: "",publishRange: "",praiseCnt: "",commentCnt: "0",}]}`},
			want: `{"total": 100,"rows": [{"goodId": "4893ed8518ff41268baf5a2f7f3e5398","reportedStaffid": "","reportedName": "","reportedPst": "","reportedOrg": "","reportedTime": "","goodTitle": "","goodContent": "","reportStaffid": "","reportName": "","reportPst": "","reportOrg": "","publishRange": "","praiseCnt": "","commentCnt": "0",}]}`}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SerializeJSON(tt.args.str); got != tt.want {
				t.Errorf("SerializeJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
