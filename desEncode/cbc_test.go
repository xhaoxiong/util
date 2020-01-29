/**
 * @Author: xiaoxiao
 * @Description:
 * @File:  cbc_test
 * @Version: 1.0.0
 * @Date: 2019/10/9 10:10 下午
 */

package desEncode

import (
	"encoding/base64"
	"testing"
)

func TestEncryptDES(t *testing.T) {
	tt := []byte("test")
	key := []byte("sfe023f_")
	//iv := []byte("abc--iv-")
	v, _ := DesEncrypt(tt, key)

	t.Log(base64.StdEncoding.EncodeToString(v))

}
