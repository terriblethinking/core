package config

import (
	"fmt"

	"errors"
	"os"

	"github.com/maximhq/bifrost/core/schemas"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Version   string                  `yaml:"version"`
	Providers []schemas.ModelProvider `yaml:"providers"`
}

func GetBaseConfig() (Config, error) {
	configDir, err := os.UserConfigDir()

	if err != nil {
		return Config{}, err
	}

	filename := fmt.Sprintf("%s/config.yaml", configDir)
	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		return Config{}, errors.New("base config does not exist")
	}

	configFile, err := os.ReadFile(filename)

	if err != nil {
		return Config{}, errors.New("could not read base config")
	}

	config := make(map[string]Config)

	err = yaml.Unmarshal(configFile, &config)

	if err != nil {
		return Config{}, errors.New("could not unmarshal base config")
	}

	return Config{}, nil

}
