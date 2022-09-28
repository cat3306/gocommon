package cryptoutil

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

func pKCS7Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

func pKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	padding := int(origData[length-1])
	return origData[:(length - padding)]
}

//AES加密,CBC
func AesEncrypt(origStr string, key []byte) (string, error) {
	origData := []byte(origStr)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	origData = pKCS7Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	dst := make([]byte, len(origData))
	blockMode.CryptBlocks(dst, origData)
	return base64.StdEncoding.EncodeToString(dst), nil
}

//AES解密
func AesDecrypt(srcStr string, key []byte) (string, error) {
	if srcStr == "" {
		return "", errors.New("srcStr nil")
	}
	src, err := base64.StdEncoding.DecodeString(srcStr)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(src))
	blockMode.CryptBlocks(origData, src)
	origData = pKCS7UnPadding(origData)
	return string(origData), nil
}
