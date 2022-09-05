package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"runtime"
)

type Config struct {
	AssetsPath  string `yaml:"assets-path"`
	ScriptsPath string `yaml:"scripts-path"`
	Port        int    `yaml:"port"`
}

func (c *Config) Init() error {
	bytes, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		return err
	}

	runtime.GOMAXPROCS(runtime.NumCPU())
	return yaml.Unmarshal(bytes, c)
}
