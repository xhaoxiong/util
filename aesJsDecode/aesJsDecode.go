/**
 * @Author xiaoxiao
 * @Description CREATE FILE aesJsDecode
 * @Date 2020/10/19 9:12 上午
 **/
package aesJsDecode

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

// Decrypt Golang解密
// ciphertext  important important 上面js的生成的密文进行了 hex.encoding 在这之前必须要进行 hex.Decoding
// 上面js代码最后返回的是16进制
// 所以收到的数据hexText还需要用hex.DecodeString(hexText)转一下，这里略了
func Decrypt(ciphertext, key []byte) ([]byte, error) {
	pkey := PaddingLeft(key, '0', 16)//和js的key补码方法一致

	block, err := aes.NewCipher(pkey) //选择加密算法
	if err != nil {
		return nil, fmt.Errorf("key 长度必须 16/24/32长度: %s", err)
	}
	blockModel := cipher.NewCBCDecrypter(block, pkey) //和前端代码对应:   mode: CryptoJS.mode.CBC,// CBC算法
	plantText := make([]byte, len(ciphertext))
	blockModel.CryptBlocks(plantText, ciphertext)
	plantText = PKCS7UnPadding(plantText) //和前端代码对应:  padding: CryptoJS.pad.Pkcs7
	return plantText, nil
}

func PKCS7UnPadding(plantText []byte) []byte {
	length := len(plantText)
	unpadding := int(plantText[length-1])
	return plantText[:(length - unpadding)]
}
//这个方案必须和js的方法是一样的
func PaddingLeft(ori []byte, pad byte, length int) []byte {
	if len(ori) >= length {
		return ori[:length]
	}
	pads := bytes.Repeat([]byte{pad}, length-len(ori))
	return append(pads, ori...)
}
