/**
*@Author: haoxiongxiao
*@Date: 2019/4/19
*@Description: CREATE GO FILE reg
 */
package reg

import (
	"regexp"
	"strings"
)

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

//处理json key没有引号的函数
func SerializeJSON(str string) string {
	return regexp.MustCompile(`[\w]+[:]`).ReplaceAllStringFunc(str, func(s string) string {
		return `"` + regexp.MustCompile(`[\w]+`).FindAllString(s, -1)[0] + `":`
	})
}

func SerializeJSON2(str string) string {
	str = strings.Replace(str, `{`, `{"`, -1)
	str = strings.Replace(str, `:`, `":`, -1)
	str = strings.Replace(str, `,`, `,"`, -1)
	str = strings.Replace(str, `,"{`, `,{`, -1)
	return str
}
