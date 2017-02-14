package cloudfoundry

import (
	cfclient "github.com/cloudfoundry-community/go-cfclient"
)

type CloudFoundryItems []CloudFoundryItem

type CloudFoundryItem struct {
	Name   string
	Client *cfclient.Client
}

func CommonByName(list1 CloudFoundryItems, list2 CloudFoundryItems) (
	common CloudFoundryItems,
	left1 CloudFoundryItems,
	left2 CloudFoundryItems,
	err error) {
	return nil, nil, nil, nil
}
