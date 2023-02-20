package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"stack-service/internal/types"
	"strings"

	"gopkg.in/yaml.v2"
)

// DefaultProxyPort is the default port for the proxy server.
const DefaultProxyPort = 8080

// Config is the configuration of the proxy server.
type Config struct {
	DefaultPort string `yaml:"defaultPort"`
}

// New creates a new configuration for the vbalancer application.
func New() *Config {
	cfg := &Config{
		DefaultPort: "",
	}

	return cfg
}

// InitializeConfig initializes the proxy server configuration.
func (c *Config) InitProxyPort() types.ResultCode {
	osEnvValue := os.Getenv("DefaultPort")
	if osEnvValue == ":" {
		return types.ErrEmptyValue
	}

	c.DefaultPort = fmt.Sprintf(":%s", osEnvValue)
	if c.DefaultPort == ":" {
		c.DefaultPort = fmt.Sprintf(":%d", DefaultProxyPort)
	}

	c.DefaultPort = strings.Trim(c.DefaultPort, " ")

	if c.DefaultPort == strings.Trim(":", " ") {
		return types.ErrEmptyValue
	}

	return types.ResultOK
}

// Load loads the configuration for the vbalancer application.
func (c *Config) Load(cfgFileName string) error {
	searchPathConfig := []string{cfgFileName, "", "./config/", "../../config/", "../config/", "../../../config"}

	var isPathFound bool

	for _, searchPath := range searchPathConfig {
		cfgFilePath := filepath.Join(searchPath, "config.yaml")
		if _, err := os.Stat(cfgFilePath); errors.Is(err, os.ErrNotExist) {
			continue
		}

		isPathFound = true
		cfgFileName = cfgFilePath

		break
	}

	if !isPathFound {
		//nolint:goerr113
		return fmt.Errorf("failed: %s", types.ErrCantFindFile.ToStr())
	}

	fileConfig, err := os.Open(cfgFileName)
	if err != nil {
		return fmt.Errorf("failed to open config file: %w", err)
	}

	defer func(f *os.File) {
		err = f.Close()
		if err != nil {
			log.Fatalf("Error can't close config file: %s, err: %s", cfgFileName, err)
		}
	}(fileConfig)

	err = c.DecodeConfigFileYaml(fileConfig)
	if err != nil {
		return fmt.Errorf("can't decode config file: %s, err: %w", cfgFileName, err)
	}

	return nil
}

// DecodeConfigFileYaml decodes the YAML configuration file.
func (c *Config) DecodeConfigFileYaml(configYaml *os.File) error {
	decoder := yaml.NewDecoder(configYaml)
	err := decoder.Decode(c)

	if err != nil {
		return fmt.Errorf("failed to decode config yml file: %w", err)
	}

	return nil
}
