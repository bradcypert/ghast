package config

import (
	"io/ioutil"
	"log"

	"github.com/jeremywohl/flatten"
	"gopkg.in/yaml.v2"
)

func Parse(configPath string) (map[interface{}]interface{}, error) {
	m := make(map[interface{}]interface{})
	data, err := ioutil.ReadFile(configPath)

	if err != nil {
		return &m, err
	}

	err = yaml.Unmarshal(data, &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return &m, err
}

func ParsedConfigToContainerKeys(parsed map[interface{}]interface{}) (map[string]interface{}, err) {
	return flatten.FlattenString(parsed, "", flatten.DotStyle)
}
