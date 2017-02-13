package loaders

import (
	"errors"
	"github.com/Altoros/cf-audit/cmd"
	cfclient "github.com/cloudfoundry-community/go-cfclient"
	"github.com/cloudfoundry-incubator/candiedyaml"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	"os"
)

type fileConfigLoader struct {
	ConfigFilePath string
	logger         boshlog.Logger
}

type fileConfig struct {
	CloudFoundries CloudFoundries `yaml:"foundations"`
}

func NewVaultConfigLoader(opts cmd.CommandOpts, logger boshlog.Logger) (*ConfigLoader, error) {
	if opts.ConfigFilePath == "" || opts.VaultTokenOpt == "" || opts.FoundationNamesOpt {
		return nil, errors.New("You need to define config file opts (config-file)")
	}
	return &fileConfigLoader{opts.VaultAddr, opts.VaultToken, logger}, nil
}

func (l *fileConfigLoader) LoadCloudFoundries() (*CloudFoundries, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	var config fileConfig
	if err := candiedyaml.NewDecoder(file).Decode(&config); err != nil {
		return nil, err
	}
	// TODO: add validations here
	return config.CloudFoundries, nil
}
