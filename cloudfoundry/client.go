package client

import (
	cfclient "github.com/cloudfoundry-community/go-cfclient"
)

type CloudFoundryCredentials struct {
	Host     string
	Username string
	Password string
}

func NewCloudFoundryClient(ApiUrl string, Username string, Password string) (*cfclient.Client, error) {
	clientConfig := &cfclient.Config{
		ApiAddress: f.Api,
		Username:   f.Username,
		Password:   f.Password,
	}

	client, err := cfclient.NewClient(clientConfig)
	if err != nil {
		return nil, err
	}

	return client, nil
}
