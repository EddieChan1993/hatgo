package wx

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
)

type ResEncryptData struct {
	UnionId   string    `json:"unionid"` //多平台用户唯一标识
	Watermark Watermark `json:"watermark"`
}

type Watermark struct {
	Appid string `json:"appid"`
}

//ase解码微信公开数据
func DecryptWXOpenData(sessionKey, encryptData, iv string) (*ResEncryptData, error) {
	decodeBytes, err := base64.StdEncoding.DecodeString(encryptData)
	if err != nil {
		return nil, err
	}
	sessionKeyBytes, err := base64.StdEncoding.DecodeString(sessionKey)
	if err != nil {
		return nil, err
	}
	ivBytes, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return nil, err
	}
	dataBytes, err := AesDecrypt(decodeBytes, sessionKeyBytes, ivBytes)
	m := new(ResEncryptData)
	err = json.Unmarshal(dataBytes, &m)
	if err != nil {
		return nil, err
	}
	//temp := m.Watermark.Appid
	appid := m.Watermark.Appid
	if appid != appidXCX {
		return nil, fmt.Errorf("invalid appid, get !%s!", appid)
	}
	if err != nil {
		return nil, err
	}
	return m, nil
}

func AesDecrypt(crypted, key, iv []byte) ([]byte, error) {
	//即密钥16，24，32长度对应AES-128, AES-192, AES-256。
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS7UnPadding(origData)
	return origData, nil
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
