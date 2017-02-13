package cmd

import (
	// boshdir "github.com/cloudfoundry/bosh-cli/director"
	"github.com/Altoros/cf-audit/comparator"
	"github.com/Altoros/cf-audit/config"
	boshui "github.com/cloudfoundry/bosh-cli/ui"
)

type DiffCmd struct {
	ui          boshui.UI
	foundations config.CloudFoundries
	comporator  comporator.Comporator
}

func NewDiffCmd(ui boshui.UI) DiffCmd {
	return DiffCmd{ui: ui}
}

func (c DiffCmd) Run(configLoader config.ConfigLoader, opts DiffCmdOpts) error {
	foundations, err := configLoader.LoadCloudFoundries()
	if err != nil {
		return err
	}

	// foundations.GetCloudFoundryClient()

	factory := comporator.FindCollisions(foundations)
	c := factory.New(foundations)
	err := c.Compare()
	collisions := c.Collisions()
	//formator := GetFormator(opts.fmtOpts)

	formator.Print(collisions)
	// InfoTable{info, c.ui}.Print()

	return nil
}
