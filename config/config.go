package config

import (
	"errors"
	cfclient "github.com/cloudfoundry-community/go-cfclient"
	"github.com/cloudfoundry-incubator/candiedyaml"
	"os"
)

type Foundations []Foundation

type Counfig struct {
	Foundations Foundations `yaml:"foundations"`
}

type Foundation struct {
	Name     string
	Api      string
	Username string
	Password string
}

func LoadFromFile(path string) (Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return Config{}, err
	}

	var config Config
	if err := candiedyaml.NewDecoder(file).Decode(&config); err != nil {
		return Config{}, err
	}
	// TODO: add validations here
	return config, nil
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
