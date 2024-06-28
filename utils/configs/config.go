package configs

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// MySQLConfig represents MySQL database configuration
type MySQLConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

// Config represents the overall application configuration
type Config struct {
	MySQL MySQLConfig `yaml:"mysql"`
	Redis RedisConfig `yaml:"redis"`
}
type RedisConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

// LoadConfig loads configuration from the given file path
func LoadConfig(filePath string) (Config, error) {
	var config Config

	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("Failed to read YAML file %s: %v", filePath, err)
		return config, err
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Printf("Failed to unmarshal YAML file %s: %v", filePath, err)
		return config, err
	}

	return config, nil
}
