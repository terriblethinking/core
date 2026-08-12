package config

import (
	"fmt"

	"errors"
	"os"

	"github.com/kirsle/configdir"
	"github.com/maximhq/bifrost/core/schemas"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Version   string                  `yaml:"version"`
	Providers []schemas.ModelProvider `yaml:"providers"`
}

func GetBaseConfig() (Config, error) {
	configPath := configdir.LocalConfig("harness")

	filename := fmt.Sprintf("%s/config.yaml", configPath)
	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		return Config{}, errors.New("base config does not exist")
	}

	configFile, err := os.ReadFile(filename)

	if err != nil {
		return Config{}, errors.New("could not read base config")
	}

	var config Config

	err = yaml.Unmarshal(configFile, &config)

	if err != nil {
		return Config{}, errors.New("could not unmarshal base config")
	}

	return config, nil

}
