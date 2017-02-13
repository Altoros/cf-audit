package config

import (
	// "errors"
	cfclient "github.com/cloudfoundry-community/go-cfclient"
	//"github.com/cloudfoundry-incubator/candiedyaml"
	//"os"
)

type CloudFoundries []Foundation

type Foundation struct {
	Name     string
	Api      string
	Username string
	Password string
}

func (f *Foundation) CloudFoundryClient() (*cfclient.Client, error) {
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
