package main

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

type Config struct {
	AssetsPath  string     `yaml:"assets-path" json:"-"`
	ScriptsPath string     `yaml:"scripts-path" json:"-"`
	Port        int        `yaml:"port" json:"-"`
	Networks    []*Network `yaml:"networks" json:"networks"`
}

func (c *Config) Init() error {
	bytes, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		return err
	}

	runtime.GOMAXPROCS(runtime.NumCPU())
	return yaml.Unmarshal(bytes, c)
}
