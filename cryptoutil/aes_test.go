package cryptoutil

import (
	"fmt"
	"testing"
)

func TestAes(t *testing.T) {
	str := "海上月是天上月,眼前人是心上人"
	key := []byte("#HvL%$o0oNNoOZnk#o2qbqCeQB1iXeIR")
	d, err := AesEncrypt(str, key)
	fmt.Println(d, err)
	fmt.Println(AesDecrypt(d, key))
}
func TestAesDecrypt(t *testing.T) {
	fmt.Println(AesDecrypt("4W0hT5SCfe1QsCYNCUEh9g==", []byte("#HvL%$o0oNNoOZnk#o2qbqCeQB1iXeIR")))
}
func TestB(t *testing.T) {
}
