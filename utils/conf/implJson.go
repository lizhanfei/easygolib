package conf

import (
	jsoniter "github.com/json-iterator/go"
	"os"
)

type ImplJson struct {
}

func (this *ImplJson) Load(filePath string, t interface{}) error {
	if yamlFile, err := os.ReadFile(filePath); err != nil {
		return err
	} else if err = jsoniter.Unmarshal(yamlFile, t); err != nil {
		return err
	}
	return nil
}
