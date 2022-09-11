package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"runtime"
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
	Server   *Server    `yaml:"server" json:"-"`
	Networks []*Network `yaml:"networks" json:"networks"`
}

func (c *Config) Init() error {
	bytes, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		return err
	}

	runtime.GOMAXPROCS(runtime.NumCPU())
	return yaml.Unmarshal(bytes, c)
}
