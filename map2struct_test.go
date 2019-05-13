/**
*@Package:util
*@Author: haoxiongxiao
*@Date: 2019/5/13
*@Description: create go file in util package
 */
package util

import "testing"

func TestMap2Struct(t *testing.T) {

	type args struct {
		obj interface{}
		m   map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Map2Struct(tt.args.obj, tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("Map2Struct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
