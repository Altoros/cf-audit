package config

import (
	"errors"
	"github.com/Altoros/cf-audit/cmd"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
)

type ConfigLoader interface {
	LoadCloudFoundries() (CloudFoundries, error)
}

func NewConfigLoader(opts cmd.CommandOpts, logger boshlog.Logger) (ConfigLoader, error) {
	if opts.VaultAddrOpt {
		logger.Info("Reading foundation connection info from Vault")
		loader, err := loaders.NewVaultConfigLoader(opts, logger)
		return loader, err
	} else if opts.ConfigFileOpt {
		logger.Info("Reading foundation connection info from Vault")
		loader, err := loaders.NewFileConfigLoader(opts, logger)
		return loader, err
	} else {
		return nil, errors.New("Do not know how to load CloudFoundries config")
	}
}

// func LoadCloudFoundries() (CloudFoundries, error) {}

// func loadCloudFoundriesFromFile(file string) (CloudFoundries, error) {
// 	// use opts or
// 	return nil, nil
// }

// func loadCloudFoundriesFromVault() (CloudFoundries, error) {
// 	// use opts to get vault creds
// 	// use env vars or
// 	return nil, nil
// }
// func checkFoundationClients() (CloudFoundries, error) {}

func noCloudFoundriesFoundError(source string) error {
	return errors.New("noCloudFoundriesFoundError in", source)
}
