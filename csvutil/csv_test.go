package csvutil

import (
	"fmt"
	"github.com/cat3306/gocommon/strutil"
	"io/ioutil"
	"testing"
)

func TestGenCsvRaw(t *testing.T) {
	c := make([][]string, 0)
	title := []string{"id", "name", "age", "height", "gender"}
	m := struct {
		Id     int
		Name   string
		Age    int
		Height float64
		Gender int
	}{}
	m.Id = 1
	m.Name = "joker"
	m.Age = 18
	m.Gender = 1
	m.Height = 1.88
	c = append(c, title)
	s := strutil.StructToStringArray(m)
	fmt.Println(s)
	c = append(c, strutil.StructToStringArray(m))
	raw, err := GenCsvRaw(c)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ioutil.WriteFile("person.csv", raw, 0777)
	fmt.Println(err)
}
