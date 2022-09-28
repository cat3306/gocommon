package cryptoutil

import (
	"fmt"
	"testing"
)

func TestRsaEncrypt(t *testing.T) {
	prvKey, err := RawRSAKey("private.pem")
	if err != nil {
		fmt.Println(err)
		return
	}
	pubKey, err := RawRSAKey("public.pem")
	if err != nil {
		fmt.Println(err)
		return
	}
	pass := RsaEncrypt([]byte("love"), pubKey)
	fmt.Println(string(pass))
	text := RsaDecrypt(pass, prvKey)
	fmt.Println(string(text))
}
