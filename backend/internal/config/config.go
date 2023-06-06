package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

type Network struct {
	ID   string `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`
	URL  string `json:"url" yaml:"url"`
}

type Server struct {
	URL  string `yaml:"url"`
	Port uint16 `yaml:"port"`
}

type Config struct {
	Server   *Server   `yaml:"server" json:"-"`
	Networks []Network `yaml:"networks" json:"networks"`
}

func Import(path string) (Config, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	config := Config{}
	err = yaml.Unmarshal(bytes, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
