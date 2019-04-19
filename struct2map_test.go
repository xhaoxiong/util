/**
*@Author: haoxiongxiao
*@Date: 2019/4/18
*@Description: CREATE GO FILE util
 */
package util

import (
	"reflect"
	"testing"
)

func TestStruct2Map(t *testing.T) {
	type args struct {
		obj interface{}
	}

	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		// TODO: Add test cases.

		struct {
			name string
			args args
			want map[string]interface{}
		}{name: "1", args: args{struct {
			Name int
			Val  int
		}{Name: 1, Val: 2}}, want: map[string]interface{}{"Name": 1, "Val": 2}},
		{name: "2", args: args{struct {
			Name string
			Val  string
		}{Name: "3", Val: "4"}}, want: map[string]interface{}{"Name": "3", "Val": "4"}},
		{name: "3", args: args{struct {
			Name string
			Val  string
		}{Name: "5", Val: "6"}}, want: map[string]interface{}{"Name": "5", "Val": "6"}},
		{name: "4", args: args{struct {
			Name string
			Val  string
		}{Name: "7", Val: "8"}}, want: map[string]interface{}{"Name": "7", "Val": "8"}},
		{name: "5", args: args{struct {
			Name string
			Val  string
		}{Name: "9", Val: "10"}}, want: map[string]interface{}{"Name": "9", "Val": "10"}},
		{name: "6", args: args{struct {
			Name string
			Val  string
		}{Name: "3", Val: "4"}}, want: map[string]interface{}{"Name": "3", "Val": "4"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Struct2Map(tt.args.obj); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Struct2Map() = %v, want %v", got, tt.want)
			}
		})
	}
}
