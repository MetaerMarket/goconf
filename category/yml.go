package category

import (
	"os"

	"gopkg.in/yaml.v3"
)

func ReadYaml(file string, obj interface{}) (interface{}, error) {
	b, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(b, &obj)
	if err != nil {
		return nil, err
	}

	return obj, nil
}
