package cmd

import (
	// boshdir "github.com/cloudfoundry/bosh-cli/director"
	// cfconfig "github.com/Altoros/cf-audit/cloudfoundry/config"
	cf "github.com/Altoros/cf-audit/cloudfoundry"
	// "github.com/Altoros/cf-audit/comparator"
	boshcmd "github.com/cloudfoundry/bosh-cli/cmd"
	// boshui "github.com/cloudfoundry/bosh-cli/ui"

	"github.com/davecgh/go-spew/spew"
)

type DiffCmd struct {
	// ui          boshui.UI

	asd        string
	globalOpts CommandOpts
	deps       boshcmd.BasicDeps
	// foundations config.CloudFoundries
	// comporator comparator.Comporator
}

func NewDiffCmd(c Cmd) DiffCmd {
	return DiffCmd{globalOpts: c.CommandOpts, deps: c.deps}
}

func (c DiffCmd) Run(opts *DiffOpts) error {
	cfFactory, err := cf.NewFactory(
		opts.ConfigFileOpt,
		opts.VaultAddrOpt,
		opts.VaultTokenOpt,
		opts.CloudFoundriesNamesOpt,
		c.deps.Logger)
	if err != nil {
		return err
	}

	// spew.Dump(cfFactory)

	cloudfoundries, err := cfFactory.LoadCloudFoundries()

	spew.Dump(cloudfoundries)

	if err != nil {
		return err
	}

	cloudfoundries.FindCollisions()
	// collisions, err := cloudfoundries.FindCollisions()

	// c := factory.New(foundations)
	// err := c.Compare()
	// collisions := c.Collisions()
	//formator := GetFormator(opts.fmtOpts)

	// formator.Print(collisions)
	// InfoTable{info, c.ui}.Print()

	return nil
}
