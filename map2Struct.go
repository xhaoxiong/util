/**
*@Package:util
*@Author: haoxiongxiao
*@Date: 2019/5/13
*@Description: create go file in util package
 */
package util

import "encoding/json"

func Map2Struct(obj interface{}, m map[string]interface{}) error {
	data, _ := json.Marshal(m)

	if err := json.Unmarshal(data, &obj); err != nil {
		return err
	}
	return nil
}
