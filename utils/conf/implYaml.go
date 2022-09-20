package conf

import (
	"gopkg.in/yaml.v3"
	"os"
)

type ImplYaml struct {
}

func (this *ImplYaml) Load(filePath string, t interface{}) error {
	if yamlFile, err := os.ReadFile(filePath); err != nil {
		return err
	} else if err = yaml.Unmarshal(yamlFile, t); err != nil {
		return err
	}
	return nil
}
