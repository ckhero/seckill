package config

import (
	"fmt"
	"github.com/jinzhu/configor"
	"os"
	"path/filepath"
	"sync"
)

type Config struct {
	Redis redis
}

type redis struct {
	Host string `yaml:"host"`
	Port string	`yaml:"port"`
	Password string `yaml:"password"`
}

var (
	appName string
	runEnv string
	AppConfig *Config
	once sync.Once
)

func GetApp() string {
	return appName;
}

func GetEnv() string {
	return runEnv;
}

func InitConfig(app, env string)  {
	once.Do(func() {
		appName = app
		runEnv = env
		pwd, _ := os.Getwd();
		path := filepath.Join(pwd, "config", env) + ".yml";
		AppConfig = loadConfig(path)
	});
}

func loadConfig(path string) *Config  {
	config := new (Config)
	if err :=configor.Load(config, path); err != nil {
		panic(fmt.Errorf("load config fail:%s\n", err));
	}
	return config;
}