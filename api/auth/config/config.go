package config

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	RpcServerConf RpcServerConf `yaml:"RpcServerConf"`
	DataSource    DataSource    `yaml:"DataSource"`
	Etcd          Etcd          `yaml:"Etcd"`
	Cache         Etcd          `yaml:"Cache"`
}

type RpcServerConf struct {
	ListenOn       string `yaml:"ListenOn"`
	ServiceAddress string `yaml:"ServiceAddress"`
}

type DataSource struct {
	Address  string `yaml:"Address`
	Port     string `yaml:"Port"`
	User     string `yaml:"User"`
	Password string `yaml:"Password"`
	Database string `yaml:"Database"`
}

type Etcd struct {
	Hosts []string `yaml:"Hosts"`
	Key   string   `yaml:"Key"`
}

type Cache struct {
	Hosts []string `yaml:"Hosts"`
}

func (config Config) Init() Config {
	content, err := ioutil.ReadFile("./auth/config/config.yaml")
	if err != nil {
		log.Fatalf("解析config.yaml读取错误: %v", err)
	}
	if yaml.Unmarshal(content, &config) != nil {
		log.Fatalf("解析config.yaml出错: %v", err)
	}
	return config
}
