package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Telegram struct {
		BotAPI string `yaml:"bot-api"`
	} `yaml:"telegram"`
	Database struct {
		Url string `yaml:"conn-url"`
	} `yaml:"database"`
	Server struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		RestPort string `yaml:"restport"`
	} `yaml:"server"`
	Parser struct {
		DelaySec int64 `yaml:"delay-sec"`
	} `yaml:"parser"`
}

func New(filename string) (*Config, error) {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read config, err: <%v>", err)
	}

	var config Config

	if err := yaml.Unmarshal(raw, &config); err != nil {
		return nil, fmt.Errorf("failed to unmarshall config, err: <%v>", err)
	}

	return &config, nil
}
