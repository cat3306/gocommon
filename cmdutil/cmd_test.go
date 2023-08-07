package utils

import (
	"fmt"
	"testing"
)

func TestCmd(t *testing.T) {
	s, _ := Cmd("cat cmd_test.go")
	fmt.Println(s)
}
