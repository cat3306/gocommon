package table

import (
	"fmt"
	"testing"
)

type House struct {
	Name  string
	Sigil string
	Motto string
}

func Test1(t *testing.T) {
	s := []House{
		{"", "direwolf", "asd is 阿斯顿"},
		{"", "实时", "Fire and 安师大"},
		{"", "asd", "Hear Me Roar"},
	}
	//s:=House{
	//	Name:  "中国ker",
	//	Sigil: "s是s",
	//	Motto: "sa、就d",
	//}
	tt := Table(s)

	fmt.Println(tt)
}
