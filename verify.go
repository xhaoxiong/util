/**
*@Author: haoxiongxiao
*@Date: 2019/4/17
*@Description: CREATE GO FILE util
 */
package util

// 身份证校验码的计算方法：
//	1、将前面的身份证号码17位数分别乘以不同的系数，系数见：coefficient
// 	2、将这17位数字和系数相乘的结果相加，用加出来和除以11，得到余数Remainder
//	3、余数Remainder作为位置值，在数组code中找到对应的值，就是身份证号码的第18位数值

import (
	"regexp"
	"strings"
)

var (
	coefficient []int32 = []int32{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	code        []byte  = []byte{'1', '0', 'X', '9', '8', '7', '6', '5', '4', '3', '2'}
)

// 校验一个身份证是否是合法的身份证
func Verification(idCardNo string) bool {
	if len(idCardNo) != 18 {
		return false
	}

	idByte := []byte(strings.ToUpper(idCardNo))

	sum := int32(0)
	for i := 0; i < 17; i++ {
		sum += int32(byte(idByte[i])-byte('0')) * coefficient[i]
	}
	return code[sum%11] == idByte[17]
}

func VerifiyEmail(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}
