package gateway

import (
	"errors"

	"github.com/maximhq/bifrost/core"
	"github.com/maximhq/bifrost/core/schemas"

	"context"

	"github.com/terriblethinking/core/account"
	"github.com/terriblethinking/core/config"
)

func InitGateway() (*bifrost.Bifrost, error) {

	client, initErr := bifrost.Init(context.Background(), schemas.BifrostConfig{
		Account: &account.Account{},
	})

	for {
		if initErr != nil {

			if initErr.Error() == "base config does not exist" {

				// If the error is connected with the config, it might
				// simply not exist, and we check for that.

				err := config.InitAllPaths()

				// If the config module had an error initing the config, we
				// simply panic. There really are no more things we can do
				// at this point.

				if err != nil {
					return nil, err
				}

			} else {

				return nil, errors.New(initErr.Error())

			}

		} else {

			break

		}

		client, initErr = bifrost.Init(context.Background(), schemas.BifrostConfig{
			Account: &account.Account{},
		})
	}

	return client, nil
}
