/**
*@Author: haoxiongxiao
*@Date: 2019/3/8
*@Description: CREATE GO FILE util
 */
package util

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
)

func Str2Md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	rs := hex.EncodeToString(h.Sum(nil))
	return rs
}

func Str2SHA256(data string) string {
	t := sha256.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}
