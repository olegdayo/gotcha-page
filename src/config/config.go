package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"runtime"
)

type Config struct {
	Paths  *Paths `yaml:"paths"`
	Port   string
	CPUNum int
}

type Paths struct {
	Templates    string `yaml:"templates"`
	Assets       string `yaml:"assets"`
	Scripts      string `yaml:"scripts"`
	MainTemplate string `yaml:"main-template"`
}

func Init() (*Config, error) {
	bytes, err := ioutil.ReadFile("src/config/config.yaml")
	if err != nil {
		return nil, err
	}

	config := new(Config)
	err = yaml.Unmarshal(bytes, config)
	if err != nil {
		return nil, err
	}

	config.CPUNum = runtime.NumCPU()
	config.Port = os.Getenv("PORT")
	if config.Port == "" {
		config.Port = "8080"
	}
	return config, nil
}
