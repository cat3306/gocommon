package ziputil

import (
	"io/ioutil"
	"testing"
)

func TestXxx(t *testing.T) {
	raw, err := ioutil.ReadFile("./1.log")
	if err != nil {
		t.Fatal(err)
	}
	raw1, err := ioutil.ReadFile("./2.log")
	if err != nil {
		t.Fatal(err)
	}
	bs, err := BytesZip([]RawFile{
		{Name: "1", Data: raw},
		{Data: raw1, Name: "2"},
	})
	if err != nil {
		t.Fatal(err)
	}
	ioutil.WriteFile("./2.zip", bs, 0777)
}
