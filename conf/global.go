package conf

import (
	"github.com/labstack/gommon/log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var Global config

type LogPath struct {
	Logdb    string `yaml:"logdb"`
	LogLocal string `yaml:"logLocal"`
}
type config struct {
	Ignoredbs string  `yaml:"ignoredbs"`
	LogPath   LogPath `yaml:"logPath"`
}

func (c *config) GetConf() config {
	yamlFile, err := ioutil.ReadFile("conf/global.yaml")
	if err != nil {
		log.Error("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Error("Unmarshal: %v", err)
	}
	return *c
}

