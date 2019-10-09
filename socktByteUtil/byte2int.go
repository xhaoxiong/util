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
	"encoding/binary"
)

func Byte2int(bt []byte) (v int) {
	buf := bytes.NewReader(bt)
	binary.Read(buf, binary.BigEndian, &v)
	return
}
