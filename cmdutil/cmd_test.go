package utils

import "testing"

func TestCmd(t *testing.T) {
	t.Log(Cmd("cat cmd_test.go"))
}
