package cloudfoundry

import (
	cfclient "github.com/cloudfoundry-community/go-cfclient"
)

func NewCloudFoundryClient(apiAddress string, username string, password string) (*cfclient.Client, error) {
	clientConfig := &cfclient.Config{
		ApiAddress:        apiAddress,
		Username:          username,
		Password:          password,
		SkipSslValidation: true,
	}

	client, err := cfclient.NewClient(clientConfig)
	if err != nil {
		return nil, err
	}

	return client, nil
}
