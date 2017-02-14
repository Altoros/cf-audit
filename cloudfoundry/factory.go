package cloudfoundry

import (
	"errors"
	"io/ioutil"

	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	yaml "gopkg.in/yaml.v2"

	"github.com/davecgh/go-spew/spew"
)

type configLoader interface {
	LoadCloudFoundries() (CloudFoundries, error)
}

func NewFactory(
	configFilePath string,
	vaultAddr string,
	vaultToken string,
	cfNamesList []string,
	logger boshlog.Logger) (configLoader, error) {

	if configFilePath != "" {
		loader, err := newFileConfigLoader(configFilePath, logger)
		return loader, err
	} else if vaultAddr == "" || vaultToken == "" {
		if len(cfNamesList) < 2 {
			return nil, errors.New("You need to define list of CloudFoundries you want to use")
		}
		loader, err := newVaultConfigLoader(vaultAddr, vaultToken, cfNamesList, logger)
		return loader, err
	}

	return nil, errors.New("Do not know how to load CloudFoundries config")
}

// LOAD FROM FILE

type fileConfigLoader struct {
	configFilePath string
	logger         boshlog.Logger
}

func newFileConfigLoader(configFilePath string, logger boshlog.Logger) (configLoader, error) {
	if configFilePath == "" {
		return nil, errors.New("You need to define config file opts (config-file)")
	}
	return &fileConfigLoader{configFilePath, logger}, nil
}

type CfConfig struct {
	CloudFoundries []CloudFoundryRef `yaml:"cloudfoundries"`
}

type CloudFoundryRef struct {
	Name       string `yaml:"name"`
	ApiAddress string `yaml:"api"`
	Username   string `yaml:"username"`
	Password   string `yaml:"password"`
}

func (l *fileConfigLoader) LoadCloudFoundries() (CloudFoundries, error) {
	data, err := ioutil.ReadFile(l.configFilePath)
	if err != nil {
		return nil, err
	}

	config := CfConfig{}
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}
	spew.Dump(config)

	// TODO: add validations here
	var cfs CloudFoundries
	for _, cfref := range config.CloudFoundries {
		cf, err := NewCloudFoundry(cfref.Name, cfref.ApiAddress, cfref.Username, cfref.Password)
		if err != nil {
			return nil, err
		}
		cfs = append(cfs, cf)
	}

	spew.Dump(cfs)

	return cfs, nil
}

// LOAD FROM VAULT

type vaultConfigLoader struct {
	VaultAddr  string
	VaultToken string
	logger     boshlog.Logger
}

func newVaultConfigLoader(
	vaultAddr string,
	vaultToken string,
	cloudfoundriesList []string,
	logger boshlog.Logger) (configLoader, error) {

	if vaultAddr == "" || vaultToken == "" || len(cloudfoundriesList) < 2 {
		return nil, errors.New("can not create config reader")
	}

	return &vaultConfigLoader{vaultAddr, vaultToken, logger}, nil
}

func (l *vaultConfigLoader) LoadCloudFoundries() (CloudFoundries, error) {
	return nil, nil
}
