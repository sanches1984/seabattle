package config

import (
	"fmt"
	"github.com/go-yaml/yaml"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
)

const configName = "app.yaml"

var config appConfig

type ListenerConfig struct {
	Host string `yaml:"host"`
	Port int32  `yaml:"port"`
}

type RedisConfig struct {
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	Db        int    `yaml:"db"`
	Namespace string `yaml:"namespace"`
	MaxActive int    `yaml:"max_active"`
	MaxIdle   int    `yaml:"max_idle"`
}

type appConfig struct {
	Listener ListenerConfig `yaml:"listener"`
	Redis    RedisConfig    `yaml:"redis"`
}

func getRootPath() string {
	dirList := [...]string{
		"/app",
		path.Join(os.Getenv("GOPATH"), "src/github.com/sanches1984/seabattle"),
		".",
	}
	for _, dir := range dirList {
		if _, err := os.Stat(dir + "/config"); !os.IsNotExist(err) {
			return dir
		}
	}
	panic("Root path not found")
}

func getConfigPath(name string) string {
	return path.Join(getRootPath(), "config", name)
}

// Load configuration
func loadYaml() {
	var err error
	var cfg = io.ReadCloser(os.Stdin)

	cfgPath := getConfigPath(configName)
	if cfg, err = os.Open(cfgPath); err != nil {
		log.Fatal("Failed to open config:", err)
	}

	data, err := ioutil.ReadAll(cfg)
	_ = cfg.Close()
	if err != nil {
		log.Fatal("Failed to read config:", err)
	}
	log.Println("Load config from:", cfgPath)

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatal("Bad YAML format: ", err)
	}
}

func Load() {
	loadYaml()
}

func Host() string {
	return fmt.Sprintf("%s:%d", config.Listener.Host, config.Listener.Port)
}
