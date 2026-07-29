package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/maximhq/bifrost/core/schemas"
	"gopkg.in/yaml.v3"
)

type ProviderConfig struct {
	Identifier   string   `yaml:"identifier"`
	URL          string   `yaml:"url"`
	Favorites    []string `yaml:"favorites"`
	APIKeys      []string `yaml:"api_keys"`
	EnvVarAPIKey string   `yaml:"envvar_api_key"`
}

func (p ProviderConfig) CreateProvider(configDir, name, url string, apiKeys ...string) error {
	if _, err := os.Stat(fmt.Sprintf("%s/config/providers", configDir)); os.IsNotExist(err) {

	}

	return nil
}

func GetInfoForProvider(name schemas.ModelProvider) (ProviderConfig, error) {
	configDir, err := os.UserConfigDir()

	if err != nil {
		return ProviderConfig{}, err
	}

	filename := fmt.Sprintf("%s/config/providers/%s.yaml", configDir, name)
	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		return ProviderConfig{}, fmt.Errorf("config does not exist for %s", name)
	}

	configFile, err := os.ReadFile(filename)

	if err != nil {
		return ProviderConfig{}, fmt.Errorf("error reading config for %s", name)
	}

	provider := ProviderConfig{}

	err = yaml.Unmarshal(configFile, &provider)

	if err != nil {
		return ProviderConfig{}, fmt.Errorf("error unmarshaling config for %s", name)
	}

	return ProviderConfig{}, nil
}

func GetAllProviders() ([]ProviderConfig, error) {
	config, err := GetBaseConfig()

	if err != nil {
		return nil, err
	}

	providerIDs := config.Providers

	var providers []ProviderConfig

	for _, provider := range providerIDs {

		// Here we get the provider config.

		providerConfig, err := GetInfoForProvider(provider)

		// TODO Currently, when there is an error with a provider,
		// we simply return the error. This is not the ideal
		// behavior, as returning will end the function and we
		// won't be able to return any providers at all.

		if err != nil {
			return nil, err
		}

		providers = append(providers, providerConfig)
	}

	return providers, nil
}

func GetAllProviderIdentifiers() ([]schemas.ModelProvider, error) {

	providerConfigs, err := GetAllProviders()

	if err != nil {
		return nil, err
	}

	var providers []schemas.ModelProvider

	for _, provider := range providerConfigs {
		providers = append(providers, schemas.ModelProvider(provider.Identifier))
	}

	return providers, nil

}
