package ziputil

import (
	"archive/zip"
	"bytes"
)

type RawFile struct {
	Name string //压缩的文件名
	Data []byte //字节信息
}

func BytesZip(files []RawFile) ([]byte, error) {
	buff := new(bytes.Buffer)
	w := zip.NewWriter(buff)
	defer w.Close()
	for _, v := range files {
		ww, err := w.Create(v.Name)
		if err != nil {
			return nil, err
		}
		_, err = ww.Write(v.Data)
		if err != nil {
			return nil, err
		}
	}
	err := w.Close()
	if err != nil {
		return nil, err
	}
	return buff.Bytes(), nil
}
