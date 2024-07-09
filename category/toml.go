package category

import (
	"os"

	"github.com/BurntSushi/toml"
)

func ReadToml(file string, obj interface{}) (interface{}, error) {
	b, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	err = toml.Unmarshal(b, &obj)
	if err != nil {
		return nil, err
	}

	return obj, nil
}
