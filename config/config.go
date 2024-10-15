package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

var GlobalConfig Config

type Config struct {
	MQTT       MQTT       `yaml:"MQTT"`
	Middleware Middleware `yaml:"Middleware"`
	Mysql      Mysql      `yaml:"Mysql"`
}

type Mysql struct {
	Url      string `yaml:"Url"`
	Port     int    `yaml:"Port"`
	Password string `yaml:"Password"`
}

type MQTT struct {
	Server  Server  `yaml:"Server"`
	Client  Client  `yaml:"Client"`
	Message Message `yaml:"Message"`
	Topic   Topic   `yaml:"Topic"`
}

type Server struct {
	In_Server  string `yaml:"In_Server"`
	Out_Server string `yaml:"Out_Server"`
}

type Client struct {
	Username string `yaml:"Username"`
	Password string `yaml:"Password"`
	ClientId string `yaml:"ClientId"`
}

type Message struct {
	Retained bool `yaml:"Retained"`
	QOS      int  `yaml:"QOS"`
}

type Topic struct {
	In_Topic  string `yaml:"In_Topic"`
	Out_Topic string `yaml:"Out_Topic"`
}

type Middleware struct {
	ReadOnly bool `yaml:"ReadOnly"`
}

func Init() {
	dataBytes, err := os.ReadFile("./config.yaml")
	if err != nil {
		fmt.Println("read yaml fail：", err)
	}
	config := Config{}
	err = yaml.Unmarshal(dataBytes, &config)
	log.Println("read yaml successfully")
	log.Printf("yaml:%+v\n", config)
	GlobalConfig = config
}
