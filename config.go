package main

import (
	"github.com/kovetskiy/ko"
	"gopkg.in/yaml.v2"
)

type Config map[string]string

func getConfig(path string) (Config, error) {
	config := Config{}
	err := ko.Load(path, &config, yaml.Unmarshal)
	if err != nil {
		return nil, err
	}

	return config, nil
}
