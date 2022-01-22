package config

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	RpcServerConf RpcServerConf `yaml:"RpcServerConf"`
	DataSource    DataSource    `yaml:"DataSource"`
	Etcds         []Etcd        `yaml:"Etcds"`
	Cache         Cache         `yaml:"Cache"`
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
	Name  string   `yaml:"Name"`
	Hosts []string `yaml:"Hosts"`
	Key   string   `yaml:"Key"`
}

type Cache struct {
	Hosts []string `yaml:"Hosts"`
}

func (config *Config) FindEtcdSvc(name string) (etcd Etcd) {
	etcds := config.Etcds
	for _, v := range etcds {
		if v.Name == name {
			etcd = v
			return
		}
	}
	return
}

func Init() *Config {
	var config *Config
	content, err := ioutil.ReadFile("./config/config.yaml")
	if err != nil {
		log.Fatalf("解析config.yaml读取错误: %v", err)
	}
	if yaml.Unmarshal(content, &config) != nil {
		log.Fatalf("解析config.yaml出错: %v", err)
	}
	return config
}
