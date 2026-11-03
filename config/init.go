package config

import (
	"errors"
	"fmt"
	"github.com/kirsle/configdir"
	"github.com/maximhq/bifrost/core/schemas"
	"gopkg.in/yaml.v3"
	"os"
)

func InitAllPaths() error {

	configPath := configdir.LocalConfig("harness")

	err := configdir.MakePath(configPath)

	if err != nil {
		return err
	}

	os.MkdirAll(fmt.Sprintf("%s/providers", configPath), 0777)
	os.MkdirAll(fmt.Sprintf("%s/memory", configPath), 0777)
	os.MkdirAll(fmt.Sprintf("%s/skills", configPath), 0777)
	os.MkdirAll(fmt.Sprintf("%s/tools", configPath), 0777)
	os.MkdirAll(fmt.Sprintf("%s/logs", configPath), 0777)

	if _, err := os.Stat(fmt.Sprintf("%s/config.yaml", configPath)); errors.Is(err, os.ErrNotExist) {

		config := Config{
			Version:   "1.0.0",
			Providers: []schemas.ModelProvider{},
		}

		out, err := yaml.Marshal(config)

		if err != nil {
			return err
		}

		err = os.WriteFile(fmt.Sprintf("%s/config.yaml", configPath), out, 0777)

		if err != nil {
			return err
		}

		return nil

	}

	return nil

}

