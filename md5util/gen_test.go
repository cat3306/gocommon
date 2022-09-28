package md5util

import "testing"

func TestStrLowerCaseMD5(t *testing.T) {
	t.Logf(StrLowerCaseMD5("123"))
}
func TestStrUpperCaseMD5(t *testing.T) {
	t.Logf(StrUpperCaseMD5("123"))
}
func TestUpperCaseMD5(t *testing.T) {
	t.Logf(UpperCaseMD5([]byte("123")))
}

func TestLowerCaseMD5(t *testing.T) {
	t.Logf(LowerCaseMD5([]byte("123")))
}
