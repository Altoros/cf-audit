package loaders

import (
	"errors"
	"github.com/Altoros/cf-audit/cmd"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
)

type vaultConfigLoader struct {
	VaultAddr  string
	VaultToken string
	logger     boshlog.Logger
}

func NewVaultConfigLoader(opts cmd.CommandOpts, logger boshlog.Logger) (*ConfigLoader, error) {
	if opts.VaultAddrOpt == "" || opts.VaultTokenOpt == "" || opts.FoundationNamesOpt {
		return nil, errors.New("You need to define Vault opts (vault-addr, vault-token and foundation-names)")
	}
	return &vaultConfigLoader{opts.VaultAddr, opts.VaultToken, logger}, nil
}

func (l *vaultConfigLoader) LoadCloudFoundries() (CloudFoundries, error) {
	return nil
}
