package md5util

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func LowerCaseMD5(data []byte) string {
	md5Hash := md5.New()
	md5Hash.Write(data)
	md5Data := md5Hash.Sum(nil)
	return hex.EncodeToString(md5Data)
}
func UpperCaseMD5(data []byte) string {
	return strings.ToUpper(LowerCaseMD5(data))
}
func StrLowerCaseMD5(str string) string {
	md5Hash := md5.New()
	md5Hash.Write([]byte(str))
	md5Data := md5Hash.Sum(nil)
	return hex.EncodeToString(md5Data)
}
func StrUpperCaseMD5(str string) string {
	return strings.ToUpper(StrLowerCaseMD5(str))
}
