package cmd

import (
	// boshdir "github.com/cloudfoundry/bosh-cli/director"
	boshui "github.com/cloudfoundry/bosh-cli/ui"
)

type DiffCmd struct {
	ui boshui.UI
}

func NewDiffCmd(ui boshui.UI) DiffCmd {
	return DiffCmd{ui: ui}
}

func (c DiffCmd) Run(opts DiffOpts) error {
	// info, err := c.director.Info()
	// if err != nil {
	// 	return err
	// }

	// InfoTable{info, c.ui}.Print()

	return nil
}
