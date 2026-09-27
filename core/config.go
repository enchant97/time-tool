package core

import (
	"os"
	"path/filepath"

	"github.com/kirsle/configdir"
	"github.com/pelletier/go-toml/v2"
)

const ConfigDirName = "time-tool"
const ConfigFileName = "config.toml"

type Config struct {
	Location string `toml:"location" comment:"IANA Time Zone Location name"`
	Layout   string `toml:"layout" comment:"The layout of time e.g. RFC3339"`
}

func (c *Config) DefaultUnset() {
	c.Location = DefaultIfUnset(c.Location, "Local", "")
	c.Layout = DefaultIfUnset(c.Layout, "RFC3339", "")
}

func ReadConfig() (Config, error) {
	config := Config{}
	configFilePath := filepath.Join(configdir.LocalConfig(ConfigDirName), ConfigFileName)
	b, err := os.ReadFile(configFilePath)
	if err != nil {
		return config, err
	}
	err = toml.Unmarshal(b, &config)
	config.DefaultUnset()
	return config, err
}

func WriteConfig(config Config) error {
	configPath := configdir.LocalConfig(ConfigDirName)
	if err := os.MkdirAll(configPath, 0755); err != nil {
		return err
	}
	configFilePath := filepath.Join(configPath, ConfigFileName)
	b, err := toml.Marshal(config)
	if err != nil {
		return err
	}
	os.WriteFile(configFilePath, b, 0644)
	return nil
}
