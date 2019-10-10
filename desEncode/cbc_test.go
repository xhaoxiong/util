/**
 * @Author: xiaoxiao
 * @Description:
 * @File:  cbc_test
 * @Version: 1.0.0
 * @Date: 2019/10/9 10:10 下午
 */

package desEncode

import (
	"fmt"
	"testing"
)

func TestEncryptDES(t *testing.T) {
	x:=[]byte("长长的头发，黑黑的眼睛。")
	key:=[]byte("12345678")
	x1:=EncryptDES(x,key)
	x2:=DecryptDES(x1,key)
	fmt.Print(string(x2))
}