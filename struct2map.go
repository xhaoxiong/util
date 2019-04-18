/**
*@Author: haoxiongxiao
*@Date: 2019/4/18
*@Description: CREATE GO FILE util
 */
package util

import "reflect"

func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}
