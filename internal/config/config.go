package config

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"sync"
)

var (
	once         sync.Once
	globalConfig *Config
	globalErr    error
)

// InitConfig 初始化配置
func InitConfig(filePath string) (*Config, error) {
	once.Do(func() {
		if filePath == "" {
			globalErr = errors.New("config file path is empty")
			fmt.Printf("[ERROR] %v\n", globalErr)
			return
		}

		fmt.Printf("[INFO] loading config from %v\n", filePath)
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Printf("[ERROR] %v\n", err)
			globalErr = err
			return
		}
		defer func() {
			_ = file.Close()
		}()

		buf, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Printf("[ERROR] read file error: %v\n", err)
			globalErr = fmt.Errorf("read file %s error", filePath)
			return
		}

		conf := &Config{}
		if err = yaml.Unmarshal(buf, conf); err != nil {
			globalErr = err
			fmt.Printf("[ERROR] %v\n", err)
			return
		}
		globalConfig = conf
	})
	return globalConfig, globalErr
}

func GetConfig() *Config {
	return globalConfig
}
