package config

import (
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server  Server   `yaml:"server"`
	DB      DataBase `yaml:"db"`
	Adapter Adapter  `yaml:"adapter"`
}

type Server struct {
	Name         string        `yaml:"name"`
	Port         string        `yaml:"port"`
	Host         string        `yaml:"host"`
	ReadTimeOut  time.Duration `yaml:"read_time_out"`
	WriteTimeOut time.Duration `yaml:"write_time_out"`
}

type DataBase struct {
	UserName string `yaml:"user_name"`
	Password int    `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DBName   string `yaml:"db_name"`
}

type Adapter struct {
	Url     string        `yaml:"url"`
	Token   string        `yaml:"token"`
	Key     string        `yaml:"key"`
	TimeOut time.Duration `yaml:"timeout"`
}

var path = "./internal/config/config.yaml"

func MustRead() *Config {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("can't open config file %v", err)
	}

	var config Config
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		log.Fatalf("can't unmarshal config %v", err)
	}
	return &config
}
