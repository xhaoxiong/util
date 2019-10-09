/**
 * @Author: xiaoxiao
 * @Description:
 * @File:  byte2int
 * @Version: 1.0.0
 * @Date: 2019/10/9 9:20 下午
 */

package socktByteUtil

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"fmt"
)

//整形转换成字节
func IntToBytes(n int) []byte {
	x := int32(n)

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

//字节转换成整形
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return int(x)
}

func BytesTo16(bt []byte) string {
	sign := md5.Sum(bt)
	return fmt.Sprintf("%x", sign)

}
