package account

import (
	"context"
	"fmt"
	"github.com/terriblethinking/core/config"
	"slices"

	"github.com/maximhq/bifrost/core/schemas"
)

type Account struct{}

func (a Account) GetConfiguredProviders() ([]schemas.ModelProvider, error) {

	providers, err := config.GetAllProviderIdentifiers()

	if err != nil {
		return nil, err
	}

	return providers, nil

}

func (a Account) GetKeysForProvider(ctx context.Context, provider schemas.ModelProvider) ([]schemas.Key, error) {

	providerConfig, err := config.GetInfoForProvider(provider)

	// Becuase all of the provider checking and retrival
	// happens behind the scenes in GetInfoForProvider,
	// all that we have to do here is return the error.

	if err != nil {
		return nil, err
	}

	// Here we return the keys for the requested provider.

	var APIKeys []string

	// if providerConfig.EnvVarAPIKey != "" {
	// 	APIKeys = append(APIKeys, providerConfig.EnvVarAPIKey)
	// }
	//
	// fmt.Println(APIKeys[0])

	for _, key := range providerConfig.APIKeys {
		APIKeys = append(APIKeys, key)
	}

	switch provider {
	case schemas.Ollama:
		if len(APIKeys) == 0 {

			// If the provider config doesn't have any api keys
			// configured, this would most likely mean that
			// the endpoint is local (for example llama.cpp)
			// and no api key is needed. So, we simply return an
			// empty string.

			return []schemas.Key{{
				Models: schemas.WhiteList{"*"},
				Weight: 1.0,
				OllamaKeyConfig: &schemas.OllamaKeyConfig{
					URL: *schemas.NewSecretVar(providerConfig.URL),
				},
			}}, nil
		} else {
			var keys []schemas.Key

			for _, key := range APIKeys {
				keys = append(keys, schemas.Key{
					Value:  *schemas.NewSecretVar(key),
					Models: schemas.WhiteList{"*"},
					Weight: 1.0,
					OllamaKeyConfig: &schemas.OllamaKeyConfig{
						URL: *schemas.NewSecretVar(providerConfig.URL),
					},
				})
			}

			return keys, nil
		}
	}

	if len(APIKeys) == 0 {

		// If the provider config doesn't have any api keys
		// configured, this would most likely mean that
		// the endpoint is local (for example llama.cpp)
		// and no api key is needed. So, we simply return an
		// empty string.

		return []schemas.Key{}, nil
	} else {

		// If the provider does have at least on key configured,
		// then we iterate over the entire providerConfig.APIKeys
		// and create a new []schemas.Key slice, which we then
		// return

		var keys []schemas.Key

		for _, key := range APIKeys {
			keys = append(keys, schemas.Key{
				Value:  *schemas.NewSecretVar(key),
				Models: schemas.WhiteList{"*"},
				Weight: 1.0,
			})
		}

		return keys, nil

	}
}

func (a Account) GetConfigForProvider(provider schemas.ModelProvider) (*schemas.ProviderConfig, error) {

	providers, err := config.GetAllProviderIdentifiers()

	if err != nil {
		return nil, err
	}

	if slices.Contains(providers, provider) {

		providerConfig, err := config.GetInfoForProvider(provider)

		if err != nil {
			return nil, err
		}

		// TODO Currently we simpy return the default config
		// for all configured providers, which seems far from
		// ideal.

		return &schemas.ProviderConfig{
			NetworkConfig: schemas.NetworkConfig{
				BaseURL: providerConfig.URL,
			},
			ConcurrencyAndBufferSize: schemas.DefaultConcurrencyAndBufferSize,
		}, nil
	}

	return nil, fmt.Errorf("provider %s not supported", provider)

}
