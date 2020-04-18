/**
 * @Author: xiaoxiao
 * @Description:
 * @File:  marshal
 * @Version: 1.0.0
 * @Date: 4/8/20 6:32 下午
 */
package cad

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func Marshal(elem interface{}) (string, error) {
	var ref reflect.Value
	var typo reflect.Type
	//参数类型
	elemType := reflect.TypeOf(elem).Kind().String()

	//对象
	object := reflect.ValueOf(elem)

	if elemType == "ptr" {
		//获取指针对象
		ref = object.Elem()
		if ref.Type().Kind().String() != "struct" {
			return "", errors.New("请传入结构体或者结构体指针")
		}
		typo = ref.Type()
	} else if elemType == "struct" {
		ref = object
		typo = object.Type()
	} else {
		return "", errors.New("请传入结构体或者结构体指针")
	}

	var marshalStr strings.Builder
	_, _ = fmt.Fprint(&marshalStr, `(`)
	for i := 0; i < ref.NumField(); i++ {
		field := ref.Field(i)
		key := typo.Field(i).Name
		val := field.String()
		_, _ = fmt.Fprint(&marshalStr, GetCADString(key, val))
	}
	_, _ = fmt.Fprint(&marshalStr, `)`)
	return marshalStr.String(), nil
}

func GetCADString(key, val string) string {
	var build strings.Builder
	build.WriteString(`("`)
	build.WriteString(key)
	build.WriteString(`".`)
	build.WriteString(`"`)
	build.WriteString(val)
	build.WriteString(`")`)
	return build.String()
}

