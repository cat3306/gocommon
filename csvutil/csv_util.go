package csvutil

import (
	"bytes"
	"encoding/csv"
	"fmt"
)

func GenCsvRaw(columns [][]string) ([]byte, error) {

	buf := bytes.NewBuffer(nil)
	w := csv.NewWriter(buf)
	err := w.Write([]string{"\xEF\xBB\xBF"}) //乱码
	if err != nil {
		return nil, err
	}
	err = w.WriteAll(columns)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return buf.Bytes(), nil
}
