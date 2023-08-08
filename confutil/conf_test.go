package confutil

import (
	"bytes"
	"encoding/json"
	etcd "go.etcd.io/etcd/client/v3"
	"testing"
)

type student struct {
	Age   int    `json:"age"`
	Name  string `json:"name"`
	Class int    `json:"class"`
	Sex   int    `json:"sex"`
}

func TestJsonSaveLoad(t *testing.T) {
	c := Config{}
	s := student{
		Age:   19,
		Name:  "haha",
		Class: 18,
		Sex:   1,
	}
	err := c.Save("./1.json", s)
	if err != nil {
		t.Fatal(err)
	}
	ss := student{}
	err = c.Load("./1.json", &ss)
	b1, err := json.Marshal(s)
	if err != nil {
		t.Fatal(err)
	}
	b2, err := json.Marshal(ss)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(bytes.Equal(b1, b2))
}

func TestYamlSaveLoad(t *testing.T) {
	c := Config{}
	s := student{
		Age:   19,
		Name:  "haha",
		Class: 18,
		Sex:   1,
	}
	err := c.Save("./1.yml", s)
	if err != nil {
		t.Fatal(err)
	}
	ss := student{}
	err = c.Load("./1.yml", &ss)
	b1, err := json.Marshal(s)
	if err != nil {
		t.Fatal(err)
	}
	b2, err := json.Marshal(ss)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(bytes.Equal(b1, b2))
}
func TestConfigEtcdSave(t *testing.T) {
	c := Config{}
	clt, err := etcd.New(etcd.Config{
		Endpoints: []string{"http://localhost:2379"},
	})
	if err != nil {
		t.Fatal(err)
	}
	s := student{
		Age:   1,
		Name:  "dd",
		Class: 89,
		Sex:   100,
	}
	err = c.EtcdMode(clt).Save("/haha/1.json", s)
	if err != nil {
		t.Fatal(err)
	}
	ss := student{}
	err = c.Load("/haha/1.json", &ss)
	if err != nil {
		t.Fatal(err)
	}
	b1, err := json.Marshal(s)
	if err != nil {
		t.Fatal(err)
	}
	b2, err := json.Marshal(ss)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(bytes.Equal(b1, b2))
}
