/**
 * @Author: xiaoxiao
 * @Description:
 * @File:  convert2byte
 * @Version: 1.0.0
 * @Date: 2019/9/29 3:48 下午
 */

package enconvert

import "github.com/axgle/mahonia"


//src 需要字符串 srcCode 源编码 targetCode 需转编码
func Convert2Byte(src string, srcCode string, targetCode string) []byte {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(targetCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	return cdata
}
