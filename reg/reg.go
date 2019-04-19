/**
*@Author: haoxiongxiao
*@Date: 2019/4/19
*@Description: CREATE GO FILE reg
 */
package reg

import "regexp"

const (
	regUserName = "[`~!@#$^&*()=|{}':;',\\[\\].<>/?~！@#￥……&*（）——|{}【】‘；：”“'。，、？%+_]"
)

func CheckUserName(str string) bool {
	uReg := regexp.MustCompile(regUserName)

	if s := uReg.FindString(str); s != "" {
		return false
	}
	return true

}
