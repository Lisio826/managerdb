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

type MysqlServer struct {
	Ip       string `yaml:ip`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type RedisServer struct {
	Addr         string `yaml:"addr"`
	Password     string `yaml:"password"`
	DB           int    `yaml:"db"`
	MaxRetries   int    `yaml:"maxRetries"`
	DialTimeout  int    `yaml:"dialTimeout"`
	ReadTimeout  int    `yaml:"readTimeout"`
	WriteTimeout int    `yaml:"writeTimeout"`
	PoolSize     int    `yaml:"poolSize"`
	PoolTimeout  int    `yaml:"poolTimeout"`
	IdleTimeout  int    `yaml:"idleTimeout"`
}
type Other struct {
	Ignoredbs string `yaml:"ignoredbs"`
	Sercet    string `yaml:"sercet"`
}

type config struct {
	Other       Other       `yaml:"other"`
	LogPath     LogPath     `yaml:"logPath"`
	MysqlServer MysqlServer `yaml:"mysql"`
	RedisServer RedisServer `yaml:"redis"`
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
