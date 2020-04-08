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
	st, ok := elem.(struct{})
	if !ok {
		return "", errors.New("断言失败,请传入struct类型的变量")
	}

	object := reflect.ValueOf(st)
	ref := object.Elem()
	typo := ref.Type()
	var marshalStr strings.Builder
	fmt.Fprint(&marshalStr, `(`)
	for i := 0; i < ref.NumField(); i++ {
		field := ref.Field(i)
		key := typo.Field(i).Name
		val := field.String()
		_, _ = fmt.Fprint(&marshalStr, GetCADString(key, val))
	}
	fmt.Fprint(&marshalStr, `)`)
	return marshalStr.String(), nil
}

func GetCADString(key, val string) string {
	return `("` + key + `".` + `"` + val + `")`
}
