package confutil

import (
	"encoding/json"
	"gopkg.in/yaml.v3"
	"path/filepath"
	"strings"
)

type Encoder interface {
	Unmarshal([]byte, interface{}) error
	Marshal(interface{}) ([]byte, error)
}

type yamlEncoding struct{}

func (y yamlEncoding) Unmarshal(b []byte, v interface{}) error {
	return yaml.Unmarshal(b, v)
}

func (y yamlEncoding) Marshal(v interface{}) ([]byte, error) {
	return yaml.Marshal(v)
}

// JSON encoding implements ConfigEncoding
type jsonEncoding struct{}

func (j jsonEncoding) Unmarshal(b []byte, v interface{}) error {
	return json.Unmarshal(b, v)
}

func (j jsonEncoding) Marshal(v interface{}) ([]byte, error) {
	return json.MarshalIndent(v, "", " ")
}
func file2Encoder(fileExtension string) Encoder {
	// Lower the file extension
	ext := strings.ToLower(fileExtension)
	ext = strings.TrimPrefix(ext, ".")
	// Return the appropriate encoder/decoder according
	// to the extension
	switch ext {
	case "yml", "yaml":
		// YAML
		return yamlEncoding{}
	case "json":
		return jsonEncoding{}
	default:
		// JSON
		return jsonEncoding{}
	}
}

func marshal(fileName string, v interface{}) ([]byte, error) {
	ext := filepath.Ext(fileName)
	return file2Encoder(ext).Marshal(v)
}

func unMarshal(fileName string, raw []byte, v interface{}) error {
	ext := filepath.Ext(fileName)
	return file2Encoder(ext).Unmarshal(raw, v)
}
